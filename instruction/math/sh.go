package math

import "jvmgo/instruction/base"
import "jvmgo/rtdata"
//左移右移系列指令

// int 左移
type ISHL struct{ base.NoOperandsInstruction  }

func (self *ISHL) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	s := uint32(v2) & 0x1f //取后5位即可 (int 32位， 5bit以上的，重复移动，没必要了)
	result := v1 << s
	stack.PushInt(result)
}

// int 有符号右移
type ISHR struct{ base.NoOperandsInstruction  }

func (self *ISHR) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	s := uint32(v2) & 0x1f //go 位操作必须是无符号数
	result := v1 >> s
	stack.PushInt(result)
}

// int 无符号右移
type IUSHR struct{ base.NoOperandsInstruction  }

func (self *IUSHR) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	s := uint32(v2) & 0x1f
	result := int32(uint32(v1) >> s) //无符号数右移后，再转有符号
	stack.PushInt(result)
}

// long 左移
type LSHL struct{ base.NoOperandsInstruction  }

func (self *LSHL) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopLong()
	s := uint32(v2) & 0x3f
	result := v1 << s
	stack.PushLong(result)
}

// long 有符号右移
type LSHR struct{ base.NoOperandsInstruction  }

func (self *LSHR) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopLong()
	s := uint32(v2) & 0x3f
	result := v1 >> s
	stack.PushLong(result)
}

// long 无符号右移
type LUSHR struct{ base.NoOperandsInstruction  }

func (self *LUSHR) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopLong()
	s := uint32(v2) & 0x3f
	result := int64(uint64(v1) >> s)
	stack.PushLong(result)
}
