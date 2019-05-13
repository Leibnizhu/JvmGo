package classfile
import "fmt"
import "strconv"
//class文件对应的struct
type ClassFile struct {
	magic uint32 //魔数 CAFE BABE
	minorVersion uint16 //小版本
	majorVersion uint16 //大版本
	constantPool ConstantPool //常量池
	accessFlags uint16 //类访问标志， private/public等
	thisClass uint16 //当前类，指向常量池
	superClass uint16 //父类，常量池指针
	interfaces []uint16 //实现的多个接口 多个常量池指针
	fields []*MemberInfo //类字段
	methods []*MemberInfo //类方法
	attributes []AttributeInfo //属性表
}

//解析class文件
func Parse(classData []byte) (cf *ClassFile, err error) {
	defer func () {
		if r := recover(); r != nil {
			var ok bool
			err,ok = r.(error)
			if !ok {
				err = fmt.Errorf("%v", r)
			}
		}
	}()
	cr := &ClassReader{classData}
	cf = &ClassFile{}
	cf.read(cr)
	return
}

//调用ClassReader解析class文件
func (self *ClassFile) read(reader *ClassReader) {
	self.readAndCheckMagic(reader)
	self.readAndCheckVersion(reader)
	self.constantPool = readConstantPool(reader)
	self.accessFlags = reader.readUint16()
	self.thisClass = reader.readUint16()
	self.superClass = reader.readUint16()
	self.interfaces = reader.readUint16s()
	self.fields = readMembers(reader, self.constantPool)
	self.methods = readMembers(reader, self.constantPool)
	self.attributes = readAttributes(reader, self.constantPool)
}

func (self *ClassFile) readAndCheckMagic(reader *ClassReader) {
	magic := reader.readUint32()
	if magic != 0xCAFEBABE {
		panic("java.lang.ClassFormatError: error magic: " + strconv.FormatInt(int64(magic), 16))
	}
	self.magic = magic
}

func (self *ClassFile) readAndCheckVersion(reader *ClassReader) {
	self.minorVersion = reader.readUint16()
	self.majorVersion = reader.readUint16()
	switch self.majorVersion {
	case 45: //对应JDK1.0.2-1.1, minorVersion有多个
		return
	case 46,47,48,49,50,51,52: //对应JDK1。2-8
		if self.minorVersion == 0 {
			return
		}
	}
	//其他版本不支持
	panic("java.lang.UnsupportedClassVersionError! version=" + strconv.Itoa(int(self.majorVersion)) + "." + strconv.Itoa(int(self.minorVersion)))
}

func (self *ClassFile) MinorVersion() uint16 {
	return self.minorVersion
}

func (self *ClassFile) MajorVersion() uint16 {
	return self.majorVersion
}

func (self *ClassFile) ConstantPool() ConstantPool {
	return self.constantPool
}

func (self *ClassFile) AccessFlags() uint16 {
	return self.accessFlags
}

func (self *ClassFile) Fields() []*MemberInfo {
	return self.fields
}

func (self *ClassFile) Methods() []*MemberInfo {
	return self.methods
}

//从常量池查找类名，thisClass只保存常量池指针
func (self *ClassFile) ClassName() string {
	return self.constantPool.getClassName(self.thisClass)
}

//从常量池查找父类名，superClass只保存常量池指针
func (self *ClassFile) SuperClassName() string {
	if self.superClass > 0 {
		return self.constantPool.getClassName(self.superClass)
	}
	return "" //没有父类
}

//从常量池查找接口名，interfaces只保存常量池指针
func (self *ClassFile) InterfaceNames() []string {
	interfaceNames := make([]string, len(self.interfaces)) //存储接口名的数组
	for i, cpIndex := range self.interfaces { //遍历接口名的常量池指针
		interfaceNames[i] = self.constantPool.getClassName(cpIndex)
	}
	return interfaceNames
}