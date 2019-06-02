package references

import "jvmgo/instruction/base"
import "jvmgo/rtdata"
import "jvmgo/rtdata/heap"

//new 指令后会有 invokesprcial 指令，来调用构造函数初始化对象
//调用实例方法，（静态绑定）包括实例初始化方法、私有方法、super关键字调用的方法
//因为私有方法和构造方法不需要动态绑定，提高速度；而super的方法如果用invokevirtual会死循环
type INVOKE_SPECIAL struct{ base.Index16Instruction }

// 临时处理，以后要真正实现
func (self *INVOKE_SPECIAL) Execute(frame *rtdata.Frame) {
	currentClass := frame.Method().Class() //当前类
	cp := currentClass.ConstantPool() //当前常量池
	methodRef := cp.GetConstant(self.Index).(*heap.MethodRef) //方法的符号引用
	resolvedClass := methodRef.ResolvedClass() //解析符号引用解析后的类
	resolvedMethod := methodRef.ResolvedMethod()

	if resolvedMethod.Name() == "<init>" && resolvedMethod.Class() != resolvedClass { //解析出来的方法是构造函数，那么对应的类必须是解析出来的类
		panic("java.lang.NoSuchMethodError")
	}
	if resolvedMethod.IsStatic() { //静态方法不能用当前指令执行
		panic("java.lang.IncompatibleClassChangeError")
	}
	//这时候要拿this，但在传递参数之前，不能破坏操作数栈的状态，所以不能出栈，新增了 GetRefFromTop 获取离栈顶n个位置的引用变量
	ref := frame.OperandStack().GetRefFromTop(resolvedMethod.ArgSlotCount() - 1)
	if ref == nil { //this为null，抛NPE
		panic("java.lang.NullPointerException")
	}
	//protected方法只能被不是声明方法的类，也不是其子类调用，不可访问
	if resolvedMethod.IsProtected() && //protected方法
		resolvedMethod.Class().IsSuperClassOf(currentClass) && //方法的类是当前类的父类
		resolvedMethod.Class().GetPackageName() != currentClass.GetPackageName() && //方法类和当前类不同包
		ref.Class() != currentClass && //this类不是当前类
		!ref.Class().IsSubClassOf(currentClass) { //this类也不是当前类的子类
		panic("java.lang.IllegalAccessError")
	}

	methodToBeInvoked := resolvedMethod
	//如果调用父类方法（且不是父类初始化方法），那么需要从父类查找调用的方法
	if currentClass.IsSuper() &&
		resolvedClass.IsSuperClassOf(currentClass) && //的确是调用了父类
		resolvedMethod.Name() != "<init>" {
		methodToBeInvoked = heap.LookupMethodInClass(currentClass.SuperClass(), methodRef.Name(), methodRef.Descriptor()）
	}

	//找不到方法，或者方法为抽象方法
	if methodToBeInvoked == nil || methodToBeInvoked.IsAbstract() {
		panic("java.lang.AbstractMethodError")
	}

	base.InvokeMethod(frame, methodToBeInvoked)
}