package control

import (
	"jvm_go/ch05_instructions/instructions/base"
	"jvm_go/ch05_instructions/rtda"
)

// Branch always
type GOTO struct{ base.BranchInstruction }

func (self *GOTO) Execute(frame *rtda.Frame) {
	base.Branch(frame, self.Offset)
}
