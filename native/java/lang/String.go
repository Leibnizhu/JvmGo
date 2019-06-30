package lang

import "jvmgo/native"
import "jvmgo/rtdata"
import "jvmgo/rtdata/heap"

const jlString = "java/lang/String"

func init() {
	native.Register(jlString, "intern", "()Ljava/lang/String;", intern)
}

//对应 public native String intern();
func intern(frame *rtdata.Frame) {
	this := frame.LocalVars().GetThis()
	interned := heap.InternString(this)
	frame.OperandStack().PushRef(interned)
}
