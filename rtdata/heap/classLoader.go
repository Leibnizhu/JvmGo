package heap

import "fmt"
import "jvmgo/classfile"
import "jvmgo/classpath"

/*
基础类型: boolean, byte, int ...
基础数组类型: [Z, [B, [I ...
非数组类: java/lang/Object ...
数组类： array classes: [Ljava/lang/Object; ...
*/
type ClassLoader struct {
	cp       *classpath.Classpath
	verboseFlag bool
	classMap map[string]*Class // loaded classes
}

func NewClassLoader(cp *classpath.Classpath, verboseFlag bool) *ClassLoader {
	return &ClassLoader{
		cp:       cp,
		verboseFlag: verboseFlag,
		classMap: make(map[string]*Class),
	}
}

func (self *ClassLoader) LoadClass(name string) *Class {
	if class, ok := self.classMap[name]; ok { //尝试从已加载的里面查
		return class
	}
	if name[0] == '[' { //数组类型
		return self.loadArrayClass(name)
	}
	return self.loadNonArrayClass(name) //map中没有的话，加载（暂不考虑数组类）
}

//加载数组类，主要是一些数组类相关的特殊固有属性
func (self *ClassLoader) loadArrayClass(name string) *Class {
	class := &Class{
		accessFlags: ACC_PUBLIC, //TODO 
		name: name,
		loader: self,
		initStarted: true, //数组类不需要初始化
		superClass: self.LoadClass("java/lang/Object"),
		interfaces: []*Class{
			self.LoadClass("java/lang/Cloneable"),
			self.LoadClass("java/io/Serializable"),
		},
	}
	self.classMap[name] = class
	return class
}

func (self *ClassLoader) loadNonArrayClass(name string) *Class {
	data, entry := self.readClass(name) //调ClassPath取加载类文件
	class := self.defineClass(data) //解析class文件，生成类数据
	link(class) //类的链接
	if self.verboseFlag {
		fmt.Printf("[Loaded %s from %s]\n", name, entry)
	}
	return class
}

func (self *ClassLoader) readClass(name string) ([]byte, classpath.Entry) {
	data, entry, err := self.cp.ReadClass(name)
	if err != nil {
		panic("java.lang.ClassNotFoundException: " + name)
	}
	return data, entry
}

// jvms 5.3.5
func (self *ClassLoader) defineClass(data []byte) *Class {
	class := parseClass(data) //class文件转类结构体
	class.loader = self
	resolveSuperClass(class) //解析父类
	resolveInterfaces(class) //解析接口
	self.classMap[class.name] = class //新加载的类加入缓存
	return class
}

func parseClass(data []byte) *Class {
	cf, err := classfile.Parse(data)
	if err != nil {
		//panic("java.lang.ClassFormatError")
		panic(err)
	}
	return newClass(cf)
}

// jvms 5.4.3.1
func resolveSuperClass(class *Class) {
	if class.name != "java/lang/Object" {
		class.superClass = class.loader.LoadClass(class.superClassName) //递归调用，处理父类
	}
}
func resolveInterfaces(class *Class) {
	interfaceCount := len(class.interfaceNames)
	if interfaceCount > 0 {
		class.interfaces = make([]*Class, interfaceCount) //初始化接口数组
		for i, interfaceName := range class.interfaceNames {
			class.interfaces[i] = class.loader.LoadClass(interfaceName) //递归调用处理接口
		}
	}
}

//类的链接分为验证和准备两个阶段
func link(class *Class) {
	verify(class)
	prepare(class)
}

func verify(class *Class) {
	// todo JVM规范4.10节有规定
}

// 给类变量分配空间，并赋初始值
func prepare(class *Class) {
	calcInstanceFieldSlotIds(class)
	calcStaticFieldSlotIds(class)
	allocAndInitStaticVars(class)
}

//计算实例字段个数
func calcInstanceFieldSlotIds(class *Class) {
	slotId := uint(0) //没有父类的情况从0开始编号
	if class.superClass != nil {
		slotId = class.superClass.instanceSlotCount //在 defineClass 解析阶段，已经递归解析加载了父类
	}
	for _, field := range class.fields {
		if !field.IsStatic() { //只计算非静态的实例字段
			field.slotId = slotId //顺便给字段的slotId编号赋值
			slotId++
			if field.isLongOrDouble() { //long和double占用2个位置
				slotId++
			}
		}
	}
	class.instanceSlotCount = slotId
}

//计算静态字段个数，原理和 calcInstanceFieldSlotIds 一样，区别是不需要处理父类
func calcStaticFieldSlotIds(class *Class) {
	slotId := uint(0)
	for _, field := range class.fields {
		if field.IsStatic() {
			field.slotId = slotId //顺便给字段的slotId编号赋值
			slotId++
			if field.isLongOrDouble() {
				slotId++
			}
		}
	}
	class.staticSlotCount = slotId
}

//给变量分配空间
func allocAndInitStaticVars(class *Class) {
	class.staticVars = newSlots(class.staticSlotCount)
	for _, field := range class.fields { //遍历字段，找常量进行初始化；对于静态非常量，go语言已经有默认值，不需要额外初始化
		if field.IsStatic() && field.IsFinal() {
			initStaticFinalVar(class, field)
		}
	}
}

//初始化常量
func initStaticFinalVar(class *Class, field *Field) {
	vars := class.staticVars
	cp := class.constantPool
	cpIndex := field.ConstValueIndex()
	slotId := field.SlotId()

	if cpIndex > 0 {
		switch field.Descriptor() {
		case "Z", "B", "C", "S", "I": //int short char boolean 等
			val := cp.GetConstant(cpIndex).(int32)
			vars.SetInt(slotId, val)
		case "J":
			val := cp.GetConstant(cpIndex).(int64)
			vars.SetLong(slotId, val)
		case "F":
			val := cp.GetConstant(cpIndex).(float32)
			vars.SetFloat(slotId, val)
		case "D":
			val := cp.GetConstant(cpIndex).(float64)
			vars.SetDouble(slotId, val)
		case "Ljava/lang/String;":
			panic("todo")
		}
	}
}
