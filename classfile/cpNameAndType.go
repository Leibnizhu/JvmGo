package classfile

/*
字段或方法的名字及描述符
有描述符，就可以区分重载的同名方法了
CONSTANT_NameAndType_info {
    u1 tag;
    u2 name_index;
    u2 descriptor_index; //如 (Ljava.lang.String;)V
}
*/
type ConstantNameAndTypeInfo struct {
	nameIndex       uint16
	descriptorIndex uint16
}

func (self *ConstantNameAndTypeInfo) readInfo(reader *ClassReader) {
	self.nameIndex = reader.readUint16()
	self.descriptorIndex = reader.readUint16()
}
