package classfile

type unparsedAttribute struct {
	name   string
	length uint32
	info   []byte
}

func (u unparsedAttribute) readInfo(reader *classReader) {
	u.info = reader.readBytes(u.length)
}
