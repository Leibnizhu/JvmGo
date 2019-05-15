package constants

import "jvmgo/instruction"
import "jvmgo/rtdata"

// NOP 指令啥都不做
type NOP struct{ instruction.NoOperandsInstruction }

func (self *NOP) Execute(frame *rtdata.Frame) {
}