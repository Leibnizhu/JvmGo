package math

import "jvmgo/instruction/base"
import "jvmgo/rtdata"
//NEG 取反系列指令

// double 取反
type DNEG struct{ base.NoOperandsInstruction  }

func (self *DNEG) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	val := stack.PopDouble()
	stack.PushDouble(-val)
}

// float 取反
type FNEG struct{ base.NoOperandsInstruction  }

func (self *FNEG) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	val := stack.PopFloat()
	stack.PushFloat(-val)
}

// int 取反
type INEG struct{ base.NoOperandsInstruction  }

func (self *INEG) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	val := stack.PopInt()
	stack.PushInt(-val)
}

// long 取反
type LNEG struct{ base.NoOperandsInstruction  }

func (self *LNEG) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	val := stack.PopLong()
	stack.PushLong(-val)
}
