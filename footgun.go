package footgun

/*
#include <stdlib.h>

void *footgun_malloc(size_t size) {
	return malloc(size);
}

void footgun_free(void *ptr) {
	free(ptr);
}
*/
import "C"
import (
	"github.com/CannibalVox/footgun/internal/bullet"
	"unsafe"
)

func _testConvertInt(i int) C.int {
	return C.int(i)
}

func Malloc(n C.size_t) unsafe.Pointer {
	alloc := bullet.MallocTrampoline(C.footgun_malloc, uint(n))
	if alloc == nil {
		panic("runtime: C malloc failed")
	}

	return alloc
}

func CString(s string) *C.char {
	ptr := Malloc(C.size_t(len(s)+1))
	ptrptr := (*[1<<30]byte)(ptr)
	copy(ptrptr[:], s)
	ptrptr[len(s)] = 0
	return (*C.char)(ptr)
}

func CBytes(b []byte) unsafe.Pointer {
	ptr := Malloc(C.size_t(len(b)))
	ptrptr := (*[1<<30]byte)(ptr)
	copy(ptrptr[:], b)
	return ptr
}

func GoString(s *C.char) string {
	return C.GoString(s)
}

func GoStringN(s *C.char, l C.int) string {
	return C.GoStringN(s, l)
}

func GoBytes(b unsafe.Pointer, l C.int) []byte {
	return C.GoBytes(b, l)
}

func Free(ptr unsafe.Pointer) {
	bullet.FreeTrampoline(C.footgun_free, ptr)
}
