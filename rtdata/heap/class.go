package heap

import "jvmgo/classfile"

type Class struct {
	accessFlags       uint16 //访问修饰
	name              string // 当前类名
	superClassName    string //父类名
	interfaceNames    []string //接口名
	superClass        *Class //父类
	interfaces        []*Class //接口
	constantPool      *ConstantPool //常量池
	fields            []*Field //字段表
	methods           []*Method //方法表
	loader            *ClassLoader //类加载器指针
	instanceSlotCount uint //实例变量占空间大小
	staticSlotCount   uint //类变量占空间大小
	staticVars        Slots //静态变量
}

//ClassFile 对象转 Class对象
func newClass(cf *classfile.ClassFile) *Class {
	class := &Class{}
	class.accessFlags = cf.AccessFlags()
	class.name = cf.ClassName()
	class.superClassName = cf.SuperClassName()
	class.interfaceNames = cf.InterfaceNames()
	class.constantPool = newConstantPool(class, cf.ConstantPool()) //常量池转换
	class.fields = newFields(class, cf.Fields()) //字段表转换
	class.methods = newMethods(class, cf.Methods()) //方法表转换
	return class
}

//类访问标志相关的查询方法
func (self *Class) IsPublic() bool {
	return 0 != self.accessFlags&ACC_PUBLIC
}
func (self *Class) IsFinal() bool {
	return 0 != self.accessFlags&ACC_FINAL
}
func (self *Class) IsSuper() bool {
	return 0 != self.accessFlags&ACC_SUPER
}
func (self *Class) IsInterface() bool {
	return 0 != self.accessFlags&ACC_INTERFACE
}
func (self *Class) IsAbstract() bool {
	return 0 != self.accessFlags&ACC_ABSTRACT
}
func (self *Class) IsSynthetic() bool {
	return 0 != self.accessFlags&ACC_SYNTHETIC
}
func (self *Class) IsAnnotation() bool {
	return 0 != self.accessFlags&ACC_ANNOTATION
}
func (self *Class) IsEnum() bool {
	return 0 != self.accessFlags&ACC_ENUM
}

// getter方法
func (self *Class) ConstantPool() *ConstantPool {
	return self.constantPool
}
func (self *Class) StaticVars() Slots {
	return self.staticVars
}

// jvms 5.4.4
func (self *Class) isAccessibleTo(other *Class) bool {
	return self.IsPublic() ||
		self.getPackageName() == other.getPackageName()
}

//包名，即类名前面所有字符
func (self *Class) getPackageName() string {
	if i := strings.LastIndex(self.name, "/"); i >= 0 {
		return self.name[:i]
	}
	return ""
}

//main() 入口方法
func (self *Class) GetMainMethod() *Method {
	return self.getStaticMethod("main", "([Ljava/lang/String;)V")
}

//获取静态方法的方法
func (self *Class) getStaticMethod(name, descriptor string) *Method {
	for _, method := range self.methods { //遍历查找 静态的 同名的 同签名的 方法， 应该唯一
		if method.IsStatic() &&
			method.name == name &&
			method.descriptor == descriptor {

			return method
		}
	}
	return nil
}

func (self *Class) NewObject() *Object {
	return newObject(self)
}
