package control

import "jvmgo/instruction/base"
import "jvmgo/rtdata"
//return 系列指令, 没有操作数
//return 只要将当前栈帧出栈即可
//其他 *return 指令，先将当前栈帧出栈，然后从其操作数栈出栈一个返回值，并将返回值入栈到原来上一层栈帧的操作数栈

// return void 
type RETURN struct{ base.NoOperandsInstruction }

func (self *RETURN) Execute(frame *rtdata.Frame) {
	frame.Thread().PopFrame()
}

// return 对象引用
type ARETURN struct{ base.NoOperandsInstruction }

func (self *ARETURN) Execute(frame *rtdata.Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame() //当前栈帧出栈
	invokerFrame := thread.TopFrame()
	ref := currentFrame.OperandStack().PopRef()
	invokerFrame.OperandStack().PushRef(ref)
}

// return double 
type DRETURN struct{ base.NoOperandsInstruction }

func (self *DRETURN) Execute(frame *rtdata.Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokerFrame := thread.TopFrame()
	val := currentFrame.OperandStack().PopDouble()
	invokerFrame.OperandStack().PushDouble(val)
}

// return float 
type FRETURN struct{ base.NoOperandsInstruction }

func (self *FRETURN) Execute(frame *rtdata.Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokerFrame := thread.TopFrame()
	val := currentFrame.OperandStack().PopFloat()
	invokerFrame.OperandStack().PushFloat(val)
}

// return int 
type IRETURN struct{ base.NoOperandsInstruction }

func (self *IRETURN) Execute(frame *rtdata.Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokerFrame := thread.TopFrame()
	val := currentFrame.OperandStack().PopInt()
	invokerFrame.OperandStack().PushInt(val)
}

// return double
type LRETURN struct{ base.NoOperandsInstruction }

func (self *LRETURN) Execute(frame *rtdata.Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokerFrame := thread.TopFrame()
	val := currentFrame.OperandStack().PopLong()
	invokerFrame.OperandStack().PushLong(val)
}
