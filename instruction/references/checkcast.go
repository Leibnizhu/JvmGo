package references

import "jvmgo/instruction/base"
import "jvmgo/rtdata"
import "jvmgo/rtdata/heap"

// 检查是否可以强转为指定类/接口
//与instanceof类似，但instanceof会改变操作数栈（弹出引用，推入结果），
//而checkscast不改变操作数栈，转不了抛异常
type CHECK_CAST struct{ base.Index16Instruction }

func (self *CHECK_CAST) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	ref := stack.PopRef() //第二个操作数
	stack.PushRef(ref) //出栈之后马上入栈，不影响栈结构
	if ref == nil { //null的时候可以强转
		return
	}

	cp := frame.Method().Class().ConstantPool()
	classRef := cp.GetConstant(self.Index).(*heap.ClassRef) //从常量池获取目标类的符号引用
	class := classRef.ResolvedClass() //解析类
	if !ref.IsInstanceOf(class) { //如果不是其实例或没实现其接口，不能强转，抛异常
		panic("java.lang.ClassCastException")
	}
}
