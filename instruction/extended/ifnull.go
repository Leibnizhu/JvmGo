package extended

import "jvmgo/instruction/base"
import "jvmgo/rtdata"
//出栈，根据栈顶 为空/非空 决定是否跳转

type IFNULL struct{ base.BranchInstruction }

func (self *IFNULL) Execute(frame *rtdata.Frame) {
	ref := frame.OperandStack().PopRef()
	if ref == nil {
		base.Branch(frame, self.Offset)
	}
}

type IFNONNULL struct{ base.BranchInstruction }

func (self *IFNONNULL) Execute(frame *rtdata.Frame) {
	ref := frame.OperandStack().PopRef()
	if ref != nil {
		base.Branch(frame, self.Offset)
	}
}
