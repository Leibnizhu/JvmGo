package heap

type Object struct {
	class *Class
	data  interface{} //类似C语言的void*，可以接受任何类型的值
	extra interface{} //记录Object结构体实例的额外信息
}

func newObject(class *Class) *Object {
	return &Object{
		class: class,
		data:  newSlots(class.instanceSlotCount), //实例字段空间分配
	}
}

// getter方法
func (self *Object) Class() *Class {
	return self.class
}
func (self *Object) Fields() Slots {
	return self.data.(Slots)
}
func (self *Object) Data() interface{} {
	return self.data
}
func (self *Object) Extra() interface{} {
	return self.extra
}
func (self *Object) SetExtra(extra interface{}) {
	self.extra = extra
}

//instanceof
func (self *Object) IsInstanceOf(class *Class) bool {
	return class.IsAssignableFrom(self.class) //在 rtdata/heap/classHierarchy.go 中
}

//获取对象的引用类型实例变量
func (self *Object) GetRefVar(name, descriptor string) *Object {
	field := self.class.getField(name, descriptor, false)
	slots := self.data.(Slots)
	return slots.GetRef(field.slotId)
}

//给对象的引用类型实例变量赋值
func (self *Object) SetRefVar(name, descriptor string, ref *Object) {
	field := self.class.getField(name, descriptor, false)
	slots := self.data.(Slots)
	slots.SetRef(field.slotId, ref)
}

func (self *Object) SetIntVar(name, descriptor string, val int32) {
	field := self.class.getField(name, descriptor, false)
	slots := self.data.(Slots)
	slots.SetInt(field.slotId, val)
}
func (self *Object) GetIntVar(name, descriptor string) int32 {
	field := self.class.getField(name, descriptor, false)
	slots := self.data.(Slots)
	return slots.GetInt(field.slotId)
}
