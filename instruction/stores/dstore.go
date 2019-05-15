package stores

import "jvmgo/instruction"
import "jvmgo/rtdata"

// dstore 系列指令，出栈 double 变量，放入局部变量表
type DSTORE struct{ instruction.Index8Instruction }

func (self *DSTORE) Execute(frame *rtdata.Frame) {
	_dstore(frame, uint(self.Index))
}

type DSTORE_0 struct{ instruction.NoOperandsInstruction }

func (self *DSTORE_0) Execute(frame *rtdata.Frame) {
	_dstore(frame, 0)
}

type DSTORE_1 struct{ instruction.NoOperandsInstruction }

func (self *DSTORE_1) Execute(frame *rtdata.Frame) {
	_dstore(frame, 1)
}

type DSTORE_2 struct{ instruction.NoOperandsInstruction }

func (self *DSTORE_2) Execute(frame *rtdata.Frame) {
	_dstore(frame, 2)
}

type DSTORE_3 struct{ instruction.NoOperandsInstruction }

func (self *DSTORE_3) Execute(frame *rtdata.Frame) {
	_dstore(frame, 3)
}

func _dstore(frame *rtdata.Frame, index uint) {
	val := frame.OperandStack().PopDouble()
	frame.LocalVars().SetDouble(index, val)
}
