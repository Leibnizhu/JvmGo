package base

import "jvmgo/rtdata"
import "jvmgo/rtdata/heap"

// jvms 5.5
func 	InitClass(thread *rtdata.Thread, class *heap.Class) {
	class.StartInit() //类对象的initStarted设为true，避免后面出现死循环
	scheduleClinit(thread, class)
	initSuperClass(thread, class) //如果父类没初始化要先递归初始化
}

func scheduleClinit(thread *rtdata.Thread, class *heap.Class) {
	clinit := class.GetClinitMethod() //获取clinit方法
	if clinit != nil {
		// exec <clinit>
		newFrame := thread.NewFrame(clinit)
		thread.PushFrame(newFrame) //推入新栈帧以执行clinit方法
	}
}

func initSuperClass(thread *rtdata.Thread, class *heap.Class) {
	if !class.IsInterface() {
		superClass := class.SuperClass()
		if superClass != nil && !superClass.InitStarted() { //父类未初始化
			InitClass(thread, superClass) //递归调用
		}
	}
}
