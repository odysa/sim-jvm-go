package classfile

type sourceFileAttribute struct {
	cp              ConstantPool
	sourceFileIndex uint16
}

func (s sourceFileAttribute) readInfo(reader *classReader) {
	s.sourceFileIndex = reader.readUint16()
}
func (s *sourceFileAttribute) FileName() string {
	return s.cp.getUtf8(s.sourceFileIndex)
}
