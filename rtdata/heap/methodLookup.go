package heap

func LookupMethodInClass(class *Class, name, descriptor string) *Method {
	for c := class; c != nil; c = c.superClass { //从当前类开始，逐渐往父类找，直到找到为止
		for _, method := range c.methods { //遍历当前类的所有方法，从名字和描述符判断是否找到
			if method.name == name && method.descriptor == descriptor {
				return method
			}
		}
	}
	return nil //全都没找到
}

func lookupMethodInInterfaces(ifaces []*Class, name, descriptor string) *Method {
	for _, iface := range ifaces { //遍历接口
		for _, method := range iface.methods { //遍历接口的方法，从名字和描述符判断是否找到方法
			if method.name == name && method.descriptor == descriptor {
				return method
			}
		}
		//遍历完当前接口，还是找不到的话，搜索当前接口的父接口，递归调用lookupMethodInInterfaces
		method := lookupMethodInInterfaces(iface.interfaces, name, descriptor)
		if method != nil {
			return method
		}
	}

	return nil
}
