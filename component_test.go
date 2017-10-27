package mmal

import (
	"testing"
	"unsafe"
)

func TestComponentStructSizes(t *testing.T) {
	testStructSizes(t, []sizeTest{
		{Go: unsafe.Sizeof(Component{}), C: unsafe.Sizeof(_MMAL_COMPONENT_T{})},
	})
}
