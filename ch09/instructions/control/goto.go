package control

import (
	"jvm_go/ch09/instructions/base"
	"jvm_go/ch09/rtda"
)

// Branch always
type GOTO struct{ base.BranchInstruction }

func (self *GOTO) Execute(frame *rtda.Frame) {
	base.Branch(frame, self.Offset)
}
