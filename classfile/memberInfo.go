package classfile
//class文件类字段和类方法的结构体
type MemberInfo struct {
	cp ConstantPool
	accessFlags uint16 //访问标志
	nameIndex uint16 //名字的常量池指针
	descriptorIndex uint16 //字段或方法的描述符，的常量池指针
	attributes []AttributeInfo //属性表
}

//读取字段表/方法表
func readMembers (reader *ClassReader, cp ConstantPool) []*MemberInfo {
	memberCount :=  reader.readUint16() //个数
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
		accessFlags: reader.readUint16(),
		nameIndex: reader.readUint16(),
		descriptorIndex: reader.readUint16(),
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

//获取属性表中的 Code 属性
func (self *MemberInfo) CodeAttribute() *CodeAttribute {
	for _, attrInfo := range self.attributes {
		switch attrInfo.(type) {
		case *CodeAttribute:
			return attrInfo.(*CodeAttribute)
		}
	}
	return nil
}

//获取属性表中的常量属性
func (self *MemberInfo) ConstantValueAttribute() *ConstantValueAttribute {
	for _, attrInfo := range self.attributes {
		switch attrInfo.(type) {
		case *ConstantValueAttribute:
			return attrInfo.(*ConstantValueAttribute)
		}
	}
	return nil
}