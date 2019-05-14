package rtdata

//局部变量表的元素结构体
//根据JVM规范，一位可以存int或引用值，连续2位可以存long或double值
//方案1： []int + unsafe.Pointer 、拿内存地址 ————会被GC
//方案2： []interface{} ————代码可读性差
//方案3： 结构体存整数和引用
type Slot struct {
	num int32 
	ref *Object
}
