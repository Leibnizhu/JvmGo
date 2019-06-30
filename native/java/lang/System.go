package lang

import "jvmgo/native"
import "jvmgo/rtdata"
import "jvmgo/rtdata/heap"

func init() {
	native.Register("java/lang/System", "arraycopy", "(Ljava/lang/Object;ILjava/lang/Object;II)V", arraycopy)
}

//对应 public static native void arraycopy(Object src, int srcPos, Object dest, int destPos, int length)
func arraycopy(frame *rtdata.Frame) {
	//从局部变量表拿到5个参数
	vars := frame.LocalVars()
	src := vars.GetRef(0)
	srcPos := vars.GetInt(1)
	dest := vars.GetRef(2)
	destPos := vars.GetInt(3)
	length := vars.GetInt(4)

	if src == nil || dest == nil {
		panic("java.lang.NullPointerException")
	}
	if !checkArrayCopy(src, dest) {
		panic("java.lang.ArrayStoreException")
	}
	//检查 srcPos destPos length 参数
	if srcPos < 0 || destPos < 0 || length < 0 ||
		srcPos+length > src.ArrayLength() ||
		destPos+length > dest.ArrayLength() {
		panic("java.lang.IndexOutOfBoundsException")
	}

	heap.ArrayCopy(src, dest, srcPos, destPos, length)
}

func checkArrayCopy(src, dest *heap.Object) bool {
	srcClass := src.Class()
	destClass := dest.Class()
	//检查src和dest都是数组
	if !srcClass.IsArray() || !destClass.IsArray() {
		return false
	}
	//检查数组类型:如果都是引用类型,可以拷贝,否则两者必须是同类型基础类型数组
	if srcClass.ComponentClass().IsPrimitive() || destClass.ComponentClass().IsPrimitive() {
		return srcClass == destClass
	}
	return true
}