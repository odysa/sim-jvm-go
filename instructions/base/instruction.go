package base

import "github.com/odysa/sim-jvm-go/runTimeData"

type Instruction interface {
	FetchOp(reader *BytecodeReader)
	Execute(frame *runTimeData.Frame)
}
type NoOpInstruction struct {
}
type BranchInstruction struct {
	Offset int
}

func (b *BranchInstruction) FetchOp(reader *BytecodeReader)  {
	b.Offset = int(reader.ReadInt16())
}

type Index8Instruction struct {
	Index uint
}

func (i *Index8Instruction) FetchOp(reader *BytecodeReader)  {
	i.Index = uint(reader.ReadUint8())
}

type Index16Instruction struct {
	Index uint16
}

func (i *Index16Instruction) FetchOp(reader *BytecodeReader)  {
	i.Index = reader.ReadUint16()
}