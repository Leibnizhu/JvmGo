package lang

import (
	"jvmgo/instruction/base"
	"jvmgo/native"
	"jvmgo/rtdata"
	"jvmgo/rtdata/heap"
	"strings"
)

const jlClass = "java/lang/Class"

func init() {
	native.Register(jlClass, "getPrimitiveClass", "(Ljava/lang/String;)Ljava/lang/Class;", getPrimitiveClass)
	native.Register(jlClass, "getName0", "()Ljava/lang/String;", getName0)
	native.Register(jlClass, "desiredAssertionStatus0", "(Ljava/lang/Class;)Z", desiredAssertionStatus0)
	native.Register(jlClass, "isInterface", "()Z", isInterface)
	native.Register(jlClass, "isPrimitive", "()Z", isPrimitive)
	native.Register(jlClass, "getDeclaredFields0", "(Z)[Ljava/lang/reflect/Field;", getDeclaredFields0)
	native.Register(jlClass, "forName0", "(Ljava/lang/String;ZLjava/lang/ClassLoader;Ljava/lang/Class;)Ljava/lang/Class;", forName0)
	native.Register(jlClass, "getDeclaredConstructors0", "(Z)[Ljava/lang/reflect/Constructor;", getDeclaredConstructors0)
	native.Register(jlClass, "getModifiers", "()I", getModifiers)
	native.Register(jlClass, "getSuperclass", "()Ljava/lang/Class;", getSuperclass)
	native.Register(jlClass, "getInterfaces0", "()[Ljava/lang/Class;", getInterfaces0)
	native.Register(jlClass, "isArray", "()Z", isArray)
	native.Register(jlClass, "getDeclaredMethods0", "(Z)[Ljava/lang/reflect/Method;", getDeclaredMethods0)
	native.Register(jlClass, "getComponentType", "()Ljava/lang/Class;", getComponentType)
	native.Register(jlClass, "isAssignableFrom", "(Ljava/lang/Class;)Z", isAssignableFrom)
}

//对应 static native Class<?> getPrimitiveClass(String name);
func getPrimitiveClass(frame *rtdata.Frame) {
	nameObj := frame.LocalVars().GetRef(0) //从局部变量表拿到类名
	name := heap.GoString(nameObj)         //转成go字符串

	loader := frame.Method().Class().Loader()
	class := loader.LoadClass(name).JClass()

	frame.OperandStack().PushRef(class)
}

//对应 private native String getName0();
func getName0(frame *rtdata.Frame) {
	this := frame.LocalVars().GetThis() //从局部变量表拿到this引用
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

// private static native Class<?> forName0(String name, boolean initialize, ClassLoader loader, Class<?> caller)
func forName0(frame *rtdata.Frame) {
	vars := frame.LocalVars()
	jName := vars.GetRef(0)
	initialize := vars.GetBoolean(1)
	//jLoader := vars.GetRef(2)

	goName := heap.GoString(jName)
	goName = strings.Replace(goName, ".", "/", -1)
	goClass := frame.Method().Class().Loader().LoadClass(goName)
	jClass := goClass.JClass()

	if initialize && !goClass.InitStarted() {
		// undo forName0
		thread := frame.Thread()
		frame.SetNextPC(thread.PC())
		// init class
		base.InitClass(thread, goClass)
	} else {
		stack := frame.OperandStack()
		stack.PushRef(jClass)
	}
}

// public native int getModifiers();
// ()I
func getModifiers(frame *rtdata.Frame) {
	vars := frame.LocalVars()
	this := vars.GetThis()
	class := this.Extra().(*heap.Class)
	modifiers := class.AccessFlags()

	stack := frame.OperandStack()
	stack.PushInt(int32(modifiers))
}

// public native Class<? super T> getSuperclass();
// ()Ljava/lang/Class;
func getSuperclass(frame *rtdata.Frame) {
	vars := frame.LocalVars()
	this := vars.GetThis()
	class := this.Extra().(*heap.Class)
	superClass := class.SuperClass()

	stack := frame.OperandStack()
	if superClass != nil {
		stack.PushRef(superClass.JClass())
	} else {
		stack.PushRef(nil)
	}
}

// private native Class<?>[] getInterfaces0();
// ()[Ljava/lang/Class;
func getInterfaces0(frame *rtdata.Frame) {
	vars := frame.LocalVars()
	this := vars.GetThis()
	class := this.Extra().(*heap.Class)
	interfaces := class.Interfaces()
	classArr := toClassArr(class.Loader(), interfaces)

	stack := frame.OperandStack()
	stack.PushRef(classArr)
}

// public native boolean isArray();
// ()Z
func isArray(frame *rtdata.Frame) {
	vars := frame.LocalVars()
	this := vars.GetThis()
	class := this.Extra().(*heap.Class)
	stack := frame.OperandStack()
	stack.PushBoolean(class.IsArray())
}

// public native Class<?> getComponentType();
// ()Ljava/lang/Class;
func getComponentType(frame *rtdata.Frame) {
	vars := frame.LocalVars()
	this := vars.GetThis()
	class := this.Extra().(*heap.Class)
	componentClass := class.ComponentClass()
	componentClassObj := componentClass.JClass()

	stack := frame.OperandStack()
	stack.PushRef(componentClassObj)
}

// public native boolean isAssignableFrom(Class<?> cls);
// (Ljava/lang/Class;)Z
func isAssignableFrom(frame *rtdata.Frame) {
	vars := frame.LocalVars()
	this := vars.GetThis()
	cls := vars.GetRef(1)

	thisClass := this.Extra().(*heap.Class)
	clsClass := cls.Extra().(*heap.Class)
	ok := thisClass.IsAssignableFrom(clsClass)

	stack := frame.OperandStack()
	stack.PushBoolean(ok)
}
