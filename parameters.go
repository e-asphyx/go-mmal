package mmal

/*
#include <stdlib.h>
#include <string.h>
#include <interface/mmal/mmal.h>
#include <interface/mmal/mmal_parameters.h>
*/
import "C"

//go:generate go run internal/paragen/paragen.go

import (
	"unsafe"
)

type Parameter interface {
	header() unsafe.Pointer
}

type ParameterChangeEventRequest _parameterChangeEventRequest
type ParameterZeroCopy _parameterBoolean
type ParameterBufferRequirements _parameterBufferRequirements
type ParameterStatistics _parameterStatistics
type ParameterCoreStatistics _parameterCoreStatistics
type ParameterMemUsage _parameterMemUsage
type ParameterBufferFlagFilter _parameterUint32
type ParameterSeek _parameterSeek
type ParameterPowermonEnable _parameterBoolean
type ParameterLogging _parameterLogging
type ParameterSystemTime _parameterUint64
type ParameterNoImagePadding _parameterBoolean
type ParameterLockstepEnable _parameterBoolean

type ParameterURI struct {
	raw parameterURI
}

type ParameterSupportedEncodings struct {
	raw parameterEncoding
}

type _parameterHeader struct {
	id   uint32 // Parameter ID.
	size uint32 // Size in bytes of the _parameter (including the header)
}

type _parameterBoolean struct {
	hdr    _parameterHeader
	Enable int32
}

type _parameterBufferRequirements struct {
	hdr                   _parameterHeader
	BufferNumMin          uint32
	BufferSizeMin         uint32
	BufferAlignmentMin    uint32
	BufferNumRecommended  uint32
	BufferSizeRecommended uint32
}

type _parameterBytes struct {
	hdr  _parameterHeader
	data [1]byte
	_    [3]byte
}

type _parameterChangeEventRequest struct {
	hdr      _parameterHeader
	ChangeID uint32
	Enable   int32
}

type _parameterConfigfileChunk struct {
	hdr    _parameterHeader
	Size   uint32
	Offset uint32
	data   [1]byte
	_      [3]byte
}

type _parameterConfigfile struct {
	hdr      _parameterHeader
	FileSize uint32
}

type CoreStatsDir uint32

const (
	CoreStatsRx CoreStatsDir = iota
	CoreStatsTx
	CoreStatsMax = 0x7fffffff /* Force 32 bit size for this enum */
)

type _parameterCoreStatistics struct {
	hdr   _parameterHeader
	Dir   CoreStatsDir
	Reset int32
	Stats CoreStatistics
}

type _parameterEncoding struct {
	hdr      _parameterHeader
	encoding [1]uint32
}

type _parameterFrameRate struct {
	hdr       _parameterHeader
	FrameRate Rational
}

type _parameterInt32 struct {
	hdr   _parameterHeader
	Value int32
}

type _parameterInt64 struct {
	hdr   _parameterHeader
	Value int64
}

type _parameterLogging struct {
	hdr   _parameterHeader
	Set   uint32
	Clear uint32
}

type _parameterMemUsage struct {
	hdr              _parameterHeader
	PoolMemAllocSize uint32
}

type ParamMirror uint32

const (
	ParamMirrorNone ParamMirror = iota
	ParamMirrorVertical
	ParamMirrorHorizontal
	ParamMirrorBoth
)

type _parameterMirror struct {
	hdr   _parameterHeader
	Value ParamMirror
}

type _parameterRational struct {
	hdr   _parameterHeader
	Value Rational
}

type _parameterScalefactor struct {
	hdr    _parameterHeader
	ScaleX uint32
	ScaleY uint32
}

type _parameterSeek struct {
	hdr    _parameterHeader
	Offset int64
	Flags  uint32
	_      [4]byte
}

type _parameterStatistics struct {
	hdr                _parameterHeader
	BufferCount        uint32
	FrameCount         uint32
	FramesSkipped      uint32
	FramesDiscarded    uint32
	EOSSeen            uint32
	MaximumFrameBytes  uint32
	TotalBytes         int64
	CorruptMacroblocks uint32
	_                  [4]byte
}

type _parameterString struct {
	hdr _parameterHeader
	str [1]byte
	_   [3]byte
}

