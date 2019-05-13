package classfile

/*
属性表
attribute_info {
    u2 attribute_name_index;
    u4 attribute_length; //表长度
    u1 info[attribute_length]; //14种，长度不定
}
*/
type AttributeInfo interface {
	readInfo(reader *ClassReader)
}

//读取属性表
func readAttributes(reader *ClassReader, cp ConstantPool) []AttributeInfo {
	attributesCount := reader.readUint16()
	attributes := make([]AttributeInfo, attributesCount)
	for i := range attributes {
		attributes[i] = readAttribute(reader, cp)
	}
	return attributes
}

//读取单个属性
func readAttribute(reader *ClassReader, cp ConstantPool) AttributeInfo {
	attrNameIndex := reader.readUint16() //属性名的常量池下标
	attrName := cp.getUtf8(attrNameIndex) //属性名
	attrLen := reader.readUint32() //属性长度
	attrInfo := newAttributeInfo(attrName, attrLen, cp) //根据属性名和长度，获取属性对应的人解析类
	attrInfo.readInfo(reader) //调用解析类的readIfo()方法去读取属性信息
	return attrInfo
}

//目前只实现8种
func newAttributeInfo(attrName string, attrLen uint32, cp ConstantPool) AttributeInfo {
	switch attrName {
	case "Code":
		return &CodeAttribute{cp: cp}
	case "ConstantValue":
		return &ConstantValueAttribute{}
	case "Deprecated":
		return &DeprecatedAttribute{}
	case "Exceptions":
		return &ExceptionsAttribute{}
	case "LineNumberTable":
		return &LineNumberTableAttribute{}
	case "LocalVariableTable":
		return &LocalVariableTableAttribute{}
	case "SourceFile":
		return &SourceFileAttribute{cp: cp}
	case "Synthetic":
		return &SyntheticAttribute{}
	default:
		return &UnparsedAttribute{attrName, attrLen, nil}
	}
}
