package main

import "fmt"
import "strings"
import "jvmgo/classpath"
import "jvmgo/instruction/base"
import "jvmgo/rtdata"
import "jvmgo/rtdata/heap"

//入口
type JVM struct {
	cmd         *Cmd
	classLoader *heap.ClassLoader
	mainThread  *rtdata.Thread
}

func newJVM(cmd *Cmd) *JVM {
	cp := classpath.Parse(cmd.XjreOption, cmd.cpOption)
	fmt.Printf("classpath:%s class:%s args:%s\n", cp, cmd.class, cmd.args)
	className := strings.Replace(cmd.class, ".", "/", -1) //类名改成类路径
	classfile := loadClass(className, cp)
	fmt.Println(cmd.class)
	printClassInfo(classfile)
	classLoader := heap.NewClassLoader(cp, cmd.verboseClassFlag)
	return &JVM{
		cmd:         cmd,
		classLoader: classLoader,
		mainThread:  rtdata.NewThread(),
	}
}

func (self *JVM) start() {
	self.initVM()
	self.execMain()
}

//先加载 sun.musc.VM 类,然后执行器类初始化方法
func (self *JVM) initVM() {
	vmClass := self.classLoader.LoadClass("sun/misc/VM")
	base.InitClass(self.mainThread, vmClass)
	interpret(self.mainThread, self.cmd.verboseInstFlag)
}

//先加载主类,再执行其main方法
func (self *JVM) execMain() {
	className := strings.Replace(self.cmd.class, ".", "/", -1)
	mainClass := self.classLoader.LoadClass(className)
	mainMethod := mainClass.GetMainMethod()
	if mainMethod == nil {
		fmt.Printf("Main method not found in class %s\n", self.cmd.class)
		return
	}

	argsArr := self.createArgsArray()
	frame := self.mainThread.NewFrame(mainMethod)
	frame.LocalVars().SetRef(0, argsArr)
	self.mainThread.PushFrame(frame)
	interpret(self.mainThread, self.cmd.verboseInstFlag)
}

func (self *JVM) createArgsArray() *heap.Object {
	stringClass := self.classLoader.LoadClass("java/lang/String")
	argsLen := uint(len(self.cmd.args))
	argsArr := stringClass.ArrayClass().NewArray(argsLen)
	jArgs := argsArr.Refs()
	for i, arg := range self.cmd.args {
		jArgs[i] = heap.JString(self.classLoader, arg)
	}
	return argsArr
}
