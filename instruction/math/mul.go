package math

import "jvmgo/instruction"
import "jvmgo/rtdata"
//MUL 乘法指令系列

// double 相乘
type DMUL struct{ instruction.NoOperandsInstruction  }

func (self *DMUL) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopDouble()
	v1 := stack.PopDouble()
	result := v1 * v2
	stack.PushDouble(result)
}

// float 相乘
type FMUL struct{ instruction.NoOperandsInstruction  }

func (self *FMUL) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopFloat()
	v1 := stack.PopFloat()
	result := v1 * v2
	stack.PushFloat(result)
}

// int 相乘
type IMUL struct{ instruction.NoOperandsInstruction  }

func (self *IMUL) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	result := v1 * v2
	stack.PushInt(result)
}

// long 相乘
type LMUL struct{ instruction.NoOperandsInstruction  }

func (self *LMUL) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	result := v1 * v2
	stack.PushLong(result)
}
