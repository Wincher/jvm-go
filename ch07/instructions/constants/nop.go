package constants

import (
	"jvm_go/ch07/instructions/base"
	"jvm_go/ch07/rtda"
)

// Do nothing
type NOP struct{ base.NoOperandsInstruction }

func (self *NOP) Execute(frame *rtda.Frame) {
	// really do nothing
}
