package math

import "jvmgo/instruction/base"
import "jvmgo/rtdata"
//AND 系列指令

//INT 的按位与运算
type IAND struct{ base.NoOperandsInstruction  }

func (self *IAND) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	result := v1 & v2
	stack.PushInt(result)
}

// long 的按位与运算
type LAND struct{ base.NoOperandsInstruction  }

func (self *LAND) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	result := v1 & v2
	stack.PushLong(result)
}
