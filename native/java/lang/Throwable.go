package lang

import "fmt"
import "jvmgo/native"
import "jvmgo/rtdata"

// import "jvmgo/rtdata/heap"

const jlThrowable = "java/lang/Throwable"

type StackTraceElement struct {
	fileName   string
	className  string
	methodName string
	lineNumber int
}

func (self *StackTraceElement) String() string {
	return fmt.Sprintf("%s.%s(%s:%d)",
		self.className, self.methodName, self.fileName, self.lineNumber)
}

func init() {
	native.Register(jlThrowable, "fillInStackTrace", "(I)Ljava/lang/Throwable;", fillInStackTrace)
}

//对应 private native Throwable fillInStackTrace(int dummy); 抛异常会用到,暂时空实现
func fillInStackTrace(frame *rtdata.Frame) {
	// this := frame.LocalVars().GetThis()
	// frame.OperandStack().PushRef(this)

	// stes := createStackTraceElements(this, frame.Thread())
	// this.SetExtra(stes)
}

// func createStackTraceElements(tObj *heap.Object, thread *rtdata.Thread) []*StackTraceElement {
// 	skip := distanceToObject(tObj.Class()) + 2
// 	frames := thread.GetFrames()[skip:]
// 	stes := make([]*StackTraceElement, len(frames))
// 	for i, frame := range frames {
// 		stes[i] = createStackTraceElement(frame)
// 	}
// 	return stes
// }

// func distanceToObject(class *heap.Class) int {
// 	distance := 0
// 	for c := class.SuperClass(); c != nil; c = c.SuperClass() {
// 		distance++
// 	}
// 	return distance
// }

// func createStackTraceElement(frame *rtda.Frame) *StackTraceElement {
// 	method := frame.Method()
// 	class := method.Class()
// 	return &StackTraceElement{
// 		fileName:   class.SourceFile(),
// 		className:  class.JavaName(),
// 		methodName: method.Name(),
// 		lineNumber: method.GetLineNumber(frame.NextPC() - 1),
// 	}
// }
