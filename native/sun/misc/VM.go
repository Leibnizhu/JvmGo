package misc

import "jvmgo/instruction/base"
import "jvmgo/native"
import "jvmgo/rtdata"
import "jvmgo/rtdata/heap"

func init() {
	native.Register("sun/misc/VM", "initialize", "()V", initialize)
}

//对应 private static native void initialize();
func initialize(frame *rtdata.Frame) { // 临时处理,确保VM.savedProps不返回null,让自动装箱用到的IntegerCache(-128~127缓存)可以初始化
	vmClass := frame.Method().Class()
	savedProps := vmClass.GetRefVar("savedProps", "Ljava/util/Properties;") //VM的savedProps属性
	//setProperty方法用到的String参数
	key := heap.JString(vmClass.Loader(), "foo")
	val := heap.JString(vmClass.Loader(), "bar")
  //参数入操作数栈
	frame.OperandStack().PushRef(savedProps)
	frame.OperandStack().PushRef(key)
	frame.OperandStack().PushRef(val)
	//调用setProperty方法
	propsClass := vmClass.Loader().LoadClass("java/util/Properties")
	setPropMethod := propsClass.GetInstanceMethod("setProperty", "(Ljava/lang/String;Ljava/lang/String;)Ljava/lang/Object;")
	base.InvokeMethod(frame, setPropMethod)
}
