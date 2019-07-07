package heap

import "jvmgo/classfile"

//异常处理表 包含多个ExceptionHandler
type ExceptionTable []*ExceptionHandler

//异常处理的结构
type ExceptionHandler struct {
	startPc   int       //try开始位置
	endPc     int       //try结束位置
	handlerPc int       //catch代码块位置
	catchType *ClassRef //异常类的引用
}

//把class文件的异常处理表转换成ExceptionTable
func newExceptionTable(entries []*classfile.ExceptionTableEntry, cp *ConstantPool) ExceptionTable {
	table := make([]*ExceptionHandler, len(entries)) //准备异常处理表的数组
	for i, entry := range entries {
		table[i] = &ExceptionHandler{
			startPc:   int(entry.StartPc()),
			endPc:     int(entry.EndPc()),
			handlerPc: int(entry.HandlerPc()),
			catchType: getCatchType(uint(entry.CatchType()), cp), //从常量池读取异常类型
		}
	}

	return table
}

func getCatchType(index uint, cp *ConstantPool) *ClassRef {
	if index == 0 {
		return nil // catch all, finally块捕获所有异常的处理
	}
	return cp.GetConstant(index).(*ClassRef)
}

//根据异常类和抛错位置查找异常处理
func (self ExceptionTable) findExceptionHandler(exClass *Class, pc int) *ExceptionHandler {
	for _, handler := range self {
		// jvms: The start_pc is inclusive and end_pc is exclusive
		if pc >= handler.startPc && pc < handler.endPc {
			if handler.catchType == nil { //final块
				return handler
			}
			catchClass := handler.catchType.ResolvedClass()                  //handler声明处理的异常类
			if catchClass == exClass || catchClass.IsSuperClassOf(exClass) { //异常类匹配,或是其父类,则处理
				return handler
			}
		}
	}
	return nil
}
