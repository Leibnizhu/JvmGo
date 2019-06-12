package heap

import "jvmgo/classfile"
import "strings"

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
	initStarted       bool //类是否已初始化，即<clinit>是否已经执行
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
func (self *Class) Name() string {
	return self.name
}
func (self *Class) Fields() []*Field {
	return self.fields
}
func (self *Class) Methods() []*Method {
	return self.methods
}
func (self *Class) SuperClass() *Class {
	return self.superClass
}
func (self *Class) Loader() *ClassLoader {
	return self.loader
}

func (self *Class) InitStarted() bool {
	return self.initStarted
}

func (self *Class) StartInit() {
	self.initStarted = true
}

// 当前类是否可被参数的类访问
//要么时public，要么是同运行时包
func (self *Class) isAccessibleTo(other *Class) bool {
	return self.IsPublic() ||
		self.GetPackageName() == other.GetPackageName()
}

//包名，即类名前面所有字符
func (self *Class) GetPackageName() string {
	if i := strings.LastIndex(self.name, "/"); i >= 0 {
		return self.name[:i]
	}
	return ""
}

//main() 入口方法
func (self *Class) GetMainMethod() *Method {
	return self.getStaticMethod("main", "([Ljava/lang/String;)V")
}
func (self *Class) GetClinitMethod() *Method {
	return self.getStaticMethod("<clinit>", "()V")
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

//返回与类相应的数组类
func (self *Class) ArrayClass() *Class {
	arrayClassName := getArrayClassName(self.name) //根据类名得到数组名
	return self.loader.LoadClass(arrayClassName)
}
