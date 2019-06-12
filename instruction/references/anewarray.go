package references

import "jvmgo/instruction/base"
import "jvmgo/rtdata"
import "jvmgo/rtdata/heap"

//创建引用类型数组，两个操作数
//第一个操作数紧跟指令，uint16，指向运行时常量池的类符号引用，对应数组元素的类
//第二个操作数从操作数栈推出，为数组长度
type ANEW_ARRAY struct{ base.Index16Instruction }

func (self *ANEW_ARRAY) Execute(frame *rtdata.Frame) {
	cp := frame.Method().Class().ConstantPool() //运行时常量池
	classRef := cp.GetConstant(self.Index).(*heap.ClassRef) //数组元素的类引用
	componentClass := classRef.ResolvedClass()

	// if componentClass.InitializationNotStarted() {
	// 	thread := frame.Thread()
	// 	frame.SetNextPC(thread.PC()) // undo anewarray
	// 	thread.InitClass(componentClass)
	// 	return
	// }

	stack := frame.OperandStack()
	count := stack.PopInt()
	if count < 0 {
		panic("java.lang.NegativeArraySizeException")
	}

	arrClass := componentClass.ArrayClass()
	arr := arrClass.NewArray(uint(count))
	stack.PushRef(arr)
}
