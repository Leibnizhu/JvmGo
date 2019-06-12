package heap

var primitiveTypes = map[string]string{
	"void":    "V",
	"boolean": "Z",
	"byte":    "B",
	"short":   "S",
	"int":     "I",
	"long":    "J",
	"char":    "C",
	"float":   "F",
	"double":  "D",
}

//根据类名得到数组类名
// [XXX -> [[XXX
// int -> [I
// XXX -> [LXXX;
func getArrayClassName(className string) string {
	return "[" + toDescriptor(className)
}

//数组类名，获取元素类名
// [[XXX -> [XXX
// [LXXX; -> XXX
// [I -> int
func getComponentClassName(className string) string {
	if className[0] == '[' {
		componentTypeDescriptor := className[1:] //去掉[，得到描述符
		return toClassName(componentTypeDescriptor)
	}
	panic("Not array: " + className)
}

//类名转描述符
// [XXX => [XXX
// int  => I
// XXX  => LXXX;
func toDescriptor(className string) string {
	if className[0] == '[' { //数组类型
		return className
	}
	if d, ok := primitiveTypes[className]; ok { //基础类型，需要转换
		return d
	}
	return "L" + className + ";" //其他引用类型（非数组）
}

//描述符转实际类名
// [XXX  => [XXX
// LXXX; => XXX
// I     => int
func toClassName(descriptor string) string {
	if descriptor[0] == '[' { //数组类型，直接返回
		return descriptor
	}
	if descriptor[0] == 'L' { //非数组的引用类型，去掉头部的L和尾部的;
		return descriptor[1 : len(descriptor)-1]
	}
	for className, d := range primitiveTypes { //基础类型，需要转换
		if d == descriptor {
			return className
		}
	}
	panic("Invalid descriptor: " + descriptor)
}
