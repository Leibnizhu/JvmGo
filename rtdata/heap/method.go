package heap

import "jvmgo/classfile"

type Method struct {
	ClassMember
	maxStack                uint                                //操作栈大小
	maxLocals               uint                                //局部变量表大小
	code                    []byte                              //字节码
	argSlotCount            uint                                //方法的参数的slot数
	exceptionTable          ExceptionTable                      //异常处理表
	lineNumberTable         *classfile.LineNumberTableAttribute //源代码行数
	exceptions              *classfile.ExceptionsAttribute      //
	parameterAnnotationData []byte                              // RuntimeVisibleParameterAnnotations_attribute
	annotationDefaultData   []byte                              // AnnotationDefault_attribute
	parsedDescriptor        *MethodDescriptor
}

//从 class文件 解析 Method对象
func newMethods(class *Class, cfMethods []*classfile.MemberInfo) []*Method {
	methods := make([]*Method, len(cfMethods))
	for i, cfMethod := range cfMethods {
		methods[i] = newMethod(class, cfMethod)
	}
	return methods
}

func newMethod(class *Class, cfMethod *classfile.MemberInfo) *Method {
	method := &Method{}
	method.class = class
	method.copyMemberInfo(cfMethod)
	method.copyAttributes(cfMethod)
	md := parseMethodDescriptor(method.descriptor)
	method.parsedDescriptor = md
	method.calcArgSlotCount(md.parameterTypes) //计算方法的参数slot数
	if method.IsNative() {                     //本地方法需要注入字节码和其他信息
		method.injectCodeAttribute(md.returnType)
	}
	return method
}

//从 member_info 结构读取 操作栈大小，局部变量表大小，字节码等信息
func (self *Method) copyAttributes(cfMethod *classfile.MemberInfo) {
	if codeAttr := cfMethod.CodeAttribute(); codeAttr != nil {
		self.maxStack = codeAttr.MaxStack()
		self.maxLocals = codeAttr.MaxLocals()
		self.code = codeAttr.Code()
		self.lineNumberTable = codeAttr.LineNumberTableAttribute()
		self.exceptionTable = newExceptionTable(codeAttr.ExceptionTable(),
			self.class.constantPool)
	}
	self.exceptions = cfMethod.ExceptionsAttribute()
	self.annotationData = cfMethod.RuntimeVisibleAnnotationsAttributeData()
	self.parameterAnnotationData = cfMethod.RuntimeVisibleParameterAnnotationsAttributeData()
	self.annotationDefaultData = cfMethod.AnnotationDefaultAttributeData()
}

func (self *Method) calcArgSlotCount(paramTypes []string) {
	for _, paramType := range paramTypes { //遍历方法的所有参数类型
		self.argSlotCount++
		if paramType == "J" || paramType == "D" { //long和doubel类型占两个
			self.argSlotCount++
		}
	}
	if !self.IsStatic() { //非静态方法还有个this引用作为第一个参数
		self.argSlotCount++
	}
}

func (self *Method) injectCodeAttribute(returnType string) {
	self.maxStack = 4                  // =操作数栈至少要能放下返回值，所以暂且设为4
	self.maxLocals = self.argSlotCount //本地方法帧的局部变量表只存放参数值，所以用argSlotCount即可
	switch returnType[0] {             //根据返回值类型选择相应的返回指令
	case 'V':
		self.code = []byte{0xfe, 0xb1} // return
	case 'L', '[':
		self.code = []byte{0xfe, 0xb0} // areturn
	case 'D':
		self.code = []byte{0xfe, 0xaf} // dreturn
	case 'F':
		self.code = []byte{0xfe, 0xae} // freturn
	case 'J':
		self.code = []byte{0xfe, 0xad} // lreturn
	default:
		self.code = []byte{0xfe, 0xac} // ireturn
	}
}

//访问标志香菇按的getter 方法
func (self *Method) IsSynchronized() bool {
	return 0 != self.accessFlags&ACC_SYNCHRONIZED
}
func (self *Method) IsBridge() bool {
	return 0 != self.accessFlags&ACC_BRIDGE
}
func (self *Method) IsVarargs() bool {
	return 0 != self.accessFlags&ACC_VARARGS
}
func (self *Method) IsNative() bool {
	return 0 != self.accessFlags&ACC_NATIVE
}
func (self *Method) IsAbstract() bool {
	return 0 != self.accessFlags&ACC_ABSTRACT
}
func (self *Method) IsStrict() bool {
	return 0 != self.accessFlags&ACC_STRICT
}

func (self *Method) isConstructor() bool {
	return !self.IsStatic() && self.name == "<init>"
}
func (self *Method) isClinit() bool {
	return self.IsStatic() && self.name == "<clinit>"
}

//其他 getter 方法
func (self *Method) MaxStack() uint {
	return self.maxStack
}
func (self *Method) MaxLocals() uint {
	return self.maxLocals
}
func (self *Method) Code() []byte {
	return self.code
}
func (self *Method) ParameterAnnotationData() []byte {
	return self.parameterAnnotationData
}
func (self *Method) AnnotationDefaultData() []byte {
	return self.annotationDefaultData
}
func (self *Method) ParsedDescriptor() *MethodDescriptor {
	return self.parsedDescriptor
}
func (self *Method) ArgSlotCount() uint {
	return self.argSlotCount
}

//根据异常类和抛异常的位置查找异常处理表
func (self *Method) FindExceptionHandler(exClass *Class, pc int) int {
	handler := self.exceptionTable.findExceptionHandler(exClass, pc)
	if handler != nil {
		return handler.handlerPc //找到则返回对应的处理代码位置
	}
	return -1 //找不到则返回-1
}

//获取代码行号
func (self *Method) GetLineNumber(pc int) int {
	if self.IsNative() { //本地方法
		return -2
	}
	if self.lineNumberTable == nil { //编译时通过参数控制不输出源代码行号
		return -1
	}
	return self.lineNumberTable.GetLineNumber(pc)
}

// reflection
func (self *Method) ParameterTypes() []*Class {
	if self.argSlotCount == 0 {
		return nil
	}

	paramTypes := self.parsedDescriptor.parameterTypes
	paramClasses := make([]*Class, len(paramTypes))
	for i, paramType := range paramTypes {
		paramClassName := toClassName(paramType)
		paramClasses[i] = self.class.loader.LoadClass(paramClassName)
	}

	return paramClasses
}

func (self *Method) ReturnType() *Class {
	returnType := self.parsedDescriptor.returnType
	returnClassName := toClassName(returnType)
	return self.class.loader.LoadClass(returnClassName)
}

func (self *Method) ExceptionTypes() []*Class {
	if self.exceptions == nil {
		return nil
	}

	exIndexTable := self.exceptions.ExceptionIndexTable()
	exClasses := make([]*Class, len(exIndexTable))
	cp := self.class.constantPool

	for i, exIndex := range exIndexTable {
		classRef := cp.GetConstant(uint(exIndex)).(*ClassRef)
		exClasses[i] = classRef.ResolvedClass()
	}

	return exClasses
}
