package classfile

/*
存放方法的行号信息，通过 javac -g:none 关闭这些信息输出到class
LineNumberTable_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 line_number_table_length; //行号信息表长度
    {   u2 start_pc;
        u2 line_number; //行号
    } line_number_table[line_number_table_length]; //行号信息表
}
*/
type LineNumberTableAttribute struct {
	lineNumberTable []*LineNumberTableEntry
}

type LineNumberTableEntry struct {
	startPc    uint16
	lineNumber uint16
}

func (self *LineNumberTableAttribute) readInfo(reader *ClassReader) {
	lineNumberTableLength := reader.readUint16() //行号信息表长度
	self.lineNumberTable = make([]*LineNumberTableEntry, lineNumberTableLength)
	for i := range self.lineNumberTable {
		self.lineNumberTable[i] = &LineNumberTableEntry{ //读取一个行号信息对象
			startPc:    reader.readUint16(),
			lineNumber: reader.readUint16(),
		}
	}
}

//获取pc对应的源码行号
func (self *LineNumberTableAttribute) GetLineNumber(pc int) int {
	for i := len(self.lineNumberTable) - 1; i >= 0; i-- { //遍历查找
		entry := self.lineNumberTable[i]
		if pc >= int(entry.startPc) {
			return int(entry.lineNumber)
		}
	}
	return -1
}
