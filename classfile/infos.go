package classfile

import (
	"math"
)

type constantIntegerInfo struct {
	val int32
}
type constantFloatInfo struct {
	val float32
}
type constantLongInfo struct {
	val int64
}
type constantDoubleInfo struct {
	val float64
}

// modified utf8
type constantUTF8Info struct {
	val string
}
type constantStringInfo struct {
	cp          ConstantPool
	stringIndex uint16
}
type constantClassInfo struct {
	cp        ConstantPool
	nameIndex uint16
}
type constantNameAndType struct {
	nameIndex uint16
	descriptorIndex uint16
}

func (c *constantNameAndType) readInfo(reader *classReader) {
	c.nameIndex = reader.readUint16()
	c.descriptorIndex = reader.readUint16()
}


func (c *constantClassInfo) readInfo(reader *classReader) {
	c.nameIndex = reader.readUint16()
}
func (c *constantClassInfo) Name()  {

}
func (c *constantStringInfo) readInfo(reader *classReader) {
	c.stringIndex = reader.readUint16()
}
func (c *constantStringInfo) String() {
}

func (c *constantUTF8Info) readInfo(reader *classReader) {
	length := uint32(reader.readUint16())
	bytes := reader.readBytes(int(length))
	c.val = decodeMUTF8(bytes)
}

func (c *constantDoubleInfo) readInfo(reader *classReader) {
	bytes := reader.readUint64()
	c.val = math.Float64frombits(bytes)
}

func (c *constantLongInfo) readInfo(reader *classReader) {
	bytes := reader.readUint32()
	c.val = int64(bytes)
}

func (c *constantFloatInfo) readInfo(reader *classReader) {
	bytes := reader.readUint32()
	c.val = math.Float32frombits(bytes)
}

func (c *constantIntegerInfo) readInfo(reader *classReader) {
	bytes := reader.readUint32()
	c.val = int32(bytes)
}

func decodeMUTF8(bytes []byte) string {
	return string(bytes)
}
