package heap

import "jvmgo/classfile"
//字段和方法公用信息的结构体
type MemberRef struct {
	SymRef
	name       string
	descriptor string //jvm允许同名字段有不同类型，但java不支持
}

//从class文件存储的字段或方法常量中提取数据
func (self *MemberRef) copyMemberRefInfo(refInfo *classfile.ConstantMemberrefInfo) {
	self.className = refInfo.ClassName()
	self.name, self.descriptor = refInfo.NameAndDescriptor()
}

//getter 方法
func (self *MemberRef) Name() string {
	return self.name
}
func (self *MemberRef) Descriptor() string {
	return self.descriptor
}
