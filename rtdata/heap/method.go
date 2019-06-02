	package heap

import "jvmgo/classfile"

type Method struct {
	ClassMember
	maxStack  uint //操作栈大小
	maxLocals uint //局部变量表大小
	code      []byte //字节码
	argSlotCount uint //方法的参数的slot数
}

//从 class文件 解析 Method对象
func newMethods(class *Class, cfMethods []*classfile.MemberInfo) []*Method {
	methods := make([]*Method, len(cfMethods))
	for i, cfMethod := range cfMethods {
		methods[i] = &Method{}
		methods[i].class = class
		methods[i].copyMemberInfo(cfMethod)
		methods[i].copyAttributes(cfMethod)
		methods[i].calcArgSlotCount() //计算方法的参数slot数
	}
	return methods
}

//从 member_info 结构读取 操作栈大小，局部变量表大小，字节码等信息
func (self *Method) copyAttributes(cfMethod *classfile.MemberInfo) {
	if codeAttr := cfMethod.CodeAttribute(); codeAttr != nil {
		self.maxStack = codeAttr.MaxStack()
		self.maxLocals = codeAttr.MaxLocals()
		self.code = codeAttr.Code()
	}
}

func (self *Method) calcArgSlotCount() {
	parsedDescriptor := parseMethodDescriptor(self.descriptor) //解析方法描述符，返回MethodDescriptor实例
	for _,paramType  := range parsedDescriptor.parameterTypes { //遍历方法的所有参数类型
		self.argSlotCount++
		if paramType == "J" || paramType == "D" { //long和doubel类型占两个
			self.argSlotCount++
		}
	}
	if !self.IsStatic() { //非静态方法还有个this引用作为第一个参数
		self.argSlotCount++
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
func (self *Method) ArgSlotCount() uint {
	return self.argSlotCount
}