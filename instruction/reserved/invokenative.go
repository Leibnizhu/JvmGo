package reserved

import "jvmgo/instruction/base"
import "jvmgo/rtdata"
import "jvmgo/native"
import _ "jvmgo/native/java/lang"
import _ "jvmgo/native/java/security"
import _ "jvmgo/native/java/io"
import _ "jvmgo/native/java/util/concurrent/atomic"
import _ "jvmgo/native/sun/io"
import _ "jvmgo/native/sun/misc"
import _ "jvmgo/native/sun/reflect"

// 调用本地方法指令，不需要操作数
type INVOKE_NATIVE struct{ base.NoOperandsInstruction }

//根据类名，方法名，方法描述符从本地方法注册表查找本地方法
func (self *INVOKE_NATIVE) Execute(frame *rtdata.Frame) {
	method := frame.Method()
	className := method.Class().Name()
	methodName := method.Name()
	methodDescriptor := method.Descriptor()

	nativeMethod := native.FindNativeMethod(className, methodName, methodDescriptor)
	if nativeMethod == nil {
		methodInfo := className + "." + methodName + methodDescriptor
		panic("java.lang.UnsatisfiedLinkError: " + methodInfo) //从本地方法注册表找不到本地方法，抛异常
	}

	nativeMethod(frame)
}
