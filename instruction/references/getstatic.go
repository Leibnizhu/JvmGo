package references

import "jvmgo/instruction/base"
import "jvmgo/rtdata"
import "jvmgo/rtdata/heap"

// 读取类静态变量的值
//操作数是是uint16索引，为字段的符号引用
type GET_STATIC struct{ base.Index16Instruction }

func (self *GET_STATIC) Execute(frame *rtdata.Frame) {
	cp := frame.Method().Class().ConstantPool() //当前常量池
	fieldRef := cp.GetConstant(self.Index).(*heap.FieldRef) //从常量池读取字段引用
	field := fieldRef.ResolvedField() //解析静态字段引用
	class := field.Class() //静态字段对应的类的对象
	// 如果声明字段的类未初始化，需要初始化
	if !class.InitStarted() {
		frame.RevertNextPC()
		base.InitClass(frame.Thread(), class)
		return
	}

	if !field.IsStatic() { //非静态字段无法读取
		panic("java.lang.IncompatibleClassChangeError")
	}

	descriptor := field.Descriptor()
	slotId := field.SlotId() //字段在类中的slot编号
	slots := class.StaticVars()
	stack := frame.OperandStack()

	switch descriptor[0] { //根据静态字段的描述符，对应的类型，从类对象的slot中读取值，并放如如当前栈帧的作数栈中
	case 'Z', 'B', 'C', 'S', 'I':
		stack.PushInt(slots.GetInt(slotId))
	case 'F':
		stack.PushFloat(slots.GetFloat(slotId))
	case 'J':
		stack.PushLong(slots.GetLong(slotId))
	case 'D':
		stack.PushDouble(slots.GetDouble(slotId))
	case 'L', '[':
		stack.PushRef(slots.GetRef(slotId))
	default:
		// todo
	}
}
