package lang

import "jvmgo/native"
import "jvmgo/rtdata"
import "unsafe"

const jlObject = "java/lang/Object"

func init(){
	native.Register(jlObject, "getClass", "()Ljava/lang/Class;", getClass)
	native.Register(jlObject, "hashCode", "()I", hashCode)
}

//对应  public final native Class<?> getClass() 方法
func getClass(frame *rtdata.Frame){
	this := frame.LocalVars().GetThis() //从局部变量表拿到this引用
	class := this.Class().JClass() //拿到类对象
	frame.OperandStack().PushRef(class) //如操作数栈
}

//对应 public native int hashCode();
func hashCode(frame *rtdata.Frame) {
	this := frame.LocalVars().GetThis()
	hash := int32(uintptr(unsafe.Pointer(this))) //拿到Object结构体指针,转unintptr类型,再强转int32入栈
	frame.OperandStack().PushInt(hash)
}

