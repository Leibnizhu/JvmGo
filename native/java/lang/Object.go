package lang

import "jvmgo/native"
import "jvmgo/rtdata"
import "unsafe"

const jlObject = "java/lang/Object"

func init() {
	native.Register(jlObject, "getClass", "()Ljava/lang/Class;", getClass)
	native.Register(jlObject, "hashCode", "()I", hashCode)
	native.Register(jlObject, "clone", "()Ljava/lang/Object;", clone)
	native.Register(jlObject, "notifyAll", "()V", notifyAll)
}

//对应  public final native Class<?> getClass() 方法
func getClass(frame *rtdata.Frame) {
	this := frame.LocalVars().GetThis() //从局部变量表拿到this引用
	class := this.Class().JClass()      //拿到类对象
	frame.OperandStack().PushRef(class) //如操作数栈
}

//对应 public native int hashCode();
func hashCode(frame *rtdata.Frame) {
	this := frame.LocalVars().GetThis()
	hash := int32(uintptr(unsafe.Pointer(this))) //拿到Object结构体指针,转unintptr类型,再强转int32入栈
	frame.OperandStack().PushInt(hash)
}

//对应 protected native Object clone() throws CloneNotSupportedException;
func clone(frame *rtdata.Frame) {
	this := frame.LocalVars().GetThis()
	cloneable := this.Class().Loader().LoadClass("java/lang/Cloneable")
	if !this.Class().IsImplements(cloneable) { //this引用对应的对象类未实现Cloneable接口
		panic("java.lang.CloneNotSupportedException")
	}
	frame.OperandStack().PushRef(this.Clone())
}

//对应 public final native void notifyAll();
func notifyAll(frame *rtdata.Frame) {
	// todo
}
