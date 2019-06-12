package references

import "jvmgo/instruction/base"
import "jvmgo/rtdata"

// 获取数组长度	
//一个操作数，在操作数栈顶，为数组的引用
type ARRAY_LENGTH struct{ base.NoOperandsInstruction }

func (self *ARRAY_LENGTH) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	arrRef := stack.PopRef() //操作数栈顶的数组引用
	if arrRef == nil {
		panic("java.lang.NullPointerException")
	}

	arrLen := arrRef.ArrayLength()
	stack.PushInt(arrLen) //长度推入操作数栈
}
