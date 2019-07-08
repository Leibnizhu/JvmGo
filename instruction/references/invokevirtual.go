package references

import "jvmgo/instruction/base"
import "jvmgo/rtdata"
import "jvmgo/rtdata/heap"

// 调用实例方法，基于类分发
type INVOKE_VIRTUAL struct{ base.Index16Instruction }

func (self *INVOKE_VIRTUAL) Execute(frame *rtdata.Frame) {
	currentClass := frame.Method().Class()
	cp := currentClass.ConstantPool()
	methodRef := cp.GetConstant(self.Index).(*heap.MethodRef)
	resolvedMethod := methodRef.ResolvedMethod()
	if resolvedMethod.IsStatic() { //当前指令不能执行静态方法
		panic("java.lang.IncompatibleClassChangeError")
	}
	//拿到this
	ref := frame.OperandStack().GetRefFromTop(resolvedMethod.ArgSlotCount() - 1)
	if ref == nil {
		panic("java.lang.NullPointerException")
	}
	//确保protected方法只能被当前类或子类调用
	if resolvedMethod.IsProtected() &&
		resolvedMethod.Class().IsSuperClassOf(currentClass) &&
		resolvedMethod.Class().GetPackageName() != currentClass.GetPackageName() &&
		ref.Class() != currentClass &&
		!ref.Class().IsSubClassOf(currentClass) {

		panic("java.lang.IllegalAccessError")
	}
	//根据对象的类来查找真正要执行的方法，这里跟 invokespecial 指令不一样，后者直接使用方法引用对应的方法
	methodToBeInvoked := heap.LookupMethodInClass(ref.Class(), methodRef.Name(), methodRef.Descriptor())
	if methodToBeInvoked == nil || methodToBeInvoked.IsAbstract() {
		panic("java.lang.AbstractMethodError")
	}

	base.InvokeMethod(frame, methodToBeInvoked)
}
