package heap

import "jvmgo/classfile"
//字段信息  和 方法信息 的共同父类

type ClassMember struct {
	accessFlags uint16 //访问标志
	name        string
	descriptor  string //签名，描述符
	class       *Class //所在的类对象
}

//从 ClassFile 的 MemberInfo 复制属性
func (self *ClassMember) copyMemberInfo(memberInfo *classfile.MemberInfo) {
	self.accessFlags = memberInfo.AccessFlags()
	self.name = memberInfo.Name()
	self.descriptor = memberInfo.Descriptor()
}

//访问标志 相关的get方法
func (self *ClassMember) IsPublic() bool {
	return 0 != self.accessFlags&ACC_PUBLIC
}
func (self *ClassMember) IsPrivate() bool {
	return 0 != self.accessFlags&ACC_PRIVATE
}
func (self *ClassMember) IsProtected() bool {
	return 0 != self.accessFlags&ACC_PROTECTED
}
func (self *ClassMember) IsStatic() bool {
	return 0 != self.accessFlags&ACC_STATIC
}
func (self *ClassMember) IsFinal() bool {
	return 0 != self.accessFlags&ACC_FINAL
}
func (self *ClassMember) IsSynthetic() bool {
	return 0 != self.accessFlags&ACC_SYNTHETIC
}

//其他 getter方法
func (self *ClassMember) Name() string {
	return self.name
}
func (self *ClassMember) Descriptor() string {
	return self.descriptor
}
func (self *ClassMember) Class() *Class {
	return self.class
}

// jvms 5.4.4
func (self *ClassMember) isAccessibleTo(d *Class) bool {
	if self.IsPublic() {
		return true
	}
	c := self.class
	if self.IsProtected() {
		return d == c || d.isSubClassOf(c) ||
			c.getPackageName() == d.getPackageName()
	}
	if !self.IsPrivate() {
		return c.getPackageName() == d.getPackageName()
	}
	return d == c
}
