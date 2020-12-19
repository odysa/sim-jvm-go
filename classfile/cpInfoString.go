package classfile

type constantStringInfo struct {
	cp          ConstantPool
	stringIndex uint16
}

func (c *constantStringInfo) readInfo(reader *classReader) {
	c.stringIndex = reader.readUint16()
}
func (c *constantStringInfo) String() {
}
