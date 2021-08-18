package bullet

import "unsafe"

func MallocTrampoline(fn unsafe.Pointer, n uint) unsafe.Pointer
func FreeTrampoline(fn unsafe.Pointer, ptr unsafe.Pointer)
