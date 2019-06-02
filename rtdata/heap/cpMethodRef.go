package heap

import "jvmgo/classfile"
//方法符号引用
type MethodRef struct {
	MemberRef
	method *Method
}

func newMethodRef(cp *ConstantPool, refInfo *classfile.ConstantMethodrefInfo) *MethodRef {
	ref := &MethodRef{}
	ref.cp = cp
	ref.copyMemberRefInfo(&refInfo.ConstantMemberrefInfo)
	return ref
}

func (self *MethodRef) ResolvedMethod() *Method {
	if self.method == nil {
		self.resolveMethodRef()
	}
	return self.method
}

//解析 方法的符号引用
// 类D 通过方法符号引用访问 类C 的方法
func (self *MethodRef) resolveMethodRef() {
	d := self.cp.class //当前对象的类
	c := self.ResolvedClass() //解析类C
	if c.IsInterface() { //接口
		panic("java.lang.IncompatibleClassChangeError	")
	}
	method := lookupMethod(c, self.name, self.descriptor) //根据方法名和描述符去查找方法
	if method == nil {
		panic("java.lang.NoSuchMethodError")
	}
	if !method.isAccessibleTo(d) { //没权限访问该方法
		 panic("java.lang.IllegalAccessError")
	}
	self.method = method
}

func lookupMethod(class *Class, name, descriptor string) *Method {
	method := LookupMethodInClass(class, name, descriptor) //先去类里面找
	if method == nil { //类里面找不到的话就去接口里面找
		method = lookupMethodInInterfaces(class.interfaces, name, descriptor)
	}
	return method
}
