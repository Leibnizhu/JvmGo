package heap

// 类符号引用
type SymRef struct {
	cp        *ConstantPool
	className string //符号引用的类名
	class     *Class
}

//类引用的懒解析
func (self *SymRef) ResolvedClass() *Class {
	if self.class == nil {
		self.resolveClassRef()
	}
	return self.class
}

// 类符号引用解析
// 类D 通过符号引用，引用类C时
//先使用类D的类加载器加载类C
//然后检查D是否有权限f访问C
func (self *SymRef) resolveClassRef() {
	d := self.cp.class
	c := d.loader.LoadClass(self.className)
	if !c.isAccessibleTo(d) {
		panic("java.lang.IllegalAccessError")
	}

	self.class = c
}
