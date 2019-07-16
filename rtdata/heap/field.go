package heap

import "jvmgo/classfile"

//字段信息类/结构体

type Field struct {
	ClassMember
	constValueIndex uint //常量池索引
	slotId          uint //记录编号，以便计算对象的静态变量和实例变量所需空间
}

//从 class文件信息对象 构建Field数组
func newFields(class *Class, cfFields []*classfile.MemberInfo) []*Field {
	fields := make([]*Field, len(cfFields)) //初始化数组
	for i, cfField := range cfFields {
		fields[i] = &Field{}
		fields[i].class = class
		fields[i].copyMemberInfo(cfField) //父类 ClassMember 的赋值方法
		fields[i].copyAttributes(cfField)
	}
	return fields
}

//复制属性， class文件 field_info
func (self *Field) copyAttributes(cfField *classfile.MemberInfo) {
	if valAttr := cfField.ConstantValueAttribute(); valAttr != nil { //获取常量属性
		self.constValueIndex = uint(valAttr.ConstantValueIndex())
	}
}

func (self *Field) IsVolatile() bool {
	return 0 != self.accessFlags&ACC_VOLATILE
}
func (self *Field) IsTransient() bool {
	return 0 != self.accessFlags&ACC_TRANSIENT
}
func (self *Field) IsEnum() bool {
	return 0 != self.accessFlags&ACC_ENUM
}

func (self *Field) ConstValueIndex() uint {
	return self.constValueIndex
}
func (self *Field) SlotId() uint {
	return self.slotId
}

//通过描述符判断类型，是否long或double（需要分配2个位置）
func (self *Field) isLongOrDouble() bool {
	return self.descriptor == "J" || self.descriptor == "D"
}

// reflection
func (self *Field) Type() *Class {
	className := toClassName(self.descriptor)
	return self.class.loader.LoadClass(className)
}
