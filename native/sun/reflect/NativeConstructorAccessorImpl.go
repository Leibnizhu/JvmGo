package reflect

import "jvmgo/instruction/base"
import "jvmgo/native"
import "jvmgo/rtdata"
import "jvmgo/rtdata/heap"

func init() {
	native.Register("sun/reflect/NativeConstructorAccessorImpl", "newInstance0", "(Ljava/lang/reflect/Constructor;[Ljava/lang/Object;)Ljava/lang/Object;", newInstance0)
}

//对应 private static native Object newInstance0(Constructor<?> c, Object[] os)
func newInstance0(frame *rtdata.Frame) {
	vars := frame.LocalVars()
	constructorObj := vars.GetRef(0)
	argArrObj := vars.GetRef(1)

	goConstructor := getGoConstructor(constructorObj)
	goClass := goConstructor.Class()
	if !goClass.InitStarted() {
		frame.RevertNextPC()
		base.InitClass(frame.Thread(), goClass)
		return
	}

	obj := goClass.NewObject()
	stack := frame.OperandStack()
	stack.PushRef(obj)

	// call <init>
	ops := convertArgs(obj, argArrObj, goConstructor)
	shimFrame := rtdata.NewShimFrame(frame.Thread(), ops)
	frame.Thread().PushFrame(shimFrame)

	base.InvokeMethod(shimFrame, goConstructor)
}

func getGoMethod(methodObj *heap.Object) *heap.Method {
	return _getGoMethod(methodObj, false)
}
func getGoConstructor(constructorObj *heap.Object) *heap.Method {
	return _getGoMethod(constructorObj, true)
}
func _getGoMethod(methodObj *heap.Object, isConstructor bool) *heap.Method {
	extra := methodObj.Extra()
	if extra != nil {
		return extra.(*heap.Method)
	}

	if isConstructor {
		root := methodObj.GetRefVar("root", "Ljava/lang/reflect/Constructor;")
		return root.Extra().(*heap.Method)
	} else {
		root := methodObj.GetRefVar("root", "Ljava/lang/reflect/Method;")
		return root.Extra().(*heap.Method)
	}
}

// Object[] -> []interface{}
func convertArgs(this, argArr *heap.Object, method *heap.Method) *rtdata.OperandStack {
	if method.ArgSlotCount() == 0 {
		return nil
	}

	//	argObjs := argArr.Refs()
	//	argTypes := method.ParsedDescriptor().ParameterTypes()

	ops := rtdata.NewOperandStack(method.ArgSlotCount())
	if !method.IsStatic() {
		ops.PushRef(this)
	}
	if method.ArgSlotCount() == 1 && !method.IsStatic() {
		return ops
	}

	//	for i, argType := range argTypes {
	//		argObj := argObjs[i]
	//
	//		if len(argType) == 1 {
	//			// base type
	//			// todo
	//			unboxed := box.Unbox(argObj, argType)
	//			args[i+j] = unboxed
	//			if argType.isLongOrDouble() {
	//				j++
	//			}
	//		} else {
	//			args[i+j] = argObj
	//		}
	//	}

	return ops
}
