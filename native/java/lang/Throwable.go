package lang

import (
	"fmt"
	"jvmgo/native"
	"jvmgo/rtdata"
	"jvmgo/rtdata/heap"
	"strconv"
)

// import "jvmgo/rtdata/heap"

const jlThrowable = "java/lang/Throwable"

type StackTraceElement struct {
	fileName   string //类所在文件名
	className  string //生命方法的类名
	methodName string //方法名
	lineNumber int    //正在执行的代码行数
}

func (self *StackTraceElement) String() string {
	var lineNumber string
	if self.lineNumber == -1 {
		lineNumber = "<Unknown>"
	} else if self.lineNumber == -2 {
		lineNumber = "<Native Method>"
	} else {
		lineNumber = strconv.Itoa(self.lineNumber)
	}
	return fmt.Sprintf("%s.%s(%s:%s)",
		self.className, self.methodName, self.fileName, lineNumber)
}

func init() {
	native.Register(jlThrowable, "fillInStackTrace", "(I)Ljava/lang/Throwable;", fillInStackTrace)
}

//对应 private native Throwable fillInStackTrace(int dummy); 抛异常会用到
func fillInStackTrace(frame *rtdata.Frame) {
	this := frame.LocalVars().GetThis()
	frame.OperandStack().PushRef(this)
	stes := createStackTraceElements(this, frame.Thread())
	this.SetExtra(stes)
}

//生成调用栈信息
func createStackTraceElements(tObj *heap.Object, thread *rtdata.Thread) []*StackTraceElement {
	skip := distanceToObject(tObj.Class()) + 2      //还要跳过 fillInStackTrace(int)和 fillInStackTrace() 两帧
	frames := thread.GetFrames()[skip:]             //GetFrames得到的栈帧数组是从栈顶开始的
	stes := make([]*StackTraceElement, len(frames)) //准备调用栈的信息数组
	for i, frame := range frames {
		stes[i] = createStackTraceElement(frame)
	}
	return stes
}

//因为fillInStackTrace是从实际的异常类一直调用到Throwable类,所以要计算异常类到Object还有多少层调用(即继承的层数)
func distanceToObject(class *heap.Class) int {
	distance := 0
	for c := class.SuperClass(); c != nil; c = c.SuperClass() { //找到Object为止
		distance++
	}
	return distance
}

//根据栈帧生成 StackTraceElement 对象
func createStackTraceElement(frame *rtdata.Frame) *StackTraceElement {
	method := frame.Method()
	class := method.Class()
	return &StackTraceElement{
		fileName:   class.SourceFile(),
		className:  class.JavaName(),
		methodName: method.Name(),
		lineNumber: method.GetLineNumber(frame.NextPC() - 1),
	}
}
