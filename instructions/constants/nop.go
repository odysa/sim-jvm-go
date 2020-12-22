package constants

import (
	"github.com/odysa/sim-jvm-go/instructions/base"
	"github.com/odysa/sim-jvm-go/runTimeData"
)

type NOP struct {
	base.NoOpInstruction
}

func (receiver NOP) Execute(frame runTimeData.Frame)  {

}
