package loads

import "jvmgo/instruction"
import "jvmgo/rtdata"

// dload系列方法，从本地变量表加载 double 变量载入栈
type DLOAD struct{ instruction.Index8Instruction }

func (self *DLOAD) Execute(frame *rtdata.Frame) {
	_dload(frame, self.Index)
}

type DLOAD_0 struct{ instruction.NoOperandsInstruction }

func (self *DLOAD_0) Execute(frame *rtdata.Frame) {
	_dload(frame, 0)
}

type DLOAD_1 struct{ instruction.NoOperandsInstruction }

func (self *DLOAD_1) Execute(frame *rtdata.Frame) {
	_dload(frame, 1)
}

type DLOAD_2 struct{ instruction.NoOperandsInstruction }

func (self *DLOAD_2) Execute(frame *rtdata.Frame) {
	_dload(frame, 2)
}

type DLOAD_3 struct{ instruction.NoOperandsInstruction }

func (self *DLOAD_3) Execute(frame *rtdata.Frame) {
	_dload(frame, 3)
}

func _dload(frame *rtdata.Frame, index uint) {
	val := frame.LocalVars().GetDouble(index)
	frame.OperandStack().PushDouble(val)
}
