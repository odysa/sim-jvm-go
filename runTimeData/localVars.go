package runTimeData

import "math"

type localVars []slot

func newLocalVars(size uint) localVars {
	if size > 0 {
		return make([]slot, size)
	}
	return nil
}

func (l localVars) SetInt(index uint, val int32) {
	l[index].num = val
}
func (l localVars) GetInt(index uint) int32 {
	return l[index].num
}

func (l localVars) SetFloat(index uint, val float32) {
	bits := math.Float32bits(val)
	l[index].num = int32(bits)
}
func (l localVars) GetFloat(index uint) float32 {
	bits := uint32(l[index].num)
	return math.Float32frombits(bits)
}

// long consumes two slots
func (l localVars) SetLong(index uint, val int64) {
	l[index].num = int32(val)
	l[index+1].num = int32(val >> 32)
}
func (l localVars) GetLong(index uint) int64 {
	low := uint32(l[index].num)
	high := uint32(l[index+1].num)
	return int64(high)<<32 | int64(low)
}

// double consumes two slots
func (l localVars) SetDouble(index uint, val float64) {
	bits := math.Float64bits(val)
	l.SetLong(index, int64(bits))
}
func (l localVars) GetDouble(index uint) float64 {
	bits := uint64(l.GetLong(index))
	return math.Float64frombits(bits)
}

func (l localVars) SetRef(index uint, ref *Object) {
	l[index].ref = ref
}
func (l localVars) GetRef(index uint) *Object {
	return l[index].ref
}

