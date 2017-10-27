package mmal

import (
	"testing"
	"unsafe"
)

func TestTypesStructSizes(t *testing.T) {
	testStructSizes(t, []sizeTest{
		{Go: unsafe.Sizeof(Rect{}), C: unsafe.Sizeof(_MMAL_RECT_T{})},
		{Go: unsafe.Sizeof(Rational{}), C: unsafe.Sizeof(_MMAL_RATIONAL_T{})},
		{Go: unsafe.Sizeof(CoreStatistics{}), C: unsafe.Sizeof(_MMAL_CORE_STATISTICS_T{})},
	})
}
