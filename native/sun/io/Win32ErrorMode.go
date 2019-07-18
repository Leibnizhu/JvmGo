package io

import "jvmgo/native"
import "jvmgo/rtdata"

func init() {
	native.Register("sun/io/Win32ErrorMode", "setErrorMode", "(J)J", setErrorMode)
}

func setErrorMode(frame *rtdata.Frame) {
	// todo
	frame.OperandStack().PushLong(0)
}
