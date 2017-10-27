package mmal

import (
	"testing"
	"unsafe"
)

func TestPoolStructSizes(t *testing.T) {
	testStructSizes(t, []sizeTest{
		{Go: unsafe.Sizeof(Pool{}), C: unsafe.Sizeof(_MMAL_POOL_T{})},
	})
}
