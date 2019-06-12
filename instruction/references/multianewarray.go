package references

import "jvmgo/instruction/base"
import "jvmgo/rtdata"
import "jvmgo/rtdata/heap"

// 创建多维数组
//2+n个操作数
//前两个操作数在指令后，分别是元素类的符号引用索引，及维度数
//后面n操作数在操作数栈中，表示各维度的长度
//如new int[3][4][5] 对应指令
// iconst_3 iconst_4 iconst_5 //各维度长度
// multianewarrray #5 //[[[I, 3 //新建数组
type MULTI_ANEW_ARRAY struct {
	index      uint16
	dimensions uint8
}

func (self *MULTI_ANEW_ARRAY) FetchOperands(reader *base.BytecodeReader) {
	self.index = reader.ReadUint16() //元素类索引
	self.dimensions = reader.ReadUint8() //维度数
}
func (self *MULTI_ANEW_ARRAY) Execute(frame *rtdata.Frame) {
	cp := frame.Method().Class().ConstantPool()
	classRef := cp.GetConstant(uint(self.index)).(*heap.ClassRef) //元素类的引用
	arrClass := classRef.ResolvedClass() //解析元素类，是多维数组的，以多个[开头

	stack := frame.OperandStack()
	counts := popAndCheckCounts(stack, int(self.dimensions)) //各维度长度出栈，并检查长度>0，返回长度数组
	arr := newMultiDimensionalArray(counts, arrClass) //初始化多维数组
	stack.PushRef(arr)
}

func popAndCheckCounts(stack *rtdata.OperandStack, dimensions int) []int32 {
	counts := make([]int32, dimensions) //长度数组初始化
	for i := dimensions - 1; i >= 0; i-- {
		counts[i] = stack.PopInt() //操作数栈里放了各个维度的长度
		if counts[i] < 0 {
			panic("java.lang.NegativeArraySizeException")
		}
	}
	return counts
}

//递归创建各维度子数组
func newMultiDimensionalArray(counts []int32, arrClass *heap.Class) *heap.Object {
	count := uint(counts[0]) //当前维度的长度
	arr := arrClass.NewArray(count) //新建当前维度的数组
	if len(counts) > 1 { //还有下一维度
		refs := arr.Refs() //当前维度数组对象
		for i := range refs {
			refs[i] = newMultiDimensionalArray(counts[1:], arrClass.ComponentClass()) //递归调用
		}
	}
	return arr //否则返回
}
