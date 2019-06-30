package lang

import "jvmgo/native"
import "jvmgo/rtdata"

func init(){
	native.Register("java/lang/Object", "getClass", "()Ljava/lang/Class;", getClass)
}

//对应  public final native Class<?> getClass() 方法
func getClass(frame *rtdata.Frame){
	this := frame.LocalVars().GetThis() //从局部变量表拿到this引用
	class := this.Class().JClass() //拿到类对象
	frame.OperandStack().PushRef(class) //如操作数栈
}
