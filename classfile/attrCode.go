package classfile

type codeAttribute struct {
	cp             ConstantPool
	maxStack       uint16
	maxLocals      uint16
	code           []byte
	attributes     []AttributeInfo
	exceptionTable []*exceptionTableEntry
}

type exceptionTableEntry struct {
	startPc   uint16
	endPc     uint16
	handlerPc uint16
	catchType uint16
}

func (c *codeAttribute) readInfo(reader *classReader) {
	c.maxStack = reader.readUint16()
	c.maxLocals = reader.readUint16()
	codeLength := reader.readUint32()
	c.code = reader.readBytes(codeLength)
	c.exceptionTable = readExceptionTable(reader)
	c.attributes = readAttributes(reader, c.cp)
}
func readExceptionTable(reader *classReader) []*exceptionTableEntry {
	length := reader.readUint16()
	exceptionTable := make([]*exceptionTableEntry, length)
	for i := range exceptionTable {
		exceptionTable[i] = &exceptionTableEntry{
			startPc:   reader.readUint16(),
			endPc:     reader.readUint16(),
			handlerPc: reader.readUint16(),
			catchType: reader.readUint16(),
		}
	}
	return exceptionTable
}
