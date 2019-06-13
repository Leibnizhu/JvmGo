package heap

// instanceof 和 checkcast 指令会用到
//other 是 self 的实例
func (self *Class) IsAssignableFrom(other *Class) bool {
	s, t := other, self
	//TODO 考虑数组类型
	if s == t { //同一个类，必然true
		return true
	}
	
	if !s.IsArray() { //other不是数组
		if !s.IsInterface() { //other不是接口
			if !t.IsInterface() { //self不是接口
				return s.IsSubClassOf(t) //other是self子类
			} else { //self是接口
				return s.IsImplements(t) //other实现了self接口
			}
		} else { //other是接口
			if !t.IsInterface() { //self不是接口
				return t.isJlObject() 
			} else { //self是接口
				return t.isSuperInterfaceOf(s)
			}
		}
	} else { //other是数组
		if !t.IsArray() { //self不是数组但other是数组，要么self是Object，要么是数组实现了的Cloneable和Serializable接口
			if !t.IsInterface() { //self不是接口
				return t.isJlObject()
			} else { //self是接口
				return t.isJlCloneable() || t.isJioSerializable()
			}
		} else { //self和other都是数组,拿出元素类型对比
			sc := s.ComponentClass()
			tc := t.ComponentClass()
			return sc == tc || tc.IsAssignableFrom(sc)
		}
	}

	return false;
}

// 判断self是other的子类
func (self *Class) IsSubClassOf(other *Class) bool {
	for c := self.superClass; c != nil; c = c.superClass { //不断地找self的父类，直到与other相同，或者没有上一层父类了
		if c == other {
			return true
		}
	}
	return false
}

//判断self实现了iface接口
func (self *Class) IsImplements(iface *Class) bool {
	for c := self; c != nil; c = c.superClass { //遍历所有父类
		for _, i := range c.interfaces { //遍历当前父类实现的接口
			if i == iface || i.IsSubInterfaceOf(iface) { //接口相同，或者是iface的子接口
				return true
			}
		}
	}
	return false
}

//判断 self 接口继承了iface 接口
func (self *Class) IsSubInterfaceOf(iface *Class) bool {
	for _, superInterface := range self.interfaces { //遍历当前接口继承的接口
		if superInterface == iface || superInterface.IsSubInterfaceOf(iface) { //接口相同，或者递归调用isSubInterfaceOf判断是否继承接口
			return true
		}
	}
	return false
}

// c extends self
func (self *Class) IsSuperClassOf(other *Class) bool {
	return other.IsSubClassOf(self)
}

// iface extends self
func (self *Class) isSuperInterfaceOf(iface *Class) bool {
	return iface.IsSubInterfaceOf(self)
}