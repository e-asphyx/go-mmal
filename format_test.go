package mmal

import (
	"testing"
	"unsafe"
)

func TestFormatStructSizes(t *testing.T) {
	testStructSizes(t, []sizeTest{
		{Go: unsafe.Sizeof(ESFormat{}), C: unsafe.Sizeof(_MMAL_ES_FORMAT_T{})},
		{Go: unsafe.Sizeof(VideoFormat{}), C: unsafe.Sizeof(_MMAL_VIDEO_FORMAT_T{})},
		{Go: unsafe.Sizeof(AudioFormat{}), C: unsafe.Sizeof(_MMAL_AUDIO_FORMAT_T{})},
		{Go: unsafe.Sizeof(SubpictureFormat{}), C: unsafe.Sizeof(_MMAL_SUBPICTURE_FORMAT_T{})},
	})
}
