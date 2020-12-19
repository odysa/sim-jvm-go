package classfile

type constantNameAndType struct {
	nameIndex       uint16
	descriptorIndex uint16
}

func (c *constantNameAndType) readInfo(reader *classReader) {
	c.nameIndex = reader.readUint16()
	c.descriptorIndex = reader.readUint16()
}
