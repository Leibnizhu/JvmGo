package references

import "reflect"
import "jvmgo/instruction/base"
import "jvmgo/rtdata"
import "jvmgo/rtdata/heap"

// 抛异常的指令 一个操作数,在操作数栈,为异常对象
type ATHROW struct{ base.NoOperandsInstruction }

func (self *ATHROW) Execute(frame *rtdata.Frame) {
	ex := frame.OperandStack().PopRef() //异常对象
	if ex == nil {
		panic("java.lang.NullPointerException")
	}

	thread := frame.Thread()
	if !findAndGotoExceptionHandler(thread, ex) { //找完当前线程所有栈帧了,还是找不到处理代码
		handleUncaughtException(thread, ex) //处理没有catch代码的异常
	}
}

//找异常处理代码并跳转
func findAndGotoExceptionHandler(thread *rtdata.Thread, ex *heap.Object) bool {
	for { //从当前栈帧一直往上找到能处理异常的代码
		frame := thread.CurrentFrame() //从当前栈帧开始遍历
		pc := frame.NextPC() - 1
		handlerPC := frame.Method().FindExceptionHandler(ex.Class(), pc) //从当前方法找处理的代码位置
		if handlerPC > 0 {                                               //找到了
			stack := frame.OperandStack()
			stack.Clear()              //清理操作数栈
			stack.PushRef(ex)          //异常对象入栈
			frame.SetNextPC(handlerPC) //调到catch块代码位置
			return true
		}

		thread.PopFrame()          //找不到则当前栈帧出栈,从上一个调用者去找处理代码
		if thread.IsStackEmpty() { //找完当前线程所有栈帧了,还是找不到处理代码
			break
		}
	}
	return false
}

func handleUncaughtException(thread *rtdata.Thread, ex *heap.Object) {
	thread.ClearStack() //清理jvm栈

	jMsg := ex.GetRefVar("detailMessage", "Ljava/lang/String;") //异常的详细信息
	goMsg := heap.GoString(jMsg)                                //异常信息转成go的字符串
	println(ex.Class().JavaName() + ": " + goMsg)

	stes := reflect.ValueOf(ex.Extra())
	for i := 0; i < stes.Len(); i++ {
		ste := stes.Index(i).Interface().(interface {
			String() string
		})
		println("\tat " + ste.String())
	}
}
