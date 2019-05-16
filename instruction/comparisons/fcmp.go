package comparisons

import "jvmgo/instruction/base"
import "jvmgo/rtdata"
//FCMP指令系列，float 比较

type FCMPG struct{ base.NoOperandsInstruction  }

func (self *FCMPG) Execute(frame *rtdata.Frame) {
	_fcmp(frame, true)
}

type FCMPL struct{ base.NoOperandsInstruction  }

func (self *FCMPL) Execute(frame *rtdata.Frame) {
	_fcmp(frame, false)
}

//gFlag控制比较双方出现NaN时，返回什么，true=返回1 fasle返回-1
func _fcmp(frame *rtdata.Frame, gFlag bool) {
	stack := frame.OperandStack()
	v2 := stack.PopFloat()
	v1 := stack.PopFloat()
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
