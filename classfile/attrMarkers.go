package classfile

/*
Deprecated标记
Deprecated_attribute {
    u2 attribute_name_index;
    u4 attribute_length;  //恒为0
}
*/
type DeprecatedAttribute struct {
	MarkerAttribute
}

/*
Synthetic标记，标记源文件中没有的类，为了支持嵌套类和嵌套接口
结构如下
Synthetic_attribute {
    u2 attribute_name_index;
    u4 attribute_length;  //恒为0
}
*/
type SyntheticAttribute struct {
	MarkerAttribute
}

type MarkerAttribute struct{}

func (self *MarkerAttribute) readInfo(reader *ClassReader) {
	// 没有数据，所以不需要读
}
