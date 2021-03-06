package base

type BytecodeReader struct {
	code []byte
	pc   int
}

func (b *BytecodeReader) Reset(code []byte, pc int) {
	b.code = code
	b.pc = pc
}
func (b *BytecodeReader) ReadUint8() uint8 {
	i := b.code[b.pc]
	b.pc++
	return i
}
func (b BytecodeReader) ReadInt8() int8 {
	return int8(b.ReadUint8())
}
func (b BytecodeReader) ReadUint16() uint16 {
	low := uint16(b.ReadUint8())
	high := uint16(b.ReadUint8())
	return low<<8 | high
}
func (b BytecodeReader) ReadInt16() int16 {
	return int16(b.ReadUint16())
}
func (b BytecodeReader) ReadInt32() int32 {
	byte1 := int32(b.ReadUint8())
	byte2 := int32(b.ReadUint8())
	byte3 := int32(b.ReadUint8())
	byte4 := int32(b.ReadUint8())
	return (byte1 << 24) | (byte2 << 16) | (byte3 << 8) | byte4
}
