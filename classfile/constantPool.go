package classfile

type ConstantInfo interface {
	readInfo(reader *classReader)
}
type ConstantPool []ConstantInfo

func readConstantInfo(reader *classReader, cp ConstantPool) ConstantInfo {
	tag := reader.readUint8()
	c := newConstantInfo(tag, cp)
	c.readInfo(reader)
	return c
}
func (cp ConstantPool) getConstantInfo(index uint16) ConstantInfo {
	if res := cp[index]; res != nil {
		return res
	}
	panic("Cannot find given constant info")
}
func (cp ConstantPool) getUtf8(index uint16) string {
	res := cp.getConstantInfo(index).(*constantUTF8Info)
	return res.val
}
func (cp ConstantPool) getNameAndType(index uint16) (string, string) {
	ntInfo := cp.getConstantInfo(index).(*constantNameAndType)
	name := cp.getUtf8(ntInfo.nameIndex)
	_type := cp.getUtf8(ntInfo.descriptorIndex)
	return name, _type
}

func (cp ConstantPool) getClassName(index uint16) string {
	classInfo := cp.getConstantInfo(index).(*constantClassInfo)
	return cp.getUtf8(classInfo.nameIndex)
}
func newConstantInfo(tag uint8, cp ConstantPool) ConstantInfo {
	return nil
}
