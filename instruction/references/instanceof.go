package references

import "jvmgo/instruction/base"
import "jvmgo/rtdata"
import "jvmgo/rtdata/heap"

// 判断d对象是否某个类的实例或实现了某个接口，2个操作数
//第一个操作数是uint16索引，指向判定的类或接口的符号引用
//第二个操作数是要进行判断的对象引用
//否则0入栈，是则1入栈
type INSTANCE_OF struct{ base.Index16Instruction }

func (self *INSTANCE_OF) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	ref := stack.PopRef() //第二个操作数
	if ref == nil { //null instanceof 永远是false，0入栈
		stack.PushInt(0)
		return
	}

	cp := frame.Method().Class().ConstantPool() //当前常量池
	classRef := cp.GetConstant(self.Index).(*heap.ClassRef) //从常量池根据索引获取类或接口的符号引用
	class := classRef.ResolvedClass() //加载类
	if ref.IsInstanceOf(class) {
		stack.PushInt(1)
	} else {
		stack.PushInt(0)
	}
}
