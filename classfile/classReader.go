package classfile

import "encoding/binary"

type classReader struct {
	data []byte
}

func (c *classReader) readUint8() uint8 {
	val := c.data[0]
	c.data = c.data[1:]
	return val
}
func (c *classReader) readUint16() uint16 {
	//jvm big endian
	val := binary.BigEndian.Uint16(c.data)
	c.data = c.data[2:]
	return val
}
func (c *classReader) readUint32() uint32 {
	val := binary.BigEndian.Uint32(c.data)
	c.data = c.data[4:]
	return val
}
func (c *classReader) readUint64() uint64 {
	val := binary.BigEndian.Uint64(c.data)
	c.data = c.data[8:]
	return val
}
/*
	read uint16 table
 */
func (c *classReader) readUint16s() []uint16 {
	size := c.readUint16()
	res := make([]uint16, size)
	for i, _ := range res {
		res[i] = c.readUint16()
	}
	return res
}
func (c *classReader) readBytes(size int) []byte {
	val := c.data[:size]
	c.data = c.data[size:]
	return val
}
