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

func newConstantInfo(tag uint8, cp ConstantPool) ConstantInfo {
	return nil
}
