package comparisons

import "jvmgo/instruction"
import "jvmgo/rtdata"
//LCMP指令，long 比较

type LCMP struct{ instruction.NoOperandsInstruction  }

func (self *LCMP) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	if v1 > v2 {
		stack.PushInt(1)
	} else if v1 == v2 {
		stack.PushInt(0)
	} else {
		stack.PushInt(-1)
	}
}
