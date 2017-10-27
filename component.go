package mmal

/*
#include <stdlib.h>
#include <interface/mmal/mmal.h>
#include <interface/mmal/util/mmal_default_components.h>
*/
import "C"

import (
	"math"
	"unsafe"
)

const (
	ComponentDefaultAudioDecoder    = C.MMAL_COMPONENT_DEFAULT_AUDIO_DECODER
	ComponentDefaultAudioRenderer   = C.MMAL_COMPONENT_DEFAULT_AUDIO_RENDERER
	ComponentDefaultCamera          = C.MMAL_COMPONENT_DEFAULT_CAMERA
	ComponentDefaultCameraInfo      = C.MMAL_COMPONENT_DEFAULT_CAMERA_INFO
	ComponentDefaultClock           = C.MMAL_COMPONENT_DEFAULT_CLOCK
	ComponentDefaultContainerReader = C.MMAL_COMPONENT_DEFAULT_CONTAINER_READER
	ComponentDefaultContainerWriter = C.MMAL_COMPONENT_DEFAULT_CONTAINER_WRITER
	ComponentDefaultImageDecoder    = C.MMAL_COMPONENT_DEFAULT_IMAGE_DECODER
	ComponentDefaultImageEncoder    = C.MMAL_COMPONENT_DEFAULT_IMAGE_ENCODER
	ComponentDefaultMiracast        = C.MMAL_COMPONENT_DEFAULT_MIRACAST
	ComponentDefaultScheduler       = C.MMAL_COMPONENT_DEFAULT_SCHEDULER
	ComponentDefaultSplitter        = C.MMAL_COMPONENT_DEFAULT_SPLITTER
	ComponentDefaultVideoConverter  = C.MMAL_COMPONENT_DEFAULT_VIDEO_CONVERTER
	ComponentDefaultVideoDecoder    = C.MMAL_COMPONENT_DEFAULT_VIDEO_DECODER
	ComponentDefaultVideoEncoder    = C.MMAL_COMPONENT_DEFAULT_VIDEO_ENCODER
	ComponentDefaultVideoInjecter   = C.MMAL_COMPONENT_DEFAULT_VIDEO_INJECTER
	ComponentDefaultVideoRenderer   = C.MMAL_COMPONENT_DEFAULT_VIDEO_RENDERER
	ComponentDefaultVideoSplitter   = C.MMAL_COMPONENT_DEFAULT_VIDEO_SPLITTER
)

type Component struct {
	priv      uintptr // Pointer to the private data of the module in use
	UserData  uintptr // Pointer to private data of the client
	name      *C.char // Component name
	IsEnabled uint32  // Specifies whether the component is enabled or not

	/* All components expose a control port.
	The control port is used by clients to set / get parameters that are global to the
	component. It is also used to receive events, which again are global to the component.
	To be able to receive events, the client needs to enable and register a callback on the
	control port. */
	Control *Port

	inputNum uint32         // Number of input ports
	input    unsafe.Pointer // Array of input ports

	outputNum uint32         // Number of output ports
	output    unsafe.Pointer // Array of output ports

	clockNum uint32         // Number of clock ports
	clock    unsafe.Pointer // Array of clock ports

	portNum uint32         // Total number of ports
	port    unsafe.Pointer // Array of all the ports (control/input/output/clock)

	ID uint32 // Uniquely identifies the component's instance within the MMAL * context / process. For debugging.
}

const ptrSliceSize = math.MaxInt32 / 4

func (c *Component) Input() []*Port {
	if c.inputNum == 0 || c.input == unsafe.Pointer(uintptr(0)) {
		return nil
	}
	return (*[ptrSliceSize]*Port)(c.input)[:int(c.inputNum):int(c.inputNum)]
}

func (c *Component) Output() []*Port {
	if c.outputNum == 0 || c.output == unsafe.Pointer(uintptr(0)) {
		return nil
	}
	return (*[ptrSliceSize]*Port)(c.output)[:int(c.outputNum):int(c.outputNum)]
}

func (c *Component) Clock() []*Port {
	if c.clockNum == 0 || c.clock == unsafe.Pointer(uintptr(0)) {
		return nil
	}
	return (*[ptrSliceSize]*Port)(c.clock)[:int(c.clockNum):int(c.clockNum)]
}

func (c *Component) Port() []*Port {
	if c.portNum == 0 || c.port == unsafe.Pointer(uintptr(0)) {
		return nil
	}
	return (*[ptrSliceSize]*Port)(c.port)[:int(c.portNum):int(c.portNum)]
}

func (c *Component) cptr() *C.MMAL_COMPONENT_T {
	return (*C.MMAL_COMPONENT_T)(unsafe.Pointer(c))
}

func (c *Component) Name() string {
	return C.GoString(c.name)
}

func NewComponent(name string) (*Component, error) {
	var component *Component
	n := append([]byte(name), 0)
	status := C.mmal_component_create((*C.char)(unsafe.Pointer(&n[0])), (**C.MMAL_COMPONENT_T)(unsafe.Pointer(&component)))
	return component, newError(status)
}

func (c *Component) Acquire() {
	C.mmal_component_acquire(c.cptr())
}

func (c *Component) Release() error {
	return newError(C.mmal_component_release(c.cptr()))
}

func (c *Component) Destroy() error {
	return newError(C.mmal_component_destroy(c.cptr()))
}

func (c *Component) Enable() error {
	return newError(C.mmal_component_enable(c.cptr()))
}

func (c *Component) Disable() error {
	return newError(C.mmal_component_disable(c.cptr()))
}

type _MMAL_COMPONENT_T C.MMAL_COMPONENT_T
