package io

import (
	"jvmgo/native"
	"jvmgo/rtdata"
)

const fd = "java/io/FileDescriptor"

func init() {
	native.Register(fd, "set", "(I)J", set)
}

//对应 private static native long set(int d);
func set(frame *rtdata.Frame) {
	// todo
	frame.OperandStack().PushLong(0)
}
