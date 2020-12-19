package classfile

type AttributeInfo interface {
	readInfo(reader *classReader)
}

func newAttributeInfo(attrName string, attrLen uint32,cp ConstantPool) AttributeInfo {
	return nil
}
func readAttributes(reader *classReader,cp ConstantPool)  []AttributeInfo{
	attributesCount := reader.readUint16()
	attributes := make([]AttributeInfo,attributesCount)
	for i:=range attributes{
		attributes[i] = readAttribute(reader,cp)
	}
	return attributes
}
func readAttribute(reader *classReader,cp ConstantPool) AttributeInfo{
	attrNameIndex := reader.readUint16()
	attrName := cp.getUtf8(attrNameIndex)
	attrLen := reader.readUint32()
	attrInfo := newAttributeInfo(attrName,attrLen,cp)
	attrInfo.readInfo(reader)
	return attrInfo
}