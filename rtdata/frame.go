package rtdata

//栈帧
type Frame struct {
	lower        *Frame //下一个栈帧的指针
	localVars    LocalVars //本地变量表的指针
	operandStack *OperandStack //操作数栈的指针
	// todo
}

func NewFrame(maxLocals, maxStack uint) *Frame {
	return &Frame{
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
