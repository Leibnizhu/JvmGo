package rtdata

//栈帧
type Frame struct {
	lower        *Frame //下一个栈帧的指针
	localVars    LocalVars //本地变量表的指针
	operandStack *OperandStack //操作数栈的指针
	thread       *Thread //当前线程
	nextPC       int //下一个指令地址
}

func newFrame(thread *Thread, maxLocals, maxStack uint) *Frame {
	return &Frame{
		thread: thread,
		localVars:    newLocalVars(maxLocals),
		operandStack: newOperandStack(maxStack),
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

func (self *Frame) NextPC() int {
	return self.nextPC
}

func (self *Frame) SetNextPC(nextPC int) {
	self.nextPC = nextPC
}