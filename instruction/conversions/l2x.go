package conversions

import "jvmgo/instruction"
import "jvmgo/rtdata"
//I2X 系列 int转其他类型指令

// long 转 double
type L2D struct{ instruction.NoOperandsInstruction  }

func (self *L2D) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	l := stack.PopLong()
	d := float64(l)
	stack.PushDouble(d)
}

// long 转 float
type L2F struct{ instruction.NoOperandsInstruction  }

func (self *L2F) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	l := stack.PopLong()
	f := float32(l)
	stack.PushFloat(f)
}

// long 转 int
type L2I struct{ instruction.NoOperandsInstruction  }

func (self *L2I) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	l := stack.PopLong()
	i := int32(l)
	stack.PushInt(i)
}