type _parameterUint32 struct {
	hdr   _parameterHeader
	Value uint32
}

type _parameterUint64 struct {
	hdr   _parameterHeader
	Value uint64
}

type _parameterURI struct {
	hdr _parameterHeader
	uri [1]byte
	_   [3]byte
}

// Allocated parameters
type parameterURI []byte

func (p parameterURI) string() string {
	ptr := unsafe.Pointer(uintptr(unsafe.Pointer(&p[0])) + unsafe.Offsetof(_parameterURI{}.uri))
	return C.GoString((*C.char)(ptr))
}

func newParameterURIString(s string) parameterURI {
	return newParameterURISlice(append([]byte(s), 0))
}

type parameterEncoding []byte

func (p *ParameterSupportedEncodings) Data() []uint32 {
	return p.raw.data()
}

func NewParameterSupportedEncodings(ln int) *ParameterSupportedEncodings {
	return &ParameterSupportedEncodings{newParameterEncoding(ln)}
}

func NewParameterSupportedEncodingsSlice(data []uint32) *ParameterSupportedEncodings {
	return &ParameterSupportedEncodings{newParameterEncodingSlice(data)}
}

func (p *ParameterURI) String() string {
	return p.raw.string()
}

func NewParameterURI(ln int) *ParameterURI {
	return &ParameterURI{newParameterURI(ln)}
}

func NewParameterURIString(s string) *ParameterURI {
	return &ParameterURI{newParameterURIString(s)}
}

// for sizeof and alignof test
type _MMAL_PARAMETER_HEADER_T C.MMAL_PARAMETER_HEADER_T
type _MMAL_PARAMETER_UINT64_T C.MMAL_PARAMETER_UINT64_T
type _MMAL_PARAMETER_INT64_T C.MMAL_PARAMETER_INT64_T
type _MMAL_PARAMETER_UINT32_T C.MMAL_PARAMETER_UINT32_T
type _MMAL_PARAMETER_INT32_T C.MMAL_PARAMETER_INT32_T
type _MMAL_PARAMETER_RATIONAL_T C.MMAL_PARAMETER_RATIONAL_T
type _MMAL_PARAMETER_BOOLEAN_T C.MMAL_PARAMETER_BOOLEAN_T
type _MMAL_PARAMETER_STRING_T C.MMAL_PARAMETER_STRING_T
type _MMAL_PARAMETER_BYTES_T C.MMAL_PARAMETER_BYTES_T
type _MMAL_PARAMETER_SCALEFACTOR_T C.MMAL_PARAMETER_SCALEFACTOR_T
type _MMAL_PARAM_MIRROR_T C.MMAL_PARAM_MIRROR_T
type _MMAL_PARAMETER_MIRROR_T C.MMAL_PARAMETER_MIRROR_T
type _MMAL_PARAMETER_URI_T C.MMAL_PARAMETER_URI_T
type _MMAL_PARAMETER_ENCODING_T C.MMAL_PARAMETER_ENCODING_T
type _MMAL_PARAMETER_FRAME_RATE_T C.MMAL_PARAMETER_FRAME_RATE_T
type _MMAL_PARAMETER_CONFIGFILE_T C.MMAL_PARAMETER_CONFIGFILE_T
type _MMAL_PARAMETER_CONFIGFILE_CHUNK_T C.MMAL_PARAMETER_CONFIGFILE_CHUNK_T
type _MMAL_PARAMETER_CHANGE_EVENT_REQUEST_T C.MMAL_PARAMETER_CHANGE_EVENT_REQUEST_T
type _MMAL_PARAMETER_BUFFER_REQUIREMENTS_T C.MMAL_PARAMETER_BUFFER_REQUIREMENTS_T
type _MMAL_PARAMETER_SEEK_T C.MMAL_PARAMETER_SEEK_T
type _MMAL_PARAMETER_STATISTICS_T C.MMAL_PARAMETER_STATISTICS_T
type _MMAL_PARAMETER_CORE_STATISTICS_T C.MMAL_PARAMETER_CORE_STATISTICS_T
type _MMAL_PARAMETER_MEM_USAGE_T C.MMAL_PARAMETER_MEM_USAGE_T
type _MMAL_PARAMETER_LOGGING_T C.MMAL_PARAMETER_LOGGING_T
