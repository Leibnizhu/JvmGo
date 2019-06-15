package main

import "fmt"
import "jvmgo/instruction"
import "jvmgo/instruction/base"
import "jvmgo/rtdata"
import "jvmgo/rtdata/heap"
//解释器

//logInst : 是否打印执行信息到控制台
func interpret(method *heap.Method, logInst bool, args []string) {
	thread := rtdata.NewThread() //新建线程
	frame := thread.NewFrame(method) //新建线程栈帧
	thread.PushFrame(frame) //栈帧入栈
	jArg := createArgsArray(method.Class().Loader(), args) //根据启动参数生成字符串数组
	frame.LocalVars().SetRef(0, jArg) //启动参数放入操作数栈顶
	defer catchErr(thread) //return 诗回调
	loop(thread, logInst)
}

func createArgsArray(loader *heap.ClassLoader, args []string) *heap.Object {
	stringClass := loader.LoadClass("java/lang/String")
	argsArr := stringClass.ArrayClass().NewArray(uint(len(args)))
	jArgs := argsArr.Refs()
	for i, arg := range args {
		jArgs[i] = heap.JString(loader, arg)
	}
	return argsArr
}

func catchErr(thread *rtdata.Thread) {
	if r := recover(); r != nil {
		logFrames(thread)
		panic(r)
	}
}

func loop(thread *rtdata.Thread, logInst bool) {
	reader := &base.BytecodeReader{}

	for {
		frame := thread.CurrentFrame()
		pc := frame.NextPC() //下一个指令地址
		thread.SetPC(pc)

		reader.Reset(frame.Method().Code(), pc) //重置reader
		opcode := reader.ReadUint8() //读取操作指令
		inst := instruction.NewInstruction(opcode) //从工厂方法获取对应指令对象
		inst.FetchOperands(reader) //指令读取操作数
		frame.SetNextPC(reader.PC()) //设置下一个指令地址
		if logInst {
			logInstruction(frame, inst)
		}
		//执行代码
		inst.Execute(frame)
		if thread.IsStackEmpty() {
			break
		}
	}
}

func logInstruction(frame *rtdata.Frame, inst base.Instruction) {
	method := frame.Method()
	className := method.Class().Name()
	methodName := method.Name()
	pc := frame.Thread().PC()
	fmt.Printf("%v.%v() #%2d %T %v\n", className, methodName, pc, inst, inst)
}

func logFrames(thread *rtdata.Thread) {
	for !thread.IsStackEmpty() {
		frame := thread.PopFrame()
		method := frame.Method()
		className := method.Class().Name()
		fmt.Printf(">> pc:%4d %v.%v%v \n",
			frame.NextPC(), className, method.Name(), method.Descriptor())
	}
}