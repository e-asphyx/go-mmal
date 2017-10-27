package mmal

/*
#include <interface/mmal/mmal.h>
#include <interface/mmal/mmal_pool.h>
*/
import "C"

import (
	"unsafe"
)

type Pool struct {
	Queue      *Queue
	headersNum uint32
	header     unsafe.Pointer
}

func (p *Pool) cptr() *C.MMAL_POOL_T {
	return (*C.MMAL_POOL_T)(unsafe.Pointer(p))
}

func (p *Pool) Headers() []*BufferHeader {
	return (*[ptrSliceSize]*BufferHeader)(p.header)[:int(p.headersNum):int(p.headersNum)]
}

func NewPool(headers, payloadSize uint32) *Pool {
	return (*Pool)(unsafe.Pointer(C.mmal_pool_create(C.uint(headers), C.uint32_t(payloadSize))))
}

func (p *Pool) Destroy() {
	C.mmal_pool_destroy(p.cptr())
}

func (p *Pool) Resize(headers, payloadSize uint32) error {
	return newError(C.mmal_pool_resize(p.cptr(), C.uint(headers), C.uint32_t(payloadSize)))
}

type _MMAL_POOL_T C.MMAL_POOL_T
