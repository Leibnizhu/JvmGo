package classfile
//class文件类字段和类方法的结构体
type MemberInfo struct {
	cp ConstPool
	accessFlags uint16 //访问标志
	nameIndex uint16 //名字的常量池指针
	descriptorIndex uint16 //字段或方法的描述符，的常量池指针
	attributes []Attribute //属性表
}

//读取字段表/方法表
func readMembers (reader *ClassReader, cp ConstantPool) []*MemberInfo {
	memberCount :=  reader.readerUint16() //个数
	members := make([]*MemberInfo, memberCount) //初始化存放字段表/方法表的数组
	for i := range members {
		members[i] = readMember(reader, cp) //读取单个字段/方法
	}
	return members
}

//读取单个字段/方法
func readMember (reader *ClassReader, cp ConstantPool) *MemberInfo {
	return &MemberInfo {
		cp: cp,
		accessFlags: reader.readerUint16(),
		nameIndex: reader.readerUint16(),
		descriptorIndex: reader.readerUint16(),
		attributes: readAttributes(reader, cp), //读取属性表
	}
}

func (self *MemberInfo) AccessFlags() uint16 {
	return self.accessFlags
}

func (self *MemberInfo) Name() string {
	return self.cp.getUtf8(self.nameIndex)
}

func (self *MemberInfo) Descriptor() string {
	return self.cp.getUtf8(self.descriptorIndex)
}