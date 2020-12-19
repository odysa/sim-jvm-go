package classfile

type constantMemberRefInfo struct {
	cp               ConstantPool
	classIndex       uint16
	nameAndTypeIndex uint16
}
type constantFieldRefInfo struct{ constantMemberRefInfo }
type constantMethodRefInfo struct{ constantMemberRefInfo }
type constantInterfaceMethodRefInfo struct{ constantMemberRefInfo }

func (c *constantMemberRefInfo) readInfo(reader *classReader) {
	c.classIndex = reader.readUint16()
	c.nameAndTypeIndex = reader.readUint16()
}

func (c *constantMemberRefInfo) ClassName() string {
	return c.cp.getClassName(c.classIndex)
}
func (c *constantMemberRefInfo) NameAndDescriptor() (string, string) {
	return c.cp.getNameAndType(c.nameAndTypeIndex)
}
