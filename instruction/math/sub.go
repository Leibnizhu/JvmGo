package math

import "jvmgo/instruction/base"
import "jvmgo/rtdata"
//SUB 减法系列命令

// double 相减
type DSUB struct{ base.NoOperandsInstruction  }

func (self *DSUB) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopDouble()
	v1 := stack.PopDouble()
	result := v1 - v2
	stack.PushDouble(result)
}

// float 相减
type FSUB struct{ base.NoOperandsInstruction  }

func (self *FSUB) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopFloat()
	v1 := stack.PopFloat()
	result := v1 - v2
	stack.PushFloat(result)
}

// int 相减
type ISUB struct{ base.NoOperandsInstruction  }

func (self *ISUB) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	result := v1 - v2
	stack.PushInt(result)
}

// long 相减
type LSUB struct{ base.NoOperandsInstruction  }

func (self *LSUB) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	result := v1 - v2
	stack.PushLong(result)
}
