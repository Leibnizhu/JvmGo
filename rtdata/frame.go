package rtdata

import "jvmgo/rtdata/heap"

//栈帧
type Frame struct {
	lower        *Frame //下一个栈帧的指针
	localVars    LocalVars //本地变量表的指针
	operandStack *OperandStack //操作数栈的指针
	thread       *Thread //当前线程
	method       *heap.Method //当前方法
	nextPC       int //下一个指令地址
}

func newFrame(thread *Thread, method *heap.Method) *Frame {
	return &Frame{
		thread: thread,
		method: method,
		localVars:    newLocalVars(method.MaxLocals()),
		operandStack: newOperandStack(method.MaxStack()),
	}
}

// getters
func (self *Frame) LocalVars() LocalVars {
	return self.localVars
}
func (self *Frame) OperandStack() *OperandStack {
	return self.operandStack
}

func (self *Frame) Thread() *Thread {
	return self.thread
}

func (self *Frame) Method() *heap.Method {
	return self.method
}

func (self *Frame) NextPC() int {
	return self.nextPC
}

func (self *Frame) SetNextPC(nextPC int) {
	self.nextPC = nextPC
}

func (self *Frame) RevertNextPC() {
	self.nextPC = self.thread.pc
}
