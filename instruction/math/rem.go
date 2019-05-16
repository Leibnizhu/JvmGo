package math

import "math"
import "jvmgo/instruction/base"
import "jvmgo/rtdata"
//REM 求余运算指令系列

// double 求余
//go 没有对浮点型定义求余，这里用math包
//支持NaN，除数==0不需要抛异常
type DREM struct{ base.NoOperandsInstruction  }

func (self *DREM) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopDouble()
	v1 := stack.PopDouble()
	result := math.Mod(v1, v2)
	stack.PushDouble(result)
}

// float 求余
type FREM struct{ base.NoOperandsInstruction  }

func (self *FREM) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopFloat()
	v1 := stack.PopFloat()
	result := float32(math.Mod(float64(v1), float64(v2)))
	stack.PushFloat(result)
}

// int 求余
type IREM struct{ base.NoOperandsInstruction  }

func (self *IREM) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	if v2 == 0 {
		panic("java.lang.ArithmeticException: / by zero")
	}

	result := v1 % v2
	stack.PushInt(result)
}

// long 求余
type LREM struct{ base.NoOperandsInstruction  }

func (self *LREM) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	if v2 == 0 {
		panic("java.lang.ArithmeticException: / by zero")
	}

	result := v1 % v2
	stack.PushLong(result)
}
