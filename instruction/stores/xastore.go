package stores

import "jvmgo/instruction/base"
import "jvmgo/rtdata"
import "jvmgo/rtdata/heap"
//<t>astore系列指令，根据索引设置数组的值，
//3个操作数，都在操作数栈，分别是：
//赋值的值，数组索引，数组引用

// 对象数组赋值
type AASTORE struct{ base.NoOperandsInstruction }

func (self *AASTORE) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	ref := stack.PopRef() //赋值的值的引用
	index := stack.PopInt() //数组下标/索引
	arrRef := stack.PopRef() //数组引用

	checkNotNil(arrRef) //检查数组非null
	refs := arrRef.Refs()
	checkIndex(len(refs), index) //检查下标越界
	refs[index] = ref
}

//   byte 或 boolean 数组赋值
type BASTORE struct{ base.NoOperandsInstruction }

func (self *BASTORE) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	val := stack.PopInt()
	index := stack.PopInt()
	arrRef := stack.PopRef()

	checkNotNil(arrRef)
	bytes := arrRef.Bytes()
	checkIndex(len(bytes), index)
	bytes[index] = int8(val)
}

// char  数组赋值
type CASTORE struct{ base.NoOperandsInstruction }

func (self *CASTORE) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	val := stack.PopInt()
	index := stack.PopInt()
	arrRef := stack.PopRef()

	checkNotNil(arrRef)
	chars := arrRef.Chars()
	checkIndex(len(chars), index)
	chars[index] = uint16(val)
}

// double  数组赋值
type DASTORE struct{ base.NoOperandsInstruction }

func (self *DASTORE) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	val := stack.PopDouble()
	index := stack.PopInt()
	arrRef := stack.PopRef()

	checkNotNil(arrRef)
	doubles := arrRef.Doubles()
	checkIndex(len(doubles), index)
	doubles[index] = float64(val)
}

// float  数组赋值
type FASTORE struct{ base.NoOperandsInstruction }

func (self *FASTORE) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	val := stack.PopFloat()
	index := stack.PopInt()
	arrRef := stack.PopRef()

	checkNotNil(arrRef)
	floats := arrRef.Floats()
	checkIndex(len(floats), index)
	floats[index] = float32(val)
}

// int  数组赋值
type IASTORE struct{ base.NoOperandsInstruction }

func (self *IASTORE) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	val := stack.PopInt()
	index := stack.PopInt()
	arrRef := stack.PopRef()

	checkNotNil(arrRef)
	ints := arrRef.Ints()
	checkIndex(len(ints), index)
	ints[index] = int32(val)
}

// long  数组赋值
type LASTORE struct{ base.NoOperandsInstruction }

func (self *LASTORE) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	val := stack.PopLong()
	index := stack.PopInt()
	arrRef := stack.PopRef()

	checkNotNil(arrRef)
	longs := arrRef.Longs()
	checkIndex(len(longs), index)
	longs[index] = int64(val)
}

// short  数组赋值
type SASTORE struct{ base.NoOperandsInstruction }

func (self *SASTORE) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	val := stack.PopInt()
	index := stack.PopInt()
	arrRef := stack.PopRef()

	checkNotNil(arrRef)
	shorts := arrRef.Shorts()
	checkIndex(len(shorts), index)
	shorts[index] = int16(val)
}

func checkNotNil(ref *heap.Object) {
	if ref == nil {
		panic("java.lang.NullPointerException")
	}
}
func checkIndex(arrLen int, index int32) {
	if index < 0 || index >= int32(arrLen) {
		panic("ArrayIndexOutOfBoundsException")
	}
}
