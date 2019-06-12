package heap

type Object struct {
	class  *Class
	data interface{} //类似C语言的void*，可以接受任何类型的值
}

func newObject(class *Class) *Object {
	return &Object{
		class:  class,
		data: newSlots(class.instanceSlotCount), //实例字段空间分配
	}
}

// getter方法
func (self *Object) Class() *Class {
	return self.class
}
func (self *Object) Fields() Slots {
	return self.data.(Slots)
}

//instanceof
func (self *Object) IsInstanceOf(class *Class) bool {
	return class.IsAssignableFrom(self.class) //在 rtdata/heap/classHierarchy.go 中
}
