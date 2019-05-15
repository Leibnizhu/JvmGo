package math

import "jvmgo/instruction"
import "jvmgo/rtdata"
//DIV 除法系列指令

//  double 相除，浮点型可以计算得到NaN，无需做除数==0的判定
type DDIV struct{ instruction.NoOperandsInstruction  }

func (self *DDIV) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopDouble() //除数
	v1 := stack.PopDouble() //被除数
	result := v1 / v2
	stack.PushDouble(result)
}

// float 相除，浮点型可以计算得到NaN，无需做除数==0的判定
type FDIV struct{ instruction.NoOperandsInstruction  }

func (self *FDIV) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopFloat()
	v1 := stack.PopFloat()
	result := v1 / v2
	stack.PushFloat(result)
}

// int 相除
type IDIV struct{ instruction.NoOperandsInstruction  }

func (self *IDIV) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	if v2 == 0 {
		panic("java.lang.ArithmeticException: / by zero")
	}

	result := v1 / v2
	stack.PushInt(result)
}

// long 相除
type LDIV struct{ instruction.NoOperandsInstruction  }

func (self *LDIV) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	if v2 == 0 {
		panic("java.lang.ArithmeticException: / by zero")
	}

	result := v1 / v2
	stack.PushLong(result)
}
