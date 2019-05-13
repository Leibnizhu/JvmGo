package classfile

/*
java.lang.String字面量，
并没有持有String常量的内容，而是指向另一个常量类型为CONSTANT_Utf8的常量
结构如下：
CONSTANT_String_info {
    u1 tag;
    u2 string_index;
}
*/
type ConstantStringInfo struct {
	cp ConstantPool
	stringIndex uint16
}

//读取到string的时候，可能对应的utf8常量还没读入到 ConstPool ，
// 所以这里只保存ConstPool的指针，等String()方法调用的时候才去ConstPool读取字符串真正的值
func (self *ConstantStringInfo) readInfo(reader *ClassReader) {
	self.stringIndex = reader.readUint16()
}
func (self *ConstantStringInfo) String() string {
	return self.cp.getUtf8(self.stringIndex)
}
