package references

import "jvmgo/instruction/base"
import "jvmgo/rtdata"
import "jvmgo/rtdata/heap"

//new 指令，创建新的实例
//操作数是uint16索引
type NEW struct{ base.Index16Instruction }

func (self *NEW) Execute(frame *rtdata.Frame) {
	cp := frame.Method().Class().ConstantPool()
	classRef := cp.GetConstant(self.Index).(*heap.ClassRef) //从常量池找到一个类符号引用
	class := classRef.ResolvedClass() //拿到类数据，可能会加载
	// init class
	if !class.InitStarted() {
		frame.RevertNextPC()
		base.InitClass(frame.Thread(), class)
		return
	}

	if class.IsInterface() || class.IsAbstract() { //抽象类和接口无法实例化
		panic("java.lang.InstantiationError") //JVM规范规定的异常
	}

	ref := class.NewObject() //创建对象
	frame.OperandStack().PushRef(ref) //放入操作数栈顶
}
