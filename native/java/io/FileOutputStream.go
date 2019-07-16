package io

import "os"
import "unsafe"
import "jvmgo/native"
import "jvmgo/rtdata"

const fos = "java/io/FileOutputStream"

func init() {
	native.Register(fos, "writeBytes", "([BIIZ)V", writeBytes)
}

//对应 private native void writeBytes(byte b[], int off, int len, boolean append) throws IOException;
func writeBytes(frame *rtdata.Frame) {
	vars := frame.LocalVars()
	//this := vars.GetRef(0)
	b := vars.GetRef(1)
	off := vars.GetInt(2)
	len := vars.GetInt(3)
	//append := vars.GetBoolean(4)

	jBytes := b.Data().([]int8)
	goBytes := castInt8sToUint8s(jBytes)
	goBytes = goBytes[off : off+len]
	os.Stdout.Write(goBytes) //利用go自己的Stdout输出控制台
}

//Java的byte是有符号,go的byte是无符号,需要转换
func castInt8sToUint8s(jBytes []int8) (goBytes []byte) {
	ptr := unsafe.Pointer(&jBytes)
	goBytes = *((*[]byte)(ptr))
	return
}
