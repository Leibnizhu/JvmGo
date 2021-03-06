package main

import "fmt"
import "jvmgo/instruction"
import "jvmgo/instruction/base"
import "jvmgo/rtdata"

//解释器
//logInst : 是否打印执行信息到控制台
func interpret(thread *rtdata.Thread, logInst bool) {
	defer catchErr(thread) //return 诗回调
	loop(thread, logInst)
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

		reader.Reset(frame.Method().Code(), pc)    //重置reader
		opcode := reader.ReadUint8()               //读取操作指令
		inst := instruction.NewInstruction(opcode) //从工厂方法获取对应指令对象
		inst.FetchOperands(reader)                 //指令读取操作数
		frame.SetNextPC(reader.PC())               //设置下一个指令地址
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
