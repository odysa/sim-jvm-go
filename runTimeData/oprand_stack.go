package runTimeData

type operandStack struct {
	size  uint
	slots []slot
}

func newOperandStack(size uint) *operandStack {
	if size > 0 {
		return &operandStack{
			size:  size,
			slots: make([]slot, size),
		}
	}
	return nil
}

