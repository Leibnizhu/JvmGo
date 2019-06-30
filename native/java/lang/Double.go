package lang

import "math"
import "jvmgo/native"
import "jvmgo/rtdata"

const jlDouble = "java/lang/Double"

func init() {
	native.Register(jlDouble, "doubleToRawLongBits", "(D)J", doubleToRawLongBits)
	native.Register(jlDouble, "longBitsToDouble", "(J)D", longBitsToDouble)
}

//对应 public static native long doubleToRawLongBits(double value);
func doubleToRawLongBits(frame *rtdata.Frame) {
	value := frame.LocalVars().GetDouble(0)
	bits := math.Float64bits(value) // todo
	frame.OperandStack().PushLong(int64(bits))
}

//对应 public static native double longBitsToDouble(long bits);
func longBitsToDouble(frame *rtdata.Frame) {
	bits := frame.LocalVars().GetLong(0)
	value := math.Float64frombits(uint64(bits)) // todo
	frame.OperandStack().PushDouble(value)
}
