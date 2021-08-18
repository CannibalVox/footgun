package footgun

import (
	"testing"
	"unsafe"
)
import "github.com/stretchr/testify/require"

func TestFootgunString(t *testing.T) {
	cStr := CString("TestString")
	goStr := GoString(cStr)
	require.Equal(t, "TestString", goStr)
	Free(unsafe.Pointer(cStr))
}

func TestFootgunBytes(t *testing.T) {
	goBytes := []byte("TestString")
	cBytes := CBytes(goBytes)
	goBytes = GoBytes(cBytes, _testConvertInt(len(goBytes)))
	require.Equal(t, "TestString", string(goBytes))
	Free(cBytes)
}

func BenchmarkFootgunMallocStr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cStr := CString("TestString")
		goStr := GoString(cStr)
		if goStr == "" {b.FailNow()}
		Free(unsafe.Pointer(cStr))
	}
}

func BenchmarkFootgunMallocByte(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cBytes := CBytes([]byte{1,2,3,4,5})
		goBytes := GoBytes(cBytes, _testConvertInt(5))
		if goBytes[0] == 0 {b.FailNow()}
		Free(unsafe.Pointer(cBytes))
	}
}

