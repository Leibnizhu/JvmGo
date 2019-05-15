package loads

import "jvmgo/instruction"
import "jvmgo/rtdata"

// aload系列方法，从本地变量表加载 引用型 变量载入栈
type ALOAD struct{ instruction.Index8Instruction }

func (self *ALOAD) Execute(frame *rtdata.Frame) {
	_aload(frame, self.Index)
}

type ALOAD_0 struct{ instruction.NoOperandsInstruction }

func (self *ALOAD_0) Execute(frame *rtdata.Frame) {
	_aload(frame, 0)
}

type ALOAD_1 struct{ instruction.NoOperandsInstruction }

func (self *ALOAD_1) Execute(frame *rtdata.Frame) {
	_aload(frame, 1)
}

type ALOAD_2 struct{ instruction.NoOperandsInstruction }

func (self *ALOAD_2) Execute(frame *rtdata.Frame) {
	_aload(frame, 2)
}

type ALOAD_3 struct{ instruction.NoOperandsInstruction }

func (self *ALOAD_3) Execute(frame *rtdata.Frame) {
	_aload(frame, 3)
}

func _aload(frame *rtdata.Frame, index uint) {
	ref := frame.LocalVars().GetRef(index) //从本地变量表读取 引用类型 变量
	frame.OperandStack().PushRef(ref) //入栈
}
