package loads

import "jvmgo/instruction"
import "jvmgo/rtdata"

// lload系列方法，从本地变量表加载 long 变量载入栈
type LLOAD struct{ instruction.Index8Instruction }

func (self *LLOAD) Execute(frame *rtdata.Frame) {
	_lload(frame, self.Index)
}

type LLOAD_0 struct{ instruction.NoOperandsInstruction }

func (self *LLOAD_0) Execute(frame *rtdata.Frame) {
	_lload(frame, 0)
}

type LLOAD_1 struct{ instruction.NoOperandsInstruction }

func (self *LLOAD_1) Execute(frame *rtdata.Frame) {
	_lload(frame, 1)
}

type LLOAD_2 struct{ instruction.NoOperandsInstruction }

func (self *LLOAD_2) Execute(frame *rtdata.Frame) {
	_lload(frame, 2)
}

type LLOAD_3 struct{ instruction.NoOperandsInstruction }

func (self *LLOAD_3) Execute(frame *rtdata.Frame) {
	_lload(frame, 3)
}

func _lload(frame *rtdata.Frame, index uint) {
	val := frame.LocalVars().GetLong(index)
	frame.OperandStack().PushLong(val)
}
