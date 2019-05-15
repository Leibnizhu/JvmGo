package control

import "jvmgo/instruction"
import "jvmgo/rtdata"
//switch 跳转， case 不连续的情况

/*
lookupswitch
<0-3 byte pad>
defaultbyte1
defaultbyte2
defaultbyte3
defaultbyte4
npairs1
npairs2
npairs3
npairs4
match-offset pairs...
*/
type LOOKUP_SWITCH struct {
	defaultOffset int32 //默认跳转的offset
	npairs        int32 //有多少对 case-offset
	matchOffsets  []int32 //case-offset对
}

func (self *LOOKUP_SWITCH) FetchOperands(reader *instruction.BytecodeReader) {
	reader.SkipPadding()
	self.defaultOffset = reader.ReadInt32()
	self.npairs = reader.ReadInt32()
	self.matchOffsets = reader.ReadInt32s(self.npairs * 2)
}

func (self *LOOKUP_SWITCH) Execute(frame *rtdata.Frame) {
	key := frame.OperandStack().PopInt()
	for i := int32(0); i < self.npairs*2; i += 2 { //只能遍历所有case
		if self.matchOffsets[i] == key { //匹配中case， 拿出offset，进行跳转
			offset := self.matchOffsets[i+1]
			instruction.Branch(frame, int(offset))
			return
		}
	}
	//一个case都没匹配中，使用默认offset
	instruction.Branch(frame, int(self.defaultOffset))
}
