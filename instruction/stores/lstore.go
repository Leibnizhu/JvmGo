package stores

import "jvmgo/instruction"
import "jvmgo/rtdata"

// lstore 系列指令，出栈 long 变量，放入局部变量表
type LSTORE struct{ instruction.Index8Instruction }

func (self *LSTORE) Execute(frame *rtdata.Frame) {
	_lstore(frame, uint(self.Index))
}

type LSTORE_0 struct{ instruction.NoOperandsInstruction }

func (self *LSTORE_0) Execute(frame *rtdata.Frame) {
	_lstore(frame, 0)
}

type LSTORE_1 struct{ instruction.NoOperandsInstruction }

func (self *LSTORE_1) Execute(frame *rtdata.Frame) {
	_lstore(frame, 1)
}

type LSTORE_2 struct{ instruction.NoOperandsInstruction }

func (self *LSTORE_2) Execute(frame *rtdata.Frame) {
	_lstore(frame, 2)
}

type LSTORE_3 struct{ instruction.NoOperandsInstruction }

func (self *LSTORE_3) Execute(frame *rtdata.Frame) {
	_lstore(frame, 3)
}

func _lstore(frame *rtdata.Frame, index uint) {
	val := frame.OperandStack().PopLong()
	frame.LocalVars().SetLong(index, val)
}
