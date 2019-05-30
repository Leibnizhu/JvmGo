package main

import "fmt"
import "jvmgo/instruction"
import "jvmgo/instruction/base"
import "jvmgo/rtdata"
import "jvmgo/rtdata/heap"
//解释器

func interpret(method *heap.Method) {
	thread := rtdata.NewThread() //新建线程
	frame := thread.NewFrame(method) //新建线程栈帧
	thread.PushFrame(frame) //栈帧入栈
	defer catchErr(frame) //return 诗回调
	loop(thread, method.Code())
}

func catchErr(frame *rtdata.Frame) {
	if r := recover(); r != nil { //如果有panic， recover() 拿到异常
		fmt.Printf("LocalVars:%v\n", frame.LocalVars())
		fmt.Printf("OperandStack:%v\n", frame.OperandStack())
		panic(r)
	}
}

func loop(thread *rtdata.Thread, bytecode []byte) {
	frame := thread.PopFrame() //当前栈帧
	reader := &base.BytecodeReader{}

	for {
		pc := frame.NextPC() //下一个指令地址
		thread.SetPC(pc)

		reader.Reset(bytecode, pc) //重置reader
		opcode := reader.ReadUint8() //读取操作指令
		inst := instruction.NewInstruction(opcode) //从工厂方法获取对应指令对象
		inst.FetchOperands(reader) //指令读取操作数
		frame.SetNextPC(reader.PC()) //设置下一个指令地址

		fmt.Printf("pc:%2d inst:%T %v\n", pc, inst, inst)
		inst.Execute(frame) //执行指令
	}
}