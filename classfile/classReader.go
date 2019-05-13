package classfile
import "encoding/binary"
type ClassReader struct {
	data []byte
}

func (self *ClassReader) readUint8() uint8 {
	val := self.data[0]
	self.data = self.data[1:] //切走第一个byte
	return val
}

func (self *ClassReader) readUint16() uint16 {
	val := binary.BigEndian.Uint16(self.data)
	self.data = self.data[2:] //切走前两个byte
	return val
}

func (self *ClassReader) readUint32() uint32 {
	val := binary.BigEndian.Uint32(self.data)
	self.data = self.data[4:] //切走前4个byte
	return val
}

func (self *ClassReader) readUint64() uint64 {
	val := binary.BigEndian.Uint64(self.data)
	self.data = self.data[8:] //切走前8个byte
	return val
}

//读unit16表
func (self *ClassReader) readUint16s() []uint16 {
	n := self.readUint16() //前2个byte表示长度
	s := make([]uint16, n) //结果数组
	for i := range s {
		s[i] = self.readUint16()
	}
	return s
}

//读取指定数量的byte
func (self *ClassReader) readBytes(n uint32) []byte {
	bytes := self.data[:n] //截取前n个byte
	self.data = self.data[n:] 
	return bytes
}