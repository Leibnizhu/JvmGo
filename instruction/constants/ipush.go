package constants

import "jvmgo/instruction/base"
import "jvmgo/rtdata"

// 从操作数读取一个byte，扩展成int，push进栈顶
type BIPUSH struct {
	val int8
}

func (self *BIPUSH) FetchOperands(reader *base.BytecodeReader) {
	self.val = reader.ReadInt8()
}
func (self *BIPUSH) Execute(frame *rtdata.Frame) {
	i := int32(self.val)
	frame.OperandStack().PushInt(i)
}

// 从操作数读取一个 short，扩展成int，push进栈顶
type SIPUSH struct {
	val int16
}

func (self *SIPUSH) FetchOperands(reader *base.BytecodeReader) {
	self.val = reader.ReadInt16()
}
func (self *SIPUSH) Execute(frame *rtdata.Frame) {
	i := int32(self.val)
	frame.OperandStack().PushInt(i)
}
