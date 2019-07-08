package misc

import "jvmgo/instruction/base"
import "jvmgo/native"
import "jvmgo/rtdata"

func init() {
	native.Register("sun/misc/VM", "initialize", "()V", initialize)
}

//对应 private static native void initialize();
func initialize(frame *rtdata.Frame) { // 临时处理,确保VM.savedProps不返回null,让自动装箱用到的IntegerCache(-128~127缓存)可以初始化
	classLoader := frame.Method().Class().Loader()
	jlSysClass := classLoader.LoadClass("java/lang/System")
	initSysClass := jlSysClass.GetStaticMethod("initializeSystemClass", "()V")
	base.InvokeMethod(frame, initSysClass)
}
