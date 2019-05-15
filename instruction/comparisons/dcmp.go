package comparisons

import "jvmgo/instruction"
import "jvmgo/rtdata"
//DCMP指令系列，double 比较

type DCMPG struct{ instruction.NoOperandsInstruction  }

func (self *DCMPG) Execute(frame *rtdata.Frame) {
	_dcmp(frame, true)
}

type DCMPL struct{ instruction.NoOperandsInstruction  }

func (self *DCMPL) Execute(frame *rtdata.Frame) {
	_dcmp(frame, false)
}

//gFlag控制比较双方出现NaN时，返回什么，true=返回1 fasle返回-1
func _dcmp(frame *rtdata.Frame, gFlag bool) {
	stack := frame.OperandStack()
	v2 := stack.PopDouble()
	v1 := stack.PopDouble()
	if v1 > v2 {
		stack.PushInt(1)
	} else if v1 == v2 {
		stack.PushInt(0)
	} else if v1 < v2 {
		stack.PushInt(-1)
	} else if gFlag {
		stack.PushInt(1)
	} else {
		stack.PushInt(-1)
	}
}
