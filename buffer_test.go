package mmal

import (
	"testing"
	"unsafe"
)

func TestBufferStructSizes(t *testing.T) {
	testStructSizes(t, []sizeTest{
		{Go: unsafe.Sizeof(BufferHeaderTypeSpecific{}), C: unsafe.Sizeof(_MMAL_BUFFER_HEADER_TYPE_SPECIFIC_T{})},
		{Go: unsafe.Sizeof(BufferHeaderVideoSpecific{}), C: unsafe.Sizeof(_MMAL_BUFFER_HEADER_VIDEO_SPECIFIC_T{})},
		{Go: unsafe.Sizeof(BufferHeader{}), C: unsafe.Sizeof(_MMAL_BUFFER_HEADER_T{})},

		{Go: unsafe.Alignof(BufferHeaderTypeSpecific{}), C: unsafe.Alignof(_MMAL_BUFFER_HEADER_TYPE_SPECIFIC_T{})},
		{Go: unsafe.Alignof(BufferHeaderVideoSpecific{}), C: unsafe.Alignof(_MMAL_BUFFER_HEADER_VIDEO_SPECIFIC_T{})},
		{Go: unsafe.Alignof(BufferHeader{}), C: unsafe.Alignof(_MMAL_BUFFER_HEADER_T{})},
	})
}
