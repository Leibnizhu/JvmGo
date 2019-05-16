package conversions

import "jvmgo/instruction/base"
import "jvmgo/rtdata"

// double 转 float
type D2F struct{ base.NoOperandsInstruction  }

func (self *D2F) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	d := stack.PopDouble()
	f := float32(d)
	stack.PushFloat(f)
}

// double 转 int
type D2I struct{ base.NoOperandsInstruction  }

func (self *D2I) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	d := stack.PopDouble()
	i := int32(d)
	stack.PushInt(i)
}

// double 转 long
type D2L struct{ base.NoOperandsInstruction  }

func (self *D2L) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	d := stack.PopDouble()
	l := int64(d)
	stack.PushLong(l)
}
