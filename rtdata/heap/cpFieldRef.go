package heap

import "jvmgo/classfile"
//字段信息结构体
type FieldRef struct {
	MemberRef
	field *Field
}

func newFieldRef(cp *ConstantPool, refInfo *classfile.ConstantFieldrefInfo) *FieldRef {
	ref := &FieldRef{}
	ref.cp = cp
	ref.copyMemberRefInfo(&refInfo.ConstantMemberrefInfo)
	return ref
}

//懒解析
func (self *FieldRef) ResolvedField() *Field {
	if self.field == nil {
		self.resolveFieldRef()
	}
	return self.field
}

// 类D通过符号引用访问类C的字段
//首先解析符号引用得到类C
//然后根据字段名和描述符查找字段
func (self *FieldRef) resolveFieldRef() {
	d := self.cp.class
	c := self.ResolvedClass()
	field := lookupField(c, self.name, self.descriptor)

	if field == nil { //找不到字段
		panic("java.lang.NoSuchFieldError")
	}
	if !field.isAccessibleTo(d) { //没有访问权限
		panic("java.lang.IllegalAccessError")
	}

	self.field = field
}

//查找字段
func lookupField(c *Class, name, descriptor string) *Field {
	//先在类C的自有字段中查找，按名字和描述符匹配来查找
	for _, field := range c.fields {
		if field.name == name && field.descriptor == descriptor {
			return field
		}
	}

	//如果类C自有类找不到字段，递归从类C直接接口中查找
	for _, iface := range c.interfaces {
		if field := lookupField(iface, name, descriptor); field != nil {
			return field
		}
	}

	//还找不到的话，从类C的父类去递归查找
	if c.superClass != nil {
		return lookupField(c.superClass, name, descriptor)
	}

	return nil
}
