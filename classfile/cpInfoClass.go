package classfile

type constantClassInfo struct {
	cp        ConstantPool
	nameIndex uint16
}

func (c *constantClassInfo) readInfo(reader *classReader) {
	c.nameIndex = reader.readUint16()
}
func (c *constantClassInfo) Name() string {
	return c.cp.getClassName(c.nameIndex)
}
