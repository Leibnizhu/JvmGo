package lang

import (
	"jvmgo/native"
	"jvmgo/rtdata"
	"jvmgo/rtdata/heap"
)

const jlClassLoader = "java/lang/ClassLoader"

func init() {
	native.Register(jlClassLoader, "findBuiltinLib", "(Ljava/lang/String;)Ljava/lang/String;", findBuiltinLib)
}

//FIXME 对应  private static native String findBuiltinLib(String name);
func findBuiltinLib(frame *rtdata.Frame) {
	vars := frame.LocalVars()
	libName := vars.GetRef(0)
	if libName == nil {
		panic("java.lang.NullPointerException")
	}
	goLibName := heap.GoString(libName)
	libraryName := "lib" + goLibName + ".so"
	cl := frame.Method().Class().Loader()
	stack := frame.OperandStack()
	stack.PushRef(heap.JString(cl, libraryName))
}
