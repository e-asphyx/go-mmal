package mmal

import (
	"testing"
)

type sizeTest struct {
	C  uintptr
	Go uintptr
}

func testStructSizes(t *testing.T, tests []sizeTest) {
	for i, tst := range tests {
		if tst.Go != tst.C {
			t.Errorf("%d: Go struct size %d != C struct size %d", i, tst.Go, tst.C)
		}
	}
}
