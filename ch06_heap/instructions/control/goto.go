package control

import (
	"jvm_go/ch06_heap/instructions/base"
	"jvm_go/ch06_heap/rtda"
)

// Branch always
type GOTO struct{ base.BranchInstruction }

func (self *GOTO) Execute(frame *rtda.Frame) {
	base.Branch(frame, self.Offset)
}
