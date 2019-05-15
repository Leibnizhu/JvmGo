package stores

import "jvmgo/instruction"
import "jvmgo/rtdata"

// istore 系列指令，出栈 int 变量，放入局部变量表
type ISTORE struct{ instruction.Index8Instruction }

func (self *ISTORE) Execute(frame *rtdata.Frame) {
	_istore(frame, uint(self.Index))
}

type ISTORE_0 struct{ instruction.NoOperandsInstruction }

func (self *ISTORE_0) Execute(frame *rtdata.Frame) {
	_istore(frame, 0)
}

type ISTORE_1 struct{ instruction.NoOperandsInstruction }

func (self *ISTORE_1) Execute(frame *rtdata.Frame) {
	_istore(frame, 1)
}

type ISTORE_2 struct{ instruction.NoOperandsInstruction }

func (self *ISTORE_2) Execute(frame *rtdata.Frame) {
	_istore(frame, 2)
}

type ISTORE_3 struct{ instruction.NoOperandsInstruction }

func (self *ISTORE_3) Execute(frame *rtdata.Frame) {
	_istore(frame, 3)
}

func _istore(frame *rtdata.Frame, index uint) {
	val := frame.OperandStack().PopInt()
	frame.LocalVars().SetInt(index, val)
}
