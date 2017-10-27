package mmal

/*
#include <stdio.h>
#include <interface/mmal/mmal.h>
#include "interface/mmal/util/mmal_util.h"

void portCbWrapper(MMAL_PORT_T *port, MMAL_BUFFER_HEADER_T *buffer);
*/
import "C"

import (
	"math"
	"sync"
	"unsafe"
)

type PortType uint32

const (
	PortTypeUnknown PortType = iota // Unknown port type
	PortTypeControl                 // Control port
	PortTypeInput                   // Input port
	PortTypeOutput                  // Output port
	PortTypeClock                   // Clock port
	PortTypeInvalid PortType = 0xffffffff
)

type Port struct {
	priv uintptr // Private member used by the framework
	name *C.char // Port name. Used for debugging purposes (Read Only)

	Type     PortType // Type of the port (Read Only)
	Index    uint16   // Index of the port in its type list (Read Only)
	IndexAll uint16   // Index of the port in the list of all ports (Read Only)

	IsEnabled uint32    // Indicates whether the port is enabled or not (Read Only)
	Format    *ESFormat // Format of the elementary stream

	BufferNumMin  uint32 // Minimum number of buffers the port requires (Read Only). This is set by the component.
	BufferSizeMin uint32 // Minimum size of buffers the port requires (Read Only). This is set by the component.

	/* Minimum alignment requirement for the buffers (Read Only).
	A value of zero means no special alignment requirements.
	This is set by the component. */
	BufferAlignmentMin uint32

	/* Number of buffers the port recommends for optimal performance (Read Only).
	A value of zero means no special recommendation.
	This is set by the component. */
	BufferNumRecommended uint32

	/* Size of buffers the port recommends for optimal performance (Read Only).
	A value of zero means no special recommendation.
	This is set by the component. */
	BufferSizeRecommended uint32

	BufferNum  uint32 // Actual number of buffers the port will use. This is set by the client.
	BufferSize uint32 // Actual maximum size of the buffers that will be sent to the port. This is set by the client.

	Component *Component // Component this port belongs to (Read Only)
	UserData  uintptr    // Field reserved for use by the client

	/* Flags describing the capabilities of a port (Read Only).
	Bitwise combination of \ref portcapabilities "Port capabilities"
	values. */
	Capabilities uint32
}

func (p *Port) cptr() *C.MMAL_PORT_T {
	return (*C.MMAL_PORT_T)(unsafe.Pointer(p))
}

func (p *Port) Name() string {
	return C.GoString(p.name)
}

func (p *Port) FormatCommit() error {
	return newError(C.mmal_port_format_commit(p.cptr()))
}

type PortBufferHeaderCallback func(*Port, *BufferHeader)

var portCallbacks sync.Map

//export portCbWrapper
func portCbWrapper(port *C.MMAL_PORT_T, buffer *C.MMAL_BUFFER_HEADER_T) {
	p := (*Port)(unsafe.Pointer(port))
	if v, ok := portCallbacks.Load(p); ok {
		v.(PortBufferHeaderCallback)(p, (*BufferHeader)(unsafe.Pointer(buffer)))
	}
}

func (p *Port) Enable(cb PortBufferHeaderCallback) error {
	if cb != nil {
		portCallbacks.Store(p, cb)
	}
	return newError(C.mmal_port_enable(p.cptr(), C.MMAL_PORT_BH_CB_T(C.portCbWrapper)))
}

func (p *Port) Disable() error {
	err := newError(C.mmal_port_disable(p.cptr()))
	portCallbacks.Delete(p)
	return err
}

func (p *Port) Flush() error {
	return newError(C.mmal_port_flush(p.cptr()))
}

func (p *Port) ParameterSet(param Parameter) error {
	return newError(C.mmal_port_parameter_set(p.cptr(), (*C.MMAL_PARAMETER_HEADER_T)(param.header())))
}

func (p *Port) ParameterGet(param Parameter) error {
	return newError(C.mmal_port_parameter_get(p.cptr(), (*C.MMAL_PARAMETER_HEADER_T)(param.header())))
}

func (p *Port) SendBuffer(buffer *BufferHeader) error {
	return newError(C.mmal_port_send_buffer(p.cptr(), buffer.cptr()))
}

func (p *Port) Connect(to *Port) error {
	return newError(C.mmal_port_connect(p.cptr(), to.cptr()))
}

func (p *Port) Disconnect() error {
	return newError(C.mmal_port_disconnect(p.cptr()))
}

func (p *Port) PayloadAlloc(size uint32) []byte {
	buf := C.mmal_port_payload_alloc(p.cptr(), C.uint32_t(size))
	return (*[math.MaxInt32]byte)(unsafe.Pointer(buf))[:int(size):int(size)]
}

func (p *Port) PayloadFree(size uint32, buf []byte) {
	C.mmal_port_payload_free(p.cptr(), (*C.uint8_t)(&buf[0]))
}

func (p *Port) EventGet(event uint32) (*BufferHeader, error) {
	var bh *BufferHeader
	err := newError(C.mmal_port_event_get(p.cptr(), (**C.MMAL_BUFFER_HEADER_T)(unsafe.Pointer(&bh)), C.uint32_t(event)))
	return bh, err
}

func (p *Port) PoolCreate(headers, payloadSize uint32) *Pool {
	return (*Pool)(unsafe.Pointer(C.mmal_port_pool_create(p.cptr(), C.uint(headers), C.uint32_t(payloadSize))))
}

func (p *Port) PoolDestroy(pool *Pool) {
	C.mmal_port_pool_destroy(p.cptr(), pool.cptr())
}

type _MMAL_PORT_T C.MMAL_PORT_T
