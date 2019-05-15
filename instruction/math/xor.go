package math

import "jvmgo/instruction"
import "jvmgo/rtdata"
// XOR 按位异或系列指令

// int 按位异或
type IXOR struct{ instruction.NoOperandsInstruction  }

func (self *IXOR) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopInt()
	v2 := stack.PopInt()
	result := v1 ^ v2
	stack.PushInt(result)
}

// Long 按位异或
type LXOR struct{ instruction.NoOperandsInstruction  }

func (self *LXOR) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopLong()
	v2 := stack.PopLong()
	result := v1 ^ v2
	stack.PushLong(result)
}
