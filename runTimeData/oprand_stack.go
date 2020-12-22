package runTimeData

import "math"

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

func (stack *operandStack) pushInt(val int32) {
	stack.slots[stack.size].num = val
	stack.size++
}
func (stack *operandStack) popInt() int32 {
	stack.size--
	return stack.slots[stack.size].num
}
func (stack *operandStack) pushFloat(val float32) {
	bits := math.Float32bits(val)
	stack.slots[stack.size].num = int32(bits)
	stack.size++
}
func (stack *operandStack) popFloat() float32 {
	stack.size--
	bits := uint32(stack.slots[stack.size].num)
	return math.Float32frombits(bits)
}
func (stack *operandStack) PushLong(val int64) {
	stack.slots[stack.size].num = int32(val)
	stack.slots[stack.size+1].num = int32(val >> 32)
	stack.size += 2
}
func (stack *operandStack) PopLong() int64 {
	stack.size -= 2
	low := uint32(stack.slots[stack.size].num)
	high := uint32(stack.slots[stack.size+1].num)
	return int64(high)<<32 | int64(low)
}
func (stack *operandStack) PushDouble(val float64) {
	bits := math.Float64bits(val)
	stack.PushLong(int64(bits))
}
func (stack operandStack) PopDouble() float64 {
	bits := uint64(stack.PopLong())
	return math.Float64frombits(bits)
}
func (stack operandStack) PushRef(ref *Object) {
	stack.slots[stack.size].ref = ref
	stack.size++
}
func (stack operandStack) PopRef() *Object {
	stack.size--
	ref := stack.slots[stack.size].ref
	stack.slots[stack.size].ref = nil
	return ref
}
