package constants

import (
	"jvm_go/ch06_heap/instructions/base"
	"jvm_go/ch06_heap/rtda"
)

// Do nothing
type NOP struct{ base.NoOperandsInstruction }

func (self *NOP) Execute(frame *rtda.Frame) {
	// really do nothing
}
