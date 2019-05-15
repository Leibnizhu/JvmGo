package math

import "jvmgo/instruction"
import "jvmgo/rtdata"
// OR 按位或系列指令

// int 按位或
type IOR struct{ instruction.NoOperandsInstruction  }

func (self *IOR) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	result := v1 | v2
	stack.PushInt(result)
}

// long 按位或
type LOR struct{ instruction.NoOperandsInstruction  }

func (self *LOR) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	result := v1 | v2
	stack.PushLong(result)
}
