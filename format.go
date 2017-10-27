package mmal

/*
#include <interface/mmal/mmal.h>
#include <interface/mmal/mmal_format.h>
*/
import "C"

import (
	"math"
	"unsafe"
)

type ESType uint32

const (
	EsTypeUnknown    ESType = iota // Unknown elementary stream type
	EsTypeControl                  // Elementary stream of control commands
	EsTypeAudio                    // Audio elementary stream
	EsTypeVideo                    // Video elementary stream
	EsTypeSubpicture               // Sub-picture elementary stream (e.g. subtitles, overlays)
)

type EsSpecificFormat C.MMAL_ES_SPECIFIC_FORMAT_T

type ESFormat struct {
	Type ESType // Type of the elementary stream

	/* FourCC specifying the encoding of the elementary stream.
	See the \ref MmalEncodings "pre-defined encodings" for some
	examples.*/
	Encoding uint32

	/* FourCC specifying the specific encoding variant of
	the elementary stream. See the \ref MmalEncodingVariants
	"pre-defined encoding variants" for some examples. */
	EncodingVariant uint32

	Es *EsSpecificFormat // Type specific information for the elementary stream

	Bitrate uint32 // Bitrate in bits per second
	Flags   uint32 // Flags describing properties of the elementary stream. See \ref elementarystreamflags "Elementary stream flags".

	extraDataSize uint32         // Size of the codec specific data
	extraData     unsafe.Pointer // Codec specific data
}

func (e *ESFormat) cptr() *C.MMAL_ES_FORMAT_T {
	return (*C.MMAL_ES_FORMAT_T)(unsafe.Pointer(e))
}

func (e *ESFormat) ExtraData() []byte {
	return (*[math.MaxInt32]byte)(e.extraData)[:int(e.extraDataSize):int(e.extraDataSize)]
}

type VideoFormat struct {
	Width     uint32   // Width of frame in pixels
	Height    uint32   // Height of frame in rows of pixels
	Crop      Rect     // Visible region of the frame
	FrameRate Rational // Frame rate
	PAR       Rational // Pixel aspect ratio

	/* FourCC specifying the color space of the
	video stream. See the \ref MmalColorSpace
	"pre-defined color spaces" for some examples. */
	ColorSpace uint32
}

type AudioFormat struct {
	Channels      uint32 // Number of audio channels
	SampleRate    uint32 // Sample rate
	BitsPerSample uint32 // Bits per sample
	BlockAlign    uint32 // Size of a block of data
}

// Definition of a subpicture format. This describes the properties specific to a subpicture stream
type SubpictureFormat struct {
	XOffset uint32 // Width offset to the start of the subpicture
	YOffset uint32 // Height offset to the start of the subpicture
}

func (e *EsSpecificFormat) Audio() *AudioFormat {
	return (*AudioFormat)(unsafe.Pointer(e))
}

func (e *EsSpecificFormat) Video() *VideoFormat {
	return (*VideoFormat)(unsafe.Pointer(e))
}

func (e *EsSpecificFormat) Subpicture() *SubpictureFormat {
	return (*SubpictureFormat)(unsafe.Pointer(e))
}

func NewFormat() *ESFormat {
	f := C.mmal_format_alloc()
	return (*ESFormat)(unsafe.Pointer(f))
}

func (e *ESFormat) Free() {
	C.mmal_format_free(e.cptr())
}

func (e *ESFormat) ExtradataAlloc(size uint) error {
	return newError(C.mmal_format_extradata_alloc(e.cptr(), C.uint(size)))
}

func (e *ESFormat) CopyFrom(src *ESFormat) {
	C.mmal_format_copy(e.cptr(), src.cptr())
}

func (e *ESFormat) FullCopyFrom(src *ESFormat) error {
	return newError(C.mmal_format_full_copy(e.cptr(), src.cptr()))
}

const (
	ESFormatCompareFlagType             = C.MMAL_ES_FORMAT_COMPARE_FLAG_TYPE
	ESFormatCompareFlagEncoding         = C.MMAL_ES_FORMAT_COMPARE_FLAG_ENCODING
	ESFormatCompareFlagBitrate          = C.MMAL_ES_FORMAT_COMPARE_FLAG_BITRATE
	ESFormatCompareFlagFlags            = C.MMAL_ES_FORMAT_COMPARE_FLAG_FLAGS
	ESFormatCompareFlagExtradata        = C.MMAL_ES_FORMAT_COMPARE_FLAG_EXTRADATA
	ESFormatCompareFlagVideoResolution  = C.MMAL_ES_FORMAT_COMPARE_FLAG_VIDEO_RESOLUTION
	ESFormatCompareFlagVideoCropping    = C.MMAL_ES_FORMAT_COMPARE_FLAG_VIDEO_CROPPING
	ESFormatCompareFlagVideoFrameRate   = C.MMAL_ES_FORMAT_COMPARE_FLAG_VIDEO_FRAME_RATE
	ESFormatCompareFlagVideoAspectRatio = C.MMAL_ES_FORMAT_COMPARE_FLAG_VIDEO_ASPECT_RATIO
	ESFormatCompareFlagVideoColorSpace  = C.MMAL_ES_FORMAT_COMPARE_FLAG_VIDEO_COLOR_SPACE
	ESFormatCompareFlagESOther          = C.MMAL_ES_FORMAT_COMPARE_FLAG_ES_OTHER
)

func (e *ESFormat) Compare(fmt *ESFormat) uint32 {
	return uint32(C.mmal_format_compare(e.cptr(), fmt.cptr()))
}

type _MMAL_ES_FORMAT_T C.MMAL_ES_FORMAT_T
type _MMAL_VIDEO_FORMAT_T C.MMAL_VIDEO_FORMAT_T
type _MMAL_AUDIO_FORMAT_T C.MMAL_AUDIO_FORMAT_T
type _MMAL_SUBPICTURE_FORMAT_T C.MMAL_SUBPICTURE_FORMAT_T
