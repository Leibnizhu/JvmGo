package math

import "jvmgo/instruction"
import "jvmgo/rtdata"

// 给局部变量表中的int变量增加指定常量值
type IINC struct {
	Index uint //局部变量表的下标
	Const int32 //指定要加的常量
}

func (self *IINC) FetchOperands(reader *instruction.BytecodeReader) {
	self.Index = uint(reader.ReadUint8())
	self.Const = int32(reader.ReadInt8())
}

func (self *IINC) Execute(frame *rtdata.Frame) {
	localVars := frame.LocalVars()
	val := localVars.GetInt(self.Index)
	val += self.Const
	localVars.SetInt(self.Index, val)
}
