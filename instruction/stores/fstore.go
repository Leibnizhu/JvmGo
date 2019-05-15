package stores

import "jvmgo/instruction"
import "jvmgo/rtdata"

// fstore 系列指令，出栈 float 变量，放入局部变量表
type FSTORE struct{ instruction.Index8Instruction }

func (self *FSTORE) Execute(frame *rtdata.Frame) {
	_fstore(frame, uint(self.Index))
}

type FSTORE_0 struct{ instruction.NoOperandsInstruction }

func (self *FSTORE_0) Execute(frame *rtdata.Frame) {
	_fstore(frame, 0)
}

type FSTORE_1 struct{ instruction.NoOperandsInstruction }

func (self *FSTORE_1) Execute(frame *rtdata.Frame) {
	_fstore(frame, 1)
}

type FSTORE_2 struct{ instruction.NoOperandsInstruction }

func (self *FSTORE_2) Execute(frame *rtdata.Frame) {
	_fstore(frame, 2)
}

type FSTORE_3 struct{ instruction.NoOperandsInstruction }

func (self *FSTORE_3) Execute(frame *rtdata.Frame) {
	_fstore(frame, 3)
}

func _fstore(frame *rtdata.Frame, index uint) {
	val := frame.OperandStack().PopFloat()
	frame.LocalVars().SetFloat(index, val)
}
