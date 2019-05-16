package extended

import "jvmgo/instruction/base"
import "jvmgo/rtdata"
//GOTOW与GOTO的区别是，索引从2字节变成4字节，所以 W（wide）

type GOTO_W struct {
	offset int
}

func (self *GOTO_W) FetchOperands(reader *base.BytecodeReader) {
	self.offset = int(reader.ReadInt32())
}
func (self *GOTO_W) Execute(frame *rtdata.Frame) {
	base.Branch(frame, self.offset)
}
