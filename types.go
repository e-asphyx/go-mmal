package mmal

/*
#include <interface/mmal/mmal_types.h>
*/
import "C"

const (
	SUCCESS    = iota // Success
	ENOMEM            // Out of memory
	ENOSPC            // Out of resources (other than memory)
	EINVAL            // Argument is invalid
	ENOSYS            // Function not implemented
	ENOENT            // No such file or directory
	ENXIO             // No such device or address
	EIO               // I/O error
	ESPIPE            // Illegal seek
	ECORRUPT          // Data is corrupt \attention FIXME: not POSIX
	ENOTREADY         // Component is not ready \attention FIXME: not POSIX
	ECONFIG           // Component is not configured \attention FIXME: not POSIX
	EISCONN           // Port is already connected
	ENOTCONN          // Port is disconnected
	EAGAIN            // Resource temporarily unavailable. Try again later
	EFAULT            // Bad address
	STATUS_MAX = 0x7FFFFFFF
)

type Rect struct {
	X      int32 // x coordinate (from left)
	Y      int32 // y coordinate (from top)
	Width  int32 // width
	Height int32 // height
}

// Describes a rational number
type Rational struct {
	Num int32 // Numerator
	Den int32 // Denominator
}

type CoreStatistics struct {
	BufferCount     uint32
	FirstBufferTime uint32
	LastBufferTime  uint32
	MaxDelay        uint32
}

var errorDesc = map[uint32]string{
	SUCCESS:   "SUCCESS",
	ENOMEM:    "ENOMEM",
	ENOSPC:    "ENOSPC",
	EINVAL:    "EINVAL",
	ENOSYS:    "ENOSYS",
	ENOENT:    "ENOENT",
	ENXIO:     "ENXIO",
	EIO:       "EIO",
	ESPIPE:    "ESPIPE",
	ECORRUPT:  "ECORRUPT",
	ENOTREADY: "ENOTREADY",
	ECONFIG:   "ECONFIG",
	EISCONN:   "EISCONN",
	ENOTCONN:  "ENOTCONN",
	EAGAIN:    "EAGAIN",
	EFAULT:    "EFAULT",
}

type Error uint32

func (e Error) Error() string {
	return "MMAL: " + errorDesc[uint32(e)]
}

func newError(code C.MMAL_STATUS_T) error {
	if code == SUCCESS {
		return nil
	}

	return Error(code)
}

type _MMAL_RECT_T C.MMAL_RECT_T
type _MMAL_RATIONAL_T C.MMAL_RATIONAL_T
type _MMAL_CORE_STATISTICS_T C.MMAL_CORE_STATISTICS_T
