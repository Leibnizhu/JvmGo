package constants

import "jvmgo/instruction/base"
import "jvmgo/rtdata"

// NOP 指令啥都不做
type NOP struct{ base.NoOperandsInstruction }

func (self *NOP) Execute(frame *rtdata.Frame) {
}