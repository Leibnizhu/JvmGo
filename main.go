package main
import "fmt"
import "strings"
import "jvmgo/classpath"
import "jvmgo/classfile"
import "jvmgo/rtdata/heap"

func main(){
	cmd := parseCmd()
	if cmd.versionFlag {
		fmt.Printf("version 0.0.1 Leibniz special")
	} else if cmd.helpFlag || cmd.class == "" {
		printUsage()
	} else {
		startJVM(cmd)
	}
}

func startJVM(cmd *Cmd){
	classpath := classpath.Parse(cmd.XjreOption,cmd.cpOption)
	fmt.Printf("classpath:%s class:%s args:%s\n", classpath, cmd.class,cmd.args)
	className := strings.Replace(cmd.class, ".", "/", -1) //类名改成类路径
	classfile := loadClass(className, classpath)
	fmt.Println(cmd.class)
	printClassInfo(classfile)
	classLoader := heap.NewClassLoader(classpath, cmd.verboseClassFlag)
	mainClass := classLoader.LoadClass(className) //加载主类
	mainMethod := mainClass.GetMainMethod() //获取main()入口函数
	if mainMethod != nil {
		interpret(mainMethod, cmd.verboseInstFlag) //执行main方法
	} else {
		fmt.Printf("Main method not found in class %s\n", cmd.class)
	}
}

func loadClass(className string, classpath *classpath.Classpath) *classfile.ClassFile {
	classData, _, err := classpath.ReadClass(className) //读出类数据
	if err != nil {
		panic(err)
	}
	classfile, err := classfile.Parse(classData) //解析class文件
	if err != nil {
		panic(err)
	}
	return classfile
}

func printClassInfo(classfile *classfile.ClassFile) {
	fmt.Printf("version: %v.%v\n", classfile.MajorVersion(), classfile.MinorVersion())
	fmt.Printf("constants count: %v\n", len(classfile.ConstantPool()))
	fmt.Printf("access flags: 0x%x\n", classfile.AccessFlags())
	fmt.Printf("this class: %v\n", classfile.ClassName())
	fmt.Printf("super class: %v\n", classfile.SuperClassName())
	fmt.Printf("interfaces: %v\n", classfile.InterfaceNames())
	fmt.Printf("fields count: %v\n", len(classfile.Fields()))
	for _, f := range classfile.Fields() {
		fmt.Printf("  %s\n", f.Name())
	}
	fmt.Printf("methods count: %v\n", len(classfile.Methods()))
	for _, m := range classfile.Methods() {
		fmt.Printf("  %s\n", m.Name())
	}
}
