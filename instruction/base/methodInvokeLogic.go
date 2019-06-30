package base

import "jvmgo/rtdata"
import "jvmgo/rtdata/heap"

//调用方法的一些通用逻辑
func InvokeMethod(invokerFrame *rtdata.Frame, method *heap.Method) {
	thread := invokerFrame.Thread() //当前线程
	newFrame := thread.NewFrame(method) //创建栈帧
	thread.PushFrame(newFrame) //栈帧入栈

	argSlotCount := int(method.ArgSlotCount()) //方法参数的slot数量
	if argSlotCount > 0 { //方法有参数（注：非静态方法第一个参数是隐藏的this）
		for i := argSlotCount - 1; i >= 0; i-- { //从本地变量表的尾部开始设置变量，保持顺序一致
			 //从操作数栈出栈参数，入栈新栈帧的本地变量表
			slot := invokerFrame.OperandStack().PopSlot()
			newFrame.LocalVars().SetSlot(uint(i), slot)
		}
	}
}
