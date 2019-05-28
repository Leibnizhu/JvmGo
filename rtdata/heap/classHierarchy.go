package heap

// instanceof 和 checkcast 指令会用到
//other 是 self 的实例
func (self *Class) isAssignableFrom(other *Class) bool {
	s, t := other, self
	//TODO 考虑数组类型
	if s == t { //同一个类，必然true
		return true
	}

	if !t.IsInterface() { //self是类
		return s.isSubClassOf(t) //判断other是self的子类
	} else { //self是接口
		return s.isImplements(t) //判断other实现了self接口
	}
}

// 判断self是other的子类
func (self *Class) isSubClassOf(other *Class) bool {
	for c := self.superClass; c != nil; c = c.superClass { //不断地找self的父类，直到与other相同，或者没有上一层父类了
		if c == other {
			return true
		}
	}
	return false
}

//判断self实现了iface接口
func (self *Class) isImplements(iface *Class) bool {
	for c := self; c != nil; c = c.superClass { //遍历所有父类
		for _, i := range c.interfaces { //遍历当前父类实现的接口
			if i == iface || i.isSubInterfaceOf(iface) { //接口相同，或者是iface的子接口
				return true
			}
		}
	}
	return false
}

//判断 self 接口继承了iface 接口
func (self *Class) isSubInterfaceOf(iface *Class) bool {
	for _, superInterface := range self.interfaces { //遍历当前接口继承的接口
		if superInterface == iface || superInterface.isSubInterfaceOf(iface) { //接口相同，或者递归调用isSubInterfaceOf判断是否继承接口
			return true
		}
	}
	return false
}
