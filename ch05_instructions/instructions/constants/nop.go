package constants

import (
	"jvm_go/ch05_instructions/instructions/base"
	"jvm_go/ch05_instructions/rtda"
)

// Do nothing
type NOP struct{ base.NoOperandsInstruction }

func (self *NOP) Execute(frame *rtda.Frame) {
	// really do nothing
}
