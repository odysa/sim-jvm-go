package classfile

type classFile struct {
	magic        uint32
	minorVersion uint16
	majorVersion uint16
	accessFlags  uint16
	thisClass    uint16
	superClass   uint16
	interfaces   []uint16
}

func (c classFile) readAndCheckMagicNumber(reader *classReader) {
	magic := reader.readUint32()
	if magic != 0xCAFEBABE {
		panic("java.lang.ClassFormatError: magic!")
	}
}
func (c *classFile) readAndCheckVersion(reader *classReader)  {
	c.majorVersion = reader.readUint16()
	c.minorVersion = reader.readUint16()
	switch c.majorVersion {
	case 45:
		return
	case 46,47,48,49,50,51,52:
		if c.minorVersion ==0{
			return
		}
	}
	panic("java.lang.unsupportedClassVersionError")
}
