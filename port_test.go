package mmal

import (
	"testing"
	"unsafe"
)

func TestPortStructSizes(t *testing.T) {
	testStructSizes(t, []sizeTest{
		{Go: unsafe.Sizeof(Port{}), C: unsafe.Sizeof(_MMAL_PORT_T{})},
	})
}
