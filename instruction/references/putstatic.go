package references

import "jvmgo/instruction/base"
import "jvmgo/rtdata"
import "jvmgo/rtdata/heap"

//给类的静态变量赋值
//第一个操作数是uint16索引，为字段的符号引用
//第二个操作数是付给静态变量的值
type PUT_STATIC struct{ base.Index16Instruction }

func (self *PUT_STATIC) Execute(frame *rtdata.Frame) {
	currentMethod := frame.Method() //当前方法
	currentClass := currentMethod.Class() //当前类
	cp := currentClass.ConstantPool() //当前常量池
	fieldRef := cp.GetConstant(self.Index).(*heap.FieldRef) //从常量池获取字段引用
	field := fieldRef.ResolvedField() //解析字段引用
	class := field.Class() //字段所在的类
	// 如果声明字段的类未初始化，需要初始化
	if !class.InitStarted() {
		frame.RevertNextPC()
		base.InitClass(frame.Thread(), class)
		return
	}

	if !field.IsStatic() { //非静态字段
		panic("java.lang.IncompatibleClassChangeError")
	}
	if field.IsFinal() { //final字段，无法赋值，只能在类初始化阶段赋值
		if currentClass != class || currentMethod.Name() != "<clinit>" {
			panic("java.lang.IllegalAccessError")
		}
	}

	descriptor := field.Descriptor()
	slotId := field.SlotId()
	slots := class.StaticVars() //类的静态变量
	stack := frame.OperandStack() //当前栈帧的操作数栈

	switch descriptor[0] { //根据字段描述符表示的类型，从当前操作栈出栈要赋值的值，同时设到静态变量的slot中
	case 'Z', 'B', 'C', 'S', 'I':
		slots.SetInt(slotId, stack.PopInt())
	case 'F':
		slots.SetFloat(slotId, stack.PopFloat())
	case 'J':
		slots.SetLong(slotId, stack.PopLong())
	case 'D':
		slots.SetDouble(slotId, stack.PopDouble())
	case 'L', '[':
		slots.SetRef(slotId, stack.PopRef())
	default:
		// todo
	}
}
