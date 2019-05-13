package classfile

/*
定长，field_info中出现，表示常量表达式的值
ConstantValue_attribute {
    u2 attribute_name_index;
    u4 attribute_length; //恒等2
    u2 constantvalue_index; 
}
*/
type ConstantValueAttribute struct {
	constantValueIndex uint16 //常量池索引
}

func (self *ConstantValueAttribute) readInfo(reader *ClassReader) {
	self.constantValueIndex = reader.readUint16()
}

func (self *ConstantValueAttribute) ConstantValueIndex() uint16 {
	return self.constantValueIndex
}
