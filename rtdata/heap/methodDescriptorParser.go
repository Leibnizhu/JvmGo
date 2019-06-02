package heap

import "strings"
//解析方法描述符
type MethodDescriptorParser struct {
	raw    string //原始描述符
	offset int
	parsed *MethodDescriptor //解析结果
}

func parseMethodDescriptor(descriptor string) *MethodDescriptor {
	parser := &MethodDescriptorParser{} //新建解析器
	return parser.parse(descriptor)
}

func (self *MethodDescriptorParser) parse(descriptor string) *MethodDescriptor {
	self.raw = descriptor
	self.parsed = &MethodDescriptor{}
	self.startParams()
	self.parseParamTypes()
	self.endParams()
	self.parseReturnType()
	self.finish()
	return self.parsed
}

func (self *MethodDescriptorParser) startParams() {
	if self.readUint8() != '(' { //不以(开头，错误
		self.causePanic()
	}
}
func (self *MethodDescriptorParser) endParams() {
	if self.readUint8() != ')' { //不以)结束，错误
		self.causePanic()
	}
}
func (self *MethodDescriptorParser) finish() {
	if self.offset != len(self.raw) { //结束时没读取完所有字符，错误
		self.causePanic()
	}
}

func (self *MethodDescriptorParser) causePanic() {
	panic("BAD descriptor: " + self.raw)
}

func (self *MethodDescriptorParser) readUint8() uint8 {
	b := self.raw[self.offset]
	self.offset++
	return b
}
func (self *MethodDescriptorParser) unreadUint8() {
	self.offset--
}

func (self *MethodDescriptorParser) parseParamTypes() {
	for {
		t := self.parseFieldType()
		if t != "" {
			self.parsed.addParameterType(t) //读出参数类型，追加到描述符对象
		} else {
			break
		}
	}
}

func (self *MethodDescriptorParser) parseReturnType() {
	//无返回值
	if self.readUint8() == 'V' {
		self.parsed.returnType = "V"
		return
	}

	self.unreadUint8() //回退刚才读的不是V的字符
	t := self.parseFieldType()
	if t != "" {
		self.parsed.returnType = t
		return
	}

	self.causePanic()
}

func (self *MethodDescriptorParser) parseFieldType() string {
	switch self.readUint8() {
	case 'B':
		return "B"
	case 'C':
		return "C"
	case 'D':
		return "D"
	case 'F':
		return "F"
	case 'I':
		return "I"
	case 'J':
		return "J"
	case 'S':
		return "S"
	case 'Z':
		return "Z"
	case 'L':
		return self.parseObjectType()
	case '[':
		return self.parseArrayType()
	default:
		self.unreadUint8()
		return ""
	}
}

func (self *MethodDescriptorParser) parseObjectType() string {
	unread := self.raw[self.offset:] //截取未读取的部分描述符
	semicolonIndex := strings.IndexRune(unread, ';')
	if semicolonIndex == -1 { //找到不到;结束符，错误
		self.causePanic()
		return ""
	} else {
		objStart := self.offset - 1
		objEnd := self.offset + semicolonIndex + 1
		self.offset = objEnd
		descriptor := self.raw[objStart:objEnd] //截取当前参数类名
		return descriptor
	}
}

func (self *MethodDescriptorParser) parseArrayType() string {
	arrStart := self.offset - 1
	self.parseFieldType() //解析数组元素的类型
	arrEnd := self.offset
	descriptor := self.raw[arrStart:arrEnd] //整个数组类型，包含[
	return descriptor
}
