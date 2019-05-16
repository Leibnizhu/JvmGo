package main
import "fmt"
import "strings"
import "jvmgo/classpath"
import "jvmgo/classfile"

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
	cp := classpath.Parse(cmd.XjreOption,cmd.cpOption)
	fmt.Printf("classpath:%s class:%s args:%s\n", cp, cmd.class,cmd.args)
	className := strings.Replace(cmd.class, ".", "/", -1) //类名改成类路径
	cf := loadClass(className, cp)
	fmt.Println(cmd.class)
	printClassInfo(cf)
	mainMethod := getMainMethod(cf) //获取main()入口函数
	if mainMethod != nil {
		interpret(mainMethod)
	} else {
		fmt.Printf("Main method not found in class %s\n", cmd.class)
	}
}

func loadClass(className string, cp *classpath.Classpath) *classfile.ClassFile {
	classData, _, err := cp.ReadClass(className) //读出类数据
	if err != nil {
		panic(err)
	}
	cf, err := classfile.Parse(classData) //解析class文件
	if err != nil {
		panic(err)
	}
	return cf
}

//遍历class文件所有方法，找main方法，注意整个签名
func getMainMethod(cf *classfile.ClassFile) *classfile.MemberInfo {
	for _, m := range cf.Methods() {
		if m.Name() == "main" && m.Descriptor() == "([Ljava/lang/String;)V" {
			return m
		}
	}
	return nil
}

func printClassInfo(cf *classfile.ClassFile) {
	fmt.Printf("version: %v.%v\n", cf.MajorVersion(), cf.MinorVersion())
	fmt.Printf("constants count: %v\n", len(cf.ConstantPool()))
	fmt.Printf("access flags: 0x%x\n", cf.AccessFlags())
	fmt.Printf("this class: %v\n", cf.ClassName())
	fmt.Printf("super class: %v\n", cf.SuperClassName())
	fmt.Printf("interfaces: %v\n", cf.InterfaceNames())
	fmt.Printf("fields count: %v\n", len(cf.Fields()))
	for _, f := range cf.Fields() {
		fmt.Printf("  %s\n", f.Name())
	}
	fmt.Printf("methods count: %v\n", len(cf.Methods()))
	for _, m := range cf.Methods() {
		fmt.Printf("  %s\n", m.Name())
	}
}
