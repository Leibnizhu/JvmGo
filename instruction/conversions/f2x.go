package conversions

import "jvmgo/instruction/base"
import "jvmgo/rtdata"

// float 转 double
type F2D struct{ base.NoOperandsInstruction  }

func (self *F2D) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	f := stack.PopFloat()
	d := float64(f)
	stack.PushDouble(d)
}

// float 转 int
type F2I struct{ base.NoOperandsInstruction  }

func (self *F2I) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	f := stack.PopFloat()
	i := int32(f)
	stack.PushInt(i)
}

// float 转 long
type F2L struct{ base.NoOperandsInstruction  }

func (self *F2L) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	f := stack.PopFloat()
	l := int64(f)
	stack.PushLong(l)
}
