package stack

import "jvmgo/instruction/base"
import "jvmgo/rtdata"

// 出栈 int float等类型
type POP struct{ base.NoOperandsInstruction  }

func (self *POP) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	stack.PopSlot()
}

// double 或 long 类型出栈，占操作数栈的两个位置
type POP2 struct{ base.NoOperandsInstruction  }

func (self *POP2) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	stack.PopSlot()
	stack.PopSlot()
}
