package stack

import "jvmgo/instruction"
import "jvmgo/rtdata"

// 重复栈顶元素 int float 等类型
type DUP struct{ instruction.NoOperandsInstruction  }

/*
[...][c][b][a]
变成
[...][c][b] [a] [a]
*/
func (self *DUP) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	slot := stack.PopSlot()
	stack.PushSlot(slot)
	stack.PushSlot(slot)
}

// 重复栈顶元素， 插入栈顶2个元素以后
type DUP_X1 struct{ instruction.NoOperandsInstruction  }

/*
[...][c][b][a}
变成
[...][c] [a] [b][a]
*/
func (self *DUP_X1) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	stack.PushSlot(slot1)
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
}

// 重复栈顶元素， 插入到栈顶第3个元素以后
type DUP_X2 struct{ instruction.NoOperandsInstruction  }

/*
[...][c][b][a]
变成
[...] [a] [c][b][a]
*/
func (self *DUP_X2) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	slot3 := stack.PopSlot()
	stack.PushSlot(slot1)
	stack.PushSlot(slot3)
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
}

// 复制栈顶2个元素到 栈顶第2个元素后面
type DUP2 struct{ instruction.NoOperandsInstruction  }

/*
[...][c][b][a]
变成
[...][c] [b][a] [b][a]
*/
func (self *DUP2) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
}

// 复制栈顶2个元素到 栈顶第3个元素后面
type DUP2_X1 struct{ instruction.NoOperandsInstruction  }

/*
[...][c][b][a]
变成
[...] [b][a] [c][b][a]
*/
func (self *DUP2_X1) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	slot3 := stack.PopSlot()
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
	stack.PushSlot(slot3)
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
}

//  复制栈顶2个元素到 栈顶第4个元素后面
type DUP2_X2 struct{ instruction.NoOperandsInstruction  }

/*
[...][d][c][b][a]
变成
[...] [b][a] [d][c][b][a]
*/
func (self *DUP2_X2) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	slot3 := stack.PopSlot()
	slot4 := stack.PopSlot()
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
	stack.PushSlot(slot4)
	stack.PushSlot(slot3)
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
}