package classpath
import "os"
import "path/filepath"
import "fmt"
type Classpath struct {
	bootClasspath Entry //启动类
	extClasspath Entry //扩展类
	userClasspath Entry //用户cp
}

func Parse(jreOption, cpOption string) *Classpath {
	cp := &Classpath{}
	cp.parseBootAndExtClasspath(jreOption)
	cp.parseUserClasspath(cpOption)
	return cp
}

func (self *Classpath) parseBootAndExtClasspath(jreOption string) {
	jreDir := getJreDir(jreOption)
	jreLibPath := filepath.Join(jreDir, "lib", "*") // jre/lib/*
	self.bootClasspath = newWildcardEntry(jreLibPath)
	fmt.Printf("Boot classpath: %s\n", jreLibPath)
	jreExtPath := filepath.Join(jreDir, "lib", "ext", "*") // jre/lib/ext/*
	self.extClasspath = newWildcardEntry(jreExtPath)
	fmt.Printf("Ext classpath: %s\n", jreExtPath)
}

func getJreDir(jreOption string) string {
	if jreOption != "" && exists(jreOption) { //优先使用用户的
		return jreOption
	}
	if exists("./jre") { //其次考虑当前目录下的
		return "./jre" 
	}
	if jh := os.Getenv("JAVA_HOME"); jh != "" { //最后使用JAVAHOME下面的jre
		return filepath.Join(jh, "jre")
	}
	panic("Connot find JRE folder!")
}

func exists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func (self *Classpath) parseUserClasspath(cpOption string) {
	if cpOption == "" { //用户没输入cp则使用当前目录
		cpOption = "."
	}
	self.userClasspath = newEntry(cpOption)
}

func (self *Classpath) ReadClass(className string) ([]byte, Entry, error) {
	className = className + ".class"
	if data,entry,err := self.bootClasspath.readClass(className); err == nil { //先从boot类里面找
		return data,entry,nil
	} 
	if data,entry,err := self.extClasspath.readClass(className); err == nil { //先从ext类里面找
		return data,entry,nil
	} 
	return self.userClasspath.readClass(className) //最后使用用户cp
}

func (self *Classpath) String() string { //返回用户类路径
	return self.userClasspath.String()
}