package instruction

import "jvmgo/rtdata"

//字节码指令接口
type Instruction interface {
	FetchOperands(reader *BytecodeReader) //读取去操作数
	Execute(frame *rtdata.Frame) //执行指令
}

//无操作数的指令
type NoOperandsInstruction struct {
	// empty
}

func (self *NoOperandsInstruction) FetchOperands(reader *BytecodeReader) {
	// 不需要读取操作数
}

//跳转指令
type BranchInstruction struct {
	Offset int
}

func (self *BranchInstruction) FetchOperands(reader *BytecodeReader) {
	self.Offset = int(reader.ReadInt16()) //读取跳转的offset uint16整数
}

//存储和加载类指令，需要读取局部变量表的索引（未操作数），单字节
type Index8Instruction struct {
	Index uint
}

func (self *Index8Instruction) FetchOperands(reader *BytecodeReader) {
	self.Index = uint(reader.ReadUint8())
}

//访问运行时变量池的指令，需要读取运行时变量吃索引（未操作数），2字节
type Index16Instruction struct {
	Index uint
}

func (self *Index16Instruction) FetchOperands(reader *BytecodeReader) {
	self.Index = uint(reader.ReadUint16())
}
