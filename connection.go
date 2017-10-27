package mmal

/*
#include <interface/mmal/mmal.h>
#include <interface/mmal/util/mmal_connection.h>
*/
import "C"

import (
	"unsafe"
)

const (
	ConnectionFlagTunnelling             = C.MMAL_CONNECTION_FLAG_TUNNELLING
	ConnectionFlagAllocationOnInput      = C.MMAL_CONNECTION_FLAG_ALLOCATION_ON_INPUT
	ConnectionFlagAllocationOnOutput     = C.MMAL_CONNECTION_FLAG_ALLOCATION_ON_OUTPUT
	ConnectionFlagKeepBufferRequirements = C.MMAL_CONNECTION_FLAG_KEEP_BUFFER_REQUIREMENTS
	ConnectionFlagDirect                 = C.MMAL_CONNECTION_FLAG_DIRECT
)

type Connection struct {
	UserData    uintptr
	callback    uintptr
	IsEnabled   uint32
	Flags       uint32
	In          *Port
	Out         *Port
	Pool        *Pool
	Queue       *Queue
	name        *byte
	_           [4]byte
	TimeSetup   int64
	TimeEnable  int64
	TimeDisable int64
}

func (c *Connection) cptr() *C.MMAL_CONNECTION_T {
	return (*C.MMAL_CONNECTION_T)(unsafe.Pointer(c))
}

func (c *Connection) Name() string {
	return C.GoString((*C.char)(c.name))
}

func NewConnection(out, in *Port, flags uint32) (*Connection, error) {
	var conn *Connection
	status := C.mmal_connection_create((**C.MMAL_CONNECTION_T)(unsafe.Pointer(&conn)),
		out.cptr(), in.cptr(), C.uint32_t(flags))
	return conn, newError(status)
}

func (c *Connection) Acquire() {
	C.mmal_connection_acquire(c.cptr())
}

func (c *Connection) Release() error {
	return newError(C.mmal_connection_release(c.cptr()))
}

func (c *Connection) Destroy() error {
	return newError(C.mmal_connection_destroy(c.cptr()))
}

func (c *Connection) Enable() error {
	return newError(C.mmal_connection_enable(c.cptr()))
}

func (c *Connection) Disable() error {
	return newError(C.mmal_connection_disable(c.cptr()))
}

func (c *Connection) EventFormatChanged(buf *BufferHeader) error {
	return newError(C.mmal_connection_event_format_changed(c.cptr(), buf.cptr()))
}

type _MMAL_CONNECTION_T C.MMAL_CONNECTION_T
