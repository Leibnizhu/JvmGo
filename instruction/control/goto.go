package control

import "jvmgo/instruction"
import "jvmgo/rtdata"
//GOTO 无条件跳转指令

type GOTO struct{ instruction.BranchInstruction }

func (self *GOTO) Execute(frame *rtdata.Frame) {
	base.Branch(frame, self.Offset)
}
