package rtdata

import "jvmgo/rtdata/heap"

type Thread struct {
	pc int //程序计数器
	stack *Stack //JVM栈指针
}

func NewThread() *Thread {
	return &Thread{
		stack: newStack(1024),
	}
}

func (self *Thread) PC() int {
	return self.pc
}

func (self *Thread) SetPC(pc int) {
	self.pc = pc
}

func (self *Thread) PushFrame(frame *Frame) {
	self.stack.push(frame)
}

func (self *Thread) PopFrame() *Frame {
	return self.stack.pop()
}

func (self *Thread) CurrentFrame() *Frame {
	return self.stack.top() //只是看看栈顶，不能pop出来
}

func (self *Thread) TopFrame() *Frame {
	return self.stack.top() //只是看看栈顶，不能pop出来
}

func (self *Thread) NewFrame(method *heap.Method) *Frame {
	return newFrame(self, method)
}

func (self *Thread) IsStackEmpty() bool {
	return self.stack.isEmpty()
}