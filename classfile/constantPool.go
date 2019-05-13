package classfile
import "strconv"
type ConstantPool []ConstantInfo

func readConstantPool (reader *ClassReader) ConstantPool {
	cpCount := int(reader.readUint16()) //常量池大小，-1才是实际大小，下标从1开始
	cp := make([]ConstantInfo, cpCount) //初始化常量池数组
	for i := 1; i < cpCount; i++ { //下标 1 ~ n-1
		cp[i] = readConstantInfo(reader, cp)
		switch cp[i].(type) {
		case *ConstantLongInfo, *ConstantDoubleInfo : //long和doube占两个常量池索引
			i++
		}
	}
	return cp
}

//查找单个常量
func (self ConstantPool) getConstantInfo (index uint16) ConstantInfo {
	if cpInfo := self[index]; cpInfo != nil {
		return cpInfo
	}
	panic("Invalid constant pool index :" + strconv.Itoa(int(index)))
}

//查找字段或方法的名字和描述符
func (self ConstantPool) getNameAndType(index uint16) (string, string) {
	ntInfo := self.getConstantInfo(index).(*ConstantNameAndTypeInfo)
	name := self.getUtf8(ntInfo.nameIndex)
	_type := self.getUtf8(ntInfo.descriptorIndex)
	return name,_type
}

//查找类名
func (self ConstantPool) getClassName(index uint16) string {
	classInfo := self.getConstantInfo(index).(*ConstantClassInfo)
	return self.getUtf8(classInfo.nameIndex)
}

//查找utf-8字符串
func (self ConstantPool) getUtf8(index uint16) string {
	utf8Info := self.getConstantInfo(index).(*ConstantUtf8Info)
	return utf8Info.str
}