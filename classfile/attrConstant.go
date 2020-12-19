package classfile

type constantValueAttribute struct {
	constantValueIndex uint16
}

func (c *constantValueAttribute) readInfo(reader *classReader) {
	c.constantValueIndex = reader.readUint16()
}
func (c *constantValueAttribute) ConstantValueIndex() uint16 {
	return c.constantValueIndex
}
