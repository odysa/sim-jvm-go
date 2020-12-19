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
