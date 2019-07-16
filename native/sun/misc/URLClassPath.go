package misc

import "jvmgo/native"
import "jvmgo/rtdata"

func init() {
	native.Register("sun/misc/URLClassPath", "getLookupCacheURLs", "(Ljava/lang/ClassLoader;)[Ljava/net/URL;", getLookupCacheURLs)
}

//对应 private static native URL[] getLookupCacheURLs(ClassLoader var0);
func getLookupCacheURLs(frame *rtdata.Frame) {
	frame.OperandStack().PushRef(nil) //临时实现
}
