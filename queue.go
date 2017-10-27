package mmal

/*
#include <interface/mmal/mmal.h>
#include <interface/mmal/mmal_queue.h>
*/
import "C"

import (
	"unsafe"
)

type Queue struct{}

func (q *Queue) cptr() *C.MMAL_QUEUE_T {
	return (*C.MMAL_QUEUE_T)(unsafe.Pointer(q))
}

func NewQueue() *Queue {
	return (*Queue)(unsafe.Pointer(C.mmal_queue_create()))
}

func (q *Queue) Put(buf *BufferHeader) {
	C.mmal_queue_put(q.cptr(), buf.cptr())
}

func (q *Queue) PutBack(buf *BufferHeader) {
	C.mmal_queue_put_back(q.cptr(), buf.cptr())
}

func (q *Queue) Get() *BufferHeader {
	return (*BufferHeader)(unsafe.Pointer(C.mmal_queue_get(q.cptr())))
}

func (q *Queue) Wait() *BufferHeader {
	return (*BufferHeader)(unsafe.Pointer(C.mmal_queue_wait(q.cptr())))
}

func (q *Queue) TimedWait(timeout uint32) *BufferHeader {
	return (*BufferHeader)(unsafe.Pointer(C.mmal_queue_timedwait(q.cptr(), C.VCOS_UNSIGNED(timeout))))
}

func (q *Queue) Length() uint {
	return uint(C.mmal_queue_length(q.cptr()))
}

func (q *Queue) Destroy() {
	C.mmal_queue_destroy(q.cptr())
}
