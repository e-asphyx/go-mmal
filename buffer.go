package mmal

/*
#include <interface/mmal/mmal.h>
#include <interface/mmal/mmal_buffer.h>
*/
import "C"

import (
	"math"
	"unsafe"
)

const (
	BufferHeaderFlagEos                  = C.MMAL_BUFFER_HEADER_FLAG_EOS                 // Signals that the start of the current payload starts a frame
	BufferHeaderFlagFrameStart           = C.MMAL_BUFFER_HEADER_FLAG_FRAME_START         // Signals that the end of the current payload ends a frame
	BufferHeaderFlagFrameEnd             = C.MMAL_BUFFER_HEADER_FLAG_FRAME_END           // Signals that the current payload contains only complete frames (1 or more)
	BufferHeaderFlagFrame                = C.MMAL_BUFFER_HEADER_FLAG_FRAME               // Signals that the current payload is a keyframe (i.e. self decodable)
	BufferHeaderFlagKeyframe             = C.MMAL_BUFFER_HEADER_FLAG_KEYFRAME            // Signals a discontinuity in the stream of data (e.g. after a seek). Can be used for instance by a decoder to reset its state
	BufferHeaderFlagDiscontinuity        = C.MMAL_BUFFER_HEADER_FLAG_DISCONTINUITY       // Signals a buffer containing some kind of config data for the component (e.g. codec config data)
	BufferHeaderFlagConfig               = C.MMAL_BUFFER_HEADER_FLAG_CONFIG              // Signals an encrypted payload
	BufferHeaderFlagEncrypted            = C.MMAL_BUFFER_HEADER_FLAG_ENCRYPTED           // Signals a buffer containing side information
	BufferHeaderFlagCodecsideinfo        = C.MMAL_BUFFER_HEADER_FLAG_CODECSIDEINFO       // Signals a buffer which is the snapshot/postview image from a stills capture
	BufferHeaderFlagsSnapshot            = C.MMAL_BUFFER_HEADER_FLAGS_SNAPSHOT           // Signals a buffer which contains data known to be corrupted
	BufferHeaderFlagCorrupted            = C.MMAL_BUFFER_HEADER_FLAG_CORRUPTED           // Signals that a buffer failed to be transmitted
	BufferHeaderFlagTransmissionFailed   = C.MMAL_BUFFER_HEADER_FLAG_TRANSMISSION_FAILED // Signals the output buffer won't be used, just update reference frames
	BufferHeaderFlagDecodeonly           = C.MMAL_BUFFER_HEADER_FLAG_DECODEONLY          // User flags - can be passed in and will get returned
	BufferHeaderFlagUser0                = C.MMAL_BUFFER_HEADER_FLAG_USER0
	BufferHeaderFlagUser1                = C.MMAL_BUFFER_HEADER_FLAG_USER1
	BufferHeaderFlagUser2                = C.MMAL_BUFFER_HEADER_FLAG_USER2
	BufferHeaderFlagUser3                = C.MMAL_BUFFER_HEADER_FLAG_USER3
	BufferHeaderFlagFormatSpecificStart  = C.MMAL_BUFFER_HEADER_FLAG_FORMAT_SPECIFIC_START  //The following flags describe properties of a video buffer header
	BufferHeaderVideoFlagInterlaced      = C.MMAL_BUFFER_HEADER_VIDEO_FLAG_INTERLACED       //Signals an interlaced video frame
	BufferHeaderVideoFlagTopFieldFirst   = C.MMAL_BUFFER_HEADER_VIDEO_FLAG_TOP_FIELD_FIRST  // Signals that the top field of the current interlaced frame should be displayed first
	BufferHeaderVideoFlagDisplayExternal = C.MMAL_BUFFER_HEADER_VIDEO_FLAG_DISPLAY_EXTERNAL // Signals that the buffer should be displayed on external display if attached.
	BufferHeaderVideoFlagProtected       = C.MMAL_BUFFER_HEADER_VIDEO_FLAG_PROTECTED        // Signals that contents of the buffer requires copy protection.
)

type BufferHeaderVideoSpecific struct {
	Planes uint32    // Number of planes composing the video frame
	Offset [4]uint32 // Offsets to the different planes. These must point within the payload buffer
	Pitch  [4]uint32 // Pitch (size in bytes of a line of a plane) of the different planes

	/* Flags describing video specific properties of a buffer header
	(see \ref videobufferheaderflags "Video buffer header flags") */
	Flags uint32
}

// Type specific data that's associated with a payload buffer
type BufferHeaderTypeSpecific C.MMAL_BUFFER_HEADER_TYPE_SPECIFIC_T

func (b *BufferHeaderTypeSpecific) Video() *BufferHeaderVideoSpecific {
	return (*BufferHeaderVideoSpecific)(unsafe.Pointer(b))
}

type BufferHeader struct {
	Next      *BufferHeader // Used to link several buffer headers together
	priv      uintptr       // Data private to the framework
	Cmd       uint32        // Defines what the buffer header contains. This is a FourCC with 0 as a special value meaning stream data
	data      uintptr       // Pointer to the start of the payload buffer (should not be changed by component)
	allocSize uint32        // Allocated size in bytes of payload buffer
	length    uint32        // Number of bytes currently used in the payload buffer (starting from offset)
	offset    uint32        // Offset in bytes to the start of valid data in the payload buffer
	Flags     uint32        // Flags describing properties of a buffer header (see \ref bufferheaderflags "Buffer header flags")

	PTS int64 // Presentation timestamp in microseconds. \ref MMAL_TIME_UNKNOWN is used when the pts is unknown.

	/* Decode timestamp in microseconds (dts = pts, except in the case
	of video streams with B frames). \ref MMAL_TIME_UNKNOWN
	is used when the dts is unknown. */
	DTS int64

	Type *BufferHeaderTypeSpecific // Type specific data that's associated with a payload buffer

	UserData uintptr // Field reserved for use by the client
}

func (b *BufferHeader) cptr() *C.MMAL_BUFFER_HEADER_T {
	return (*C.MMAL_BUFFER_HEADER_T)(unsafe.Pointer(b))
}

func (b *BufferHeader) Data() []byte {
	return (*[math.MaxInt32]byte)(unsafe.Pointer(b.data + uintptr(b.offset)))[:int(b.length):int(b.length)]
}

func (b *BufferHeader) Acquire() {
	C.mmal_buffer_header_acquire(b.cptr())
}

func (b *BufferHeader) Release() {
	C.mmal_buffer_header_release(b.cptr())
}

func (b *BufferHeader) Reset() {
	C.mmal_buffer_header_reset(b.cptr())
}

func (b *BufferHeader) ReleaseContinue() {
	C.mmal_buffer_header_release_continue(b.cptr())
}

func (b *BufferHeader) Replicate(src *BufferHeader) error {
	return newError(C.mmal_buffer_header_replicate(b.cptr(), src.cptr()))
}

func (b *BufferHeader) MemLock() error {
	return newError(C.mmal_buffer_header_mem_lock(b.cptr()))
}

func (b *BufferHeader) MemUnlock() {
	C.mmal_buffer_header_mem_unlock(b.cptr())
}

type _MMAL_BUFFER_HEADER_TYPE_SPECIFIC_T C.MMAL_BUFFER_HEADER_TYPE_SPECIFIC_T
type _MMAL_BUFFER_HEADER_VIDEO_SPECIFIC_T C.MMAL_BUFFER_HEADER_VIDEO_SPECIFIC_T
type _MMAL_BUFFER_HEADER_T C.MMAL_BUFFER_HEADER_T
