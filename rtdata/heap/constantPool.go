package heap

import "fmt"
import "jvmgo/classfile"

type Constant interface{}

type ConstantPool struct {
	class  *Class //当前类
	consts []Constant //常量池
}

//class文件的常量池转换成运行时常量池
func newConstantPool(class *Class, cfCp classfile.ConstantPool) *ConstantPool {
	cpCount := len(cfCp)
	consts := make([]Constant, cpCount) //初始化常量池数组
	rtCp := &ConstantPool{class, consts}

	for i := 1; i < cpCount; i++ { //常量池下标从1开始
		cpInfo := cfCp[i]
		// classfile.ConstantPool 转 heap.Constant , 根据不同类型做不同操作
		switch cpInfo.(type) {
			//int 和 float 常量，直接从class文件常量池取出放入consts
		case *classfile.ConstantIntegerInfo:
			intInfo := cpInfo.(*classfile.ConstantIntegerInfo)
			consts[i] = intInfo.Value()
		case *classfile.ConstantFloatInfo:
			floatInfo := cpInfo.(*classfile.ConstantFloatInfo)
			consts[i] = floatInfo.Value()
			//kong 和 double 常量，直接从class文件常量池取出2个位置，放入consts
		case *classfile.ConstantLongInfo:
			longInfo := cpInfo.(*classfile.ConstantLongInfo)
			consts[i] = longInfo.Value() //int64
			i++
		case *classfile.ConstantDoubleInfo:
			doubleInfo := cpInfo.(*classfile.ConstantDoubleInfo)
			consts[i] = doubleInfo.Value() //float64
			i++
			//字符串常量，取出go 的string对象 放入consts
		case *classfile.ConstantStringInfo:
			stringInfo := cpInfo.(*classfile.ConstantStringInfo)
			consts[i] = stringInfo.String()
			//类符号引用
		case *classfile.ConstantClassInfo:
			classInfo := cpInfo.(*classfile.ConstantClassInfo)
			consts[i] = newClassRef(rtCp, classInfo)
			//字段符号医用
		case *classfile.ConstantFieldrefInfo:
			fieldrefInfo := cpInfo.(*classfile.ConstantFieldrefInfo)
			consts[i] = newFieldRef(rtCp, fieldrefInfo)
			//方法符号引用
		case *classfile.ConstantMethodrefInfo:
			methodrefInfo := cpInfo.(*classfile.ConstantMethodrefInfo)
			consts[i] = newMethodRef(rtCp, methodrefInfo)
			//接口方法引用
		case *classfile.ConstantInterfaceMethodrefInfo:
			methodrefInfo := cpInfo.(*classfile.ConstantInterfaceMethodrefInfo)
			consts[i] = newInterfaceMethodRef(rtCp, methodrefInfo)
		default:
			// todo
		}
	}

	return rtCp
}

//根据索引返回常量
func (self *ConstantPool) GetConstant(index uint) Constant {
	if c := self.consts[index]; c != nil {
		return c
	}
	panic(fmt.Sprintf("No constants at index %d", index))
}
