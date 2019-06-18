package native

import "jvmgo/rtdata"

//本地方法， 定义为一个函数参数是Frame指针，即本地方法的工作空间
type NativeMethod func(frame *rtdata.Frame)

//本地方法注册表，是个hash表，key是 类名+方法名+方法签名
var registry = map[string]NativeMethod{}

//本地方法的空实现，不做任何操作
func emptyNativeMethod(frame *rtdata.Frame) {
}

//注册本地方法，翻入hash表
func Register(className, methodName, methodDescriptor string, method NativeMethod) {
	key := className + "~" + methodName + "~" + methodDescriptor
	registry[key] = method
}

//从本地方法注册表查找方法
func FindNativeMethod(className, methodName, methodDescriptor string) NativeMethod {
	key := className + "~" + methodName + "~" + methodDescriptor
	if method, ok := registry[key]; ok {
		return method
	}
	//对于registerNatives方法，Object等类通过它注册其他本地方法，在这里没什么用，返回空实现
	if methodDescriptor == "()V" && methodName == "registerNatives" {
		return emptyNativeMethod
	}
	return nil //注册表里找不到，又不是registerNatives，则返回nil
}
