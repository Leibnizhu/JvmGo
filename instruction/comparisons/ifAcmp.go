package comparisons

import "jvmgo/instruction"
import "jvmgo/rtdata"
//ACMP 系列指令，判定两个引用型变量是否同一个对象，来进行跳转

type IF_ACMPEQ struct{ instruction.BranchInstruction }

func (self *IF_ACMPEQ) Execute(frame *rtdata.Frame) {
	if _acmp(frame) {
		instruction.Branch(frame, self.Offset)
	}
}

type IF_ACMPNE struct{ instruction.BranchInstruction }

func (self *IF_ACMPNE) Execute(frame *rtdata.Frame) {
	if !_acmp(frame) {
		instruction.Branch(frame, self.Offset)
	}
}

func _acmp(frame *rtdata.Frame) bool {
	stack := frame.OperandStack()
	ref2 := stack.PopRef()
	ref1 := stack.PopRef()
	return ref1 == ref2 // todo
}
