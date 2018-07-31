package base

import "jvm_go/ch08/rtda"

//基本逻辑 设置程序计数器位置
func Branch(frame *rtda.Frame, offset int) {
	pc := frame.Thread().PC()
	nextPC := pc + offset
	frame.SetNextPC(nextPC)
}
