package classfile

type lineNumberTableAttribute struct {
	lineNumberTable []*lineNumberTableEntry
}
type lineNumberTableEntry struct {
	startPc    uint16
	lineNumber uint16
}

func (l *lineNumberTableAttribute) readInfo(reader *classReader) {
	length := reader.readUint16()
	l.lineNumberTable = make([]*lineNumberTableEntry, length)
	for i := range l.lineNumberTable {
		l.lineNumberTable[i] = &lineNumberTableEntry{
			startPc:    reader.readUint16(),
			lineNumber: reader.readUint16(),
		}
	}
}
