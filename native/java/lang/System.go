package lang

import (
	"jvmgo/instruction/base"
	"jvmgo/native"
	"jvmgo/rtdata"
	"jvmgo/rtdata/heap"
	"runtime"
	"time"
)

const jlSystem = "java/lang/System"

func init() {
	native.Register(jlSystem, "arraycopy", "(Ljava/lang/Object;ILjava/lang/Object;II)V", arraycopy)
	native.Register(jlSystem, "initProperties", "(Ljava/util/Properties;)Ljava/util/Properties;", initProperties)
	native.Register(jlSystem, "setIn0", "(Ljava/io/InputStream;)V", setIn0)
	native.Register(jlSystem, "setOut0", "(Ljava/io/PrintStream;)V", setOut0)
	native.Register(jlSystem, "setErr0", "(Ljava/io/PrintStream;)V", setErr0)
	native.Register(jlSystem, "mapLibraryName", "(Ljava/lang/String;)Ljava/lang/String;", mapLibraryName)
	native.Register(jlSystem, "currentTimeMillis", "()J", currentTimeMillis)
	native.Register(jlSystem, "nanoTime", "()J", nanoTime)
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

//对应  private static native Properties initProperties(Properties props);
func initProperties(frame *rtdata.Frame) {
	vars := frame.LocalVars()
	props := vars.GetRef(0)

	stack := frame.OperandStack()
	stack.PushRef(props)

	// public synchronized Object setProperty(String key, String value)
	setPropMethod := props.Class().GetInstanceMethod("setProperty", "(Ljava/lang/String;Ljava/lang/String;)Ljava/lang/Object;")
	thread := frame.Thread()
	for key, val := range _sysProps() {
		jKey := heap.JString(frame.Method().Class().Loader(), key)
		jVal := heap.JString(frame.Method().Class().Loader(), val)
		ops := rtdata.NewOperandStack(3)
		ops.PushRef(props)
		ops.PushRef(jKey)
		ops.PushRef(jVal)
		shimFrame := rtdata.NewShimFrame(thread, ops)
		thread.PushFrame(shimFrame)

		base.InvokeMethod(shimFrame, setPropMethod)
	}
}

//默认系统配置
func _sysProps() map[string]string {
	return map[string]string{
		"java.version":         "1.8.0",
		"java.vendor":          "jvm.go",
		"java.vendor.url":      "https://github.com/zxh0/jvm.go",
		"java.home":            "todo",
		"java.class.version":   "52.0",
		"java.class.path":      "todo",
		"java.awt.graphicsenv": "sun.awt.CGraphicsEnvironment",
		"os.name":              runtime.GOOS,   // todo
		"os.arch":              runtime.GOARCH, // todo
		"os.version":           "",             // todo
		"file.separator":       "/",            // todo os.PathSeparator
		"path.separator":       ":",            // todo os.PathListSeparator
		"line.separator":       "\n",           // todo
		"user.name":            "",             // todo
		"user.home":            "",             // todo
		"user.dir":             ".",            // todo
		"user.country":         "CN",           // todo
		"file.encoding":        "UTF-8",
		"sun.stdout.encoding":  "UTF-8",
		"sun.stderr.encoding":  "UTF-8",
	}
}

//对应  private static native void setIn0(InputStream in);
func setIn0(frame *rtdata.Frame) {
	vars := frame.LocalVars()
	in := vars.GetRef(0)
	sysClass := frame.Method().Class()
	sysClass.SetRefVar("in", "Ljava/io/InputStream;", in)
}

//对应  private static native void setOut0(PrintStream out);
func setOut0(frame *rtdata.Frame) {
	vars := frame.LocalVars()
	out := vars.GetRef(0)
	sysClass := frame.Method().Class()
	sysClass.SetRefVar("out", "Ljava/io/PrintStream;", out)
}

//对应  private static native void setErr0(PrintStream err);
func setErr0(frame *rtdata.Frame) {
	vars := frame.LocalVars()
	err := vars.GetRef(0)
	sysClass := frame.Method().Class()
	sysClass.SetRefVar("err", "Ljava/io/PrintStream;", err)
}

//对应  public static native long currentTimeMillis();
func currentTimeMillis(frame *rtdata.Frame) {
	millis := time.Now().UnixNano() / int64(time.Millisecond) //用go自带的算毫秒数
	stack := frame.OperandStack()
	stack.PushLong(millis)
}

//对应  public static native long nanoTime();
func nanoTime(frame *rtdata.Frame) {
	nanos := time.Now().UnixNano() //用go自带的纳秒数
	stack := frame.OperandStack()
	stack.PushLong(nanos)
}

//FIXME 对应  public static native String mapLibraryName(String libname);
func mapLibraryName(frame *rtdata.Frame) {
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
