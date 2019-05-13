package classfile

/*
变长，记录方法抛出的异常表
Exceptions_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 number_of_exceptions; //异常个数
    u2 exception_index_table[number_of_exceptions]; //异常表
}
*/
type ExceptionsAttribute struct {
	exceptionIndexTable []uint16
}

func (self *ExceptionsAttribute) readInfo(reader *ClassReader) {
	self.exceptionIndexTable = reader.readUint16s() //正好是个uint16 表，调readUint16s() 方法读取即可
}

func (self *ExceptionsAttribute) ExceptionIndexTable() []uint16 {
	return self.exceptionIndexTable
}
