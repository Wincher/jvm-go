package rtda

import "jvm_go/ch06_heap/rtda/heap"

type Slot struct {
	num int32
	ref *heap.Object
}
