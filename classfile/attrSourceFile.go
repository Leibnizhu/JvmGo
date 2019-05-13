package classfile

/*
指出源文件名
可选，定长
SourceFile_attribute {
    u2 attribute_name_index;
    u4 attribute_length; //恒为2
    u2 sourcefile_index;
}
*/
type SourceFileAttribute struct {
	cp ConstantPool
	sourceFileIndex uint16 //源文件名的常量池下标
}

func (self *SourceFileAttribute) readInfo(reader *ClassReader) {
	self.sourceFileIndex = reader.readUint16()
}

func (self *SourceFileAttribute) FileName() string {
	return self.cp.getUtf8(self.sourceFileIndex)
}
