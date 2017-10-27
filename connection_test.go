package mmal

import (
	"testing"
	"unsafe"
)

func TestConnectionStructSizes(t *testing.T) {
	testStructSizes(t, []sizeTest{
		{Go: unsafe.Sizeof(Connection{}), C: unsafe.Sizeof(_MMAL_CONNECTION_T{})},
	})
}
