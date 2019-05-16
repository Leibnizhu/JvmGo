package stack

import "jvmgo/instruction/base"
import "jvmgo/rtdata"

// 交换栈顶2个元素
type SWAP struct{ base.NoOperandsInstruction  }

/*
[...][c][b][a]
变成
[...][c][a][b]
*/
func (self *SWAP) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	stack.PushSlot(slot1)
	stack.PushSlot(slot2)
}