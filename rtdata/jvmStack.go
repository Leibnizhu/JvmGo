package rtdata

// jvm 栈， 链表结构
type Stack struct {
	maxSize uint   //最大栈帧容量
	size    uint   //当前大小
	_top    *Frame //栈顶指针
}

func newStack(maxSize uint) *Stack {
	return &Stack{
		maxSize: maxSize,
	}
}

func (self *Stack) push(frame *Frame) {
	if self.size >= self.maxSize {
		panic("java.lang.StackOverflowError") //超过最大栈容量
	}

	if self._top != nil {
		frame.lower = self._top //新栈帧的下一位是旧的栈顶
	}
	self._top = frame
	self.size++
}

func (self *Stack) pop() *Frame {
	if self._top == nil {
		panic("jvm stack is empty!") //栈空
	}

	top := self._top
	self._top = top.lower //新栈顶是出栈栈帧的下一位
	top.lower = nil
	self.size--
	return top
}

func (self *Stack) top() *Frame {
	if self._top == nil {
		panic("jvm stack is empty!") //栈空
	}
	return self._top
}

func (self *Stack) isEmpty() bool {
	return self._top == nil
}

//清空
func (self *Stack) clear() {
	for !self.isEmpty() {
		self.pop()
	}
}
