package loads

import "jvmgo/instruction/base"
import "jvmgo/rtdata"

// fload系列方法，从本地变量表加载float变量载入栈
type FLOAD struct{ base.Index8Instruction }

func (self *FLOAD) Execute(frame *rtdata.Frame) {
	_fload(frame, self.Index)
}

type FLOAD_0 struct{ base.NoOperandsInstruction }

func (self *FLOAD_0) Execute(frame *rtdata.Frame) {
	_fload(frame, 0)
}

type FLOAD_1 struct{ base.NoOperandsInstruction }

func (self *FLOAD_1) Execute(frame *rtdata.Frame) {
	_fload(frame, 1)
}

type FLOAD_2 struct{ base.NoOperandsInstruction }

func (self *FLOAD_2) Execute(frame *rtdata.Frame) {
	_fload(frame, 2)
}

type FLOAD_3 struct{ base.NoOperandsInstruction }

func (self *FLOAD_3) Execute(frame *rtdata.Frame) {
	_fload(frame, 3)
}

//通用float指令执行方法
func _fload(frame *rtdata.Frame, index uint) {
	val := frame.LocalVars().GetFloat(index) //从本地变量表读取 Float 变量
	frame.OperandStack().PushFloat(val) //入栈
}
