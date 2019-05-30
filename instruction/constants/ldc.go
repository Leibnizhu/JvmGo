package constants

import "jvmgo/instruction/base"
import "jvmgo/rtdata"
//LDC系列指令，从常量池读取常量值，并推入ca操作数栈

//加载int float 字符串常量 Class实例
//ldc 和 ldc_w 区别仅在于操作数宽度
type LDC struct{ base.Index8Instruction }

func (self *LDC) Execute(frame *rtdata.Frame) {
	_ldc(frame, self.Index)
}

type LDC_W struct{ base.Index16Instruction }

func (self *LDC_W) Execute(frame *rtdata.Frame) {
	_ldc(frame, self.Index)
}

//LDC 和 LDC_W 实际逻辑
func _ldc(frame *rtdata.Frame, index uint) {
	stack := frame.OperandStack()
	cp := frame.Method().Class().ConstantPool()
	c := cp.GetConstant(index) //从当前类的常量池获取常量

	switch c.(type) {
	//如果是int或float，入栈
	case int32:
		stack.PushInt(c.(int32))
	case float32:
		stack.PushFloat(c.(float32))
	// case string:
	// case *heap.ClassRef:
	// case MethodType, MethodHandle
	default: //其他暂不支持
		panic("todo: ldc!")
	}
}

// 加载 long 和 double，与 _ldc() 类似
type LDC2_W struct{ base.Index16Instruction }

func (self *LDC2_W) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	cp := frame.Method().Class().ConstantPool()
	c := cp.GetConstant(self.Index)

	switch c.(type) {
	case int64:
		stack.PushLong(c.(int64))
	case float64:
		stack.PushDouble(c.(float64))
	default:
		panic("java.lang.ClassFormatError")
	}
}
