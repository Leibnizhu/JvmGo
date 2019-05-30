package references

import "jvmgo/instruction/base"
import "jvmgo/rtdata"

//new 指令后吗会有 invokesprcial 指令，来调用构造函数初始化对象
// special handling for superclass, private, and instance initialization method invocations
type INVOKE_SPECIAL struct{ base.Index16Instruction }

// 临时处理，以后要真正实现
func (self *INVOKE_SPECIAL) Execute(frame *rtdata.Frame) {
	frame.OperandStack().PopRef() //从操作数栈出栈一个引用，不调构造函数了
}
