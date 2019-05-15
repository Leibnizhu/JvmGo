package math

import "jvmgo/instruction"
import "jvmgo/rtdata"
//ADD 系列指令

//  double 相加
type DADD struct{ instruction.NoOperandsInstruction  }

func (self *DADD) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopDouble()
	v2 := stack.PopDouble()
	result := v1 + v2
	stack.PushDouble(result)
}

// float 相加
type FADD struct{ instruction.NoOperandsInstruction  }

func (self *FADD) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopFloat()
	v1 := stack.PopFloat()
	result := v1 + v2
	stack.PushFloat(result)
}

// int 相加
type IADD struct{ instruction.NoOperandsInstruction  }

func (self *IADD) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	result := v1 + v2
	stack.PushInt(result)
}

// long 相加
type LADD struct{ instruction.NoOperandsInstruction  }

func (self *LADD) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	result := v1 + v2
	stack.PushLong(result)
}
