package lang


import "jvmgo/native"
import "jvmgo/rtdata"
import "jvmgo/rtdata/heap"

const jlClass = "java/lang/Class"

func init() {
	native.Register(jlClass, "getPrimitiveClass", "(Ljava/lang/String;)Ljava/lang/Class;", getPrimitiveClass)
	native.Register(jlClass, "getName0", "()Ljava/lang/String;", getName0)
	native.Register(jlClass, "desiredAssertionStatus0", "(Ljava/lang/Class;)Z", desiredAssertionStatus0)
	native.Register(jlClass, "isInterface", "()Z", isInterface)
	native.Register(jlClass, "isPrimitive", "()Z", isPrimitive)
}
//对应 static native Class<?> getPrimitiveClass(String name);
func getPrimitiveClass(frame *rtdata.Frame) {
	nameObj := frame.LocalVars().GetRef(0) //从局部变量表拿到类名
	name := heap.GoString(nameObj) //转成go字符串

	loader := frame.Method().Class().Loader()
	class := loader.LoadClass(name).JClass()

	frame.OperandStack().PushRef(class)
}

//对应 private native String getName0();
func getName0(frame *rtdata.Frame) {
	this := frame.LocalVars().GetThis()//从局部变量表拿到this引用
	class := this.Extra().(*heap.Class)
	name := class.JavaName()
	nameObj := heap.JString(class.Loader(), name)

	frame.OperandStack().PushRef(nameObj)
}

//对应 private static native boolean desiredAssertionStatus0(Class<?> clazz);
func desiredAssertionStatus0(frame *rtdata.Frame) {
	// todo　暂不处理断言
	frame.OperandStack().PushBoolean(false)
}

//对应 public native boolean isInterface();
func isInterface(frame *rtdata.Frame) {
	this := frame.LocalVars().GetThis()
	class := this.Extra().(*heap.Class)
	frame.OperandStack().PushBoolean(class.IsInterface())
}

//对应 public native boolean isPrimitive();
func isPrimitive(frame *rtdata.Frame) {
	this := frame.LocalVars().GetThis()
	class := this.Extra().(*heap.Class)
	frame.OperandStack().PushBoolean(class.IsPrimitive())
}
