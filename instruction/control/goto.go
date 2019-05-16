package control

import "jvmgo/instruction/base"
import "jvmgo/rtdata"
//GOTO 无条件跳转指令

type GOTO struct{ base.BranchInstruction }

func (self *GOTO) Execute(frame *rtdata.Frame) {
	base.Branch(frame, self.Offset)
}
