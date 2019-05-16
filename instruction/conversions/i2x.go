package conversions

import "jvmgo/instruction/base"
import "jvmgo/rtdata"

// int 转 byte
type I2B struct{ base.NoOperandsInstruction  }

func (self *I2B) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	i := stack.PopInt()
	b := int32(int8(i)) //先强转int8，相当于byte了，截断了数据，再转回int32符合栈类型要求
	stack.PushInt(b)
}

// int 转 char
type I2C struct{ base.NoOperandsInstruction  }

func (self *I2C) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	i := stack.PopInt()
	c := int32(uint16(i))
	stack.PushInt(c)
}

// int 转 short
type I2S struct{ base.NoOperandsInstruction  }

func (self *I2S) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	i := stack.PopInt()
	s := int32(int16(i))
	stack.PushInt(s)
}

// int 转 long
type I2L struct{ base.NoOperandsInstruction  }

func (self *I2L) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	i := stack.PopInt()
	l := int64(i)
	stack.PushLong(l)
}

// int 转 float
type I2F struct{ base.NoOperandsInstruction  }

func (self *I2F) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	i := stack.PopInt()
	f := float32(i)
	stack.PushFloat(f)
}

// int 转 double
type I2D struct{ base.NoOperandsInstruction  }

func (self *I2D) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	i := stack.PopInt()
	d := float64(i)
	stack.PushDouble(d)
}
