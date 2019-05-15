package control

import "jvmgo/instruction"
import "jvmgo/rtdata"
//switch 跳转， case 连续的情况

/*
结构
<0-3 byte pad>
defaultbyte1
defaultbyte2
defaultbyte3
defaultbyte4
lowbyte1
lowbyte2
lowbyte3
lowbyte4
highbyte1
highbyte2
highbyte3
highbyte4
jump offsets...
*/
type TABLE_SWITCH struct {
	defaultOffset int32 //默认跳转的offset
	low           int32 //case取值范围下限
	high          int32 //case取值范围上限
	jumpOffsets   []int32 //各种case下跳转的offset
}

func (self *TABLE_SWITCH) FetchOperands(reader *instruction.BytecodeReader) {
	reader.SkipPadding() //tableswitch操作码后面有0-3位padding ，保证defaultOffset地址时4的倍数
	self.defaultOffset = reader.ReadInt32()
	self.low = reader.ReadInt32()
	self.high = reader.ReadInt32()
	jumpOffsetsCount := self.high - self.low + 1 //jumpOffsets 长度
	self.jumpOffsets = reader.ReadInt32s(jumpOffsetsCount)
}

func (self *TABLE_SWITCH) Execute(frame *rtdata.Frame) {
	index := frame.OperandStack().PopInt()

	var offset int
	if index >= self.low && index <= self.high { //是否在case的范围内
		offset = int(self.jumpOffsets[index-self.low]) //范围内的从 jumpOffsets 中拿取对应的offset
	} else { //范围外的直接用默认offset
		offset = int(self.defaultOffset)
	}

	instruction.Branch(frame, offset)
}
