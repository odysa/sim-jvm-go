package classfile

type exceptionAttribute struct {
	exceptionIndexTable []uint16
}

func (e *exceptionAttribute) readInfo(reader *classReader) {
	e.exceptionIndexTable = reader.readUint16s()
}
func (e *exceptionAttribute) ExceptionIndexTable() []uint16 {
	return e.exceptionIndexTable
}
