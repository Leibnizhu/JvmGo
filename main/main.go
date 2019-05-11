package main
import "fmt"
import "strings"
import "jvmgo/classpath"
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
	classData,_,err := cp.ReadClass(className)
	if err != nil {
		fmt.Printf("Could not find or load main class %s(%s). \n", cmd.class, className)
	}
	fmt.Printf("class data:%v\n",classData)
}