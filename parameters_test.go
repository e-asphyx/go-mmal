package mmal

import (
	"testing"
	"unsafe"
)

func TestParameterStructSizes(t *testing.T) {
	testStructSizes(t, []sizeTest{
		{Go: unsafe.Sizeof(_parameterHeader{}), C: unsafe.Sizeof(_MMAL_PARAMETER_HEADER_T{})},
		{Go: unsafe.Sizeof(_parameterUint64{}), C: unsafe.Sizeof(_MMAL_PARAMETER_UINT64_T{})},
		{Go: unsafe.Sizeof(_parameterInt64{}), C: unsafe.Sizeof(_MMAL_PARAMETER_INT64_T{})},
		{Go: unsafe.Sizeof(_parameterUint32{}), C: unsafe.Sizeof(_MMAL_PARAMETER_UINT32_T{})},
		{Go: unsafe.Sizeof(_parameterInt32{}), C: unsafe.Sizeof(_MMAL_PARAMETER_INT32_T{})},
		{Go: unsafe.Sizeof(_parameterRational{}), C: unsafe.Sizeof(_MMAL_PARAMETER_RATIONAL_T{})},
		{Go: unsafe.Sizeof(_parameterBoolean{}), C: unsafe.Sizeof(_MMAL_PARAMETER_BOOLEAN_T{})},
		{Go: unsafe.Sizeof(_parameterString{}), C: unsafe.Sizeof(_MMAL_PARAMETER_STRING_T{})},
		{Go: unsafe.Sizeof(_parameterBytes{}), C: unsafe.Sizeof(_MMAL_PARAMETER_BYTES_T{})},
		{Go: unsafe.Sizeof(_parameterScalefactor{}), C: unsafe.Sizeof(_MMAL_PARAMETER_SCALEFACTOR_T{})},
		{Go: unsafe.Sizeof(_parameterMirror{}), C: unsafe.Sizeof(_MMAL_PARAMETER_MIRROR_T{})},
		{Go: unsafe.Sizeof(_parameterURI{}), C: unsafe.Sizeof(_MMAL_PARAMETER_URI_T{})},
		{Go: unsafe.Sizeof(_parameterEncoding{}), C: unsafe.Sizeof(_MMAL_PARAMETER_ENCODING_T{})},
		{Go: unsafe.Sizeof(_parameterFrameRate{}), C: unsafe.Sizeof(_MMAL_PARAMETER_FRAME_RATE_T{})},
		{Go: unsafe.Sizeof(_parameterConfigfile{}), C: unsafe.Sizeof(_MMAL_PARAMETER_CONFIGFILE_T{})},
		{Go: unsafe.Sizeof(_parameterConfigfileChunk{}), C: unsafe.Sizeof(_MMAL_PARAMETER_CONFIGFILE_CHUNK_T{})},
		{Go: unsafe.Sizeof(_parameterHeader{}), C: unsafe.Sizeof(_MMAL_PARAMETER_HEADER_T{})},
		{Go: unsafe.Sizeof(_parameterChangeEventRequest{}), C: unsafe.Sizeof(_MMAL_PARAMETER_CHANGE_EVENT_REQUEST_T{})},
		{Go: unsafe.Sizeof(_parameterBufferRequirements{}), C: unsafe.Sizeof(_MMAL_PARAMETER_BUFFER_REQUIREMENTS_T{})},
		{Go: unsafe.Sizeof(_parameterSeek{}), C: unsafe.Sizeof(_MMAL_PARAMETER_SEEK_T{})},
		{Go: unsafe.Sizeof(_parameterStatistics{}), C: unsafe.Sizeof(_MMAL_PARAMETER_STATISTICS_T{})},
		{Go: unsafe.Sizeof(_parameterCoreStatistics{}), C: unsafe.Sizeof(_MMAL_PARAMETER_CORE_STATISTICS_T{})},
		{Go: unsafe.Sizeof(_parameterMemUsage{}), C: unsafe.Sizeof(_MMAL_PARAMETER_MEM_USAGE_T{})},
		{Go: unsafe.Sizeof(_parameterLogging{}), C: unsafe.Sizeof(_MMAL_PARAMETER_LOGGING_T{})},

		{Go: unsafe.Alignof(_parameterHeader{}), C: unsafe.Alignof(_MMAL_PARAMETER_HEADER_T{})},
		{Go: unsafe.Alignof(_parameterUint64{}), C: unsafe.Alignof(_MMAL_PARAMETER_UINT64_T{})},
		{Go: unsafe.Alignof(_parameterInt64{}), C: unsafe.Alignof(_MMAL_PARAMETER_INT64_T{})},
		{Go: unsafe.Alignof(_parameterUint32{}), C: unsafe.Alignof(_MMAL_PARAMETER_UINT32_T{})},
		{Go: unsafe.Alignof(_parameterInt32{}), C: unsafe.Alignof(_MMAL_PARAMETER_INT32_T{})},
		{Go: unsafe.Alignof(_parameterRational{}), C: unsafe.Alignof(_MMAL_PARAMETER_RATIONAL_T{})},
		{Go: unsafe.Alignof(_parameterBoolean{}), C: unsafe.Alignof(_MMAL_PARAMETER_BOOLEAN_T{})},
		{Go: unsafe.Alignof(_parameterString{}), C: unsafe.Alignof(_MMAL_PARAMETER_STRING_T{})},
		{Go: unsafe.Alignof(_parameterBytes{}), C: unsafe.Alignof(_MMAL_PARAMETER_BYTES_T{})},
		{Go: unsafe.Alignof(_parameterScalefactor{}), C: unsafe.Alignof(_MMAL_PARAMETER_SCALEFACTOR_T{})},
		{Go: unsafe.Alignof(_parameterMirror{}), C: unsafe.Alignof(_MMAL_PARAMETER_MIRROR_T{})},
		{Go: unsafe.Alignof(_parameterURI{}), C: unsafe.Alignof(_MMAL_PARAMETER_URI_T{})},
		{Go: unsafe.Alignof(_parameterEncoding{}), C: unsafe.Alignof(_MMAL_PARAMETER_ENCODING_T{})},
		{Go: unsafe.Alignof(_parameterFrameRate{}), C: unsafe.Alignof(_MMAL_PARAMETER_FRAME_RATE_T{})},
		{Go: unsafe.Alignof(_parameterConfigfile{}), C: unsafe.Alignof(_MMAL_PARAMETER_CONFIGFILE_T{})},
		{Go: unsafe.Alignof(_parameterConfigfileChunk{}), C: unsafe.Alignof(_MMAL_PARAMETER_CONFIGFILE_CHUNK_T{})},
		{Go: unsafe.Alignof(_parameterHeader{}), C: unsafe.Alignof(_MMAL_PARAMETER_HEADER_T{})},
		{Go: unsafe.Alignof(_parameterChangeEventRequest{}), C: unsafe.Alignof(_MMAL_PARAMETER_CHANGE_EVENT_REQUEST_T{})},
		{Go: unsafe.Alignof(_parameterBufferRequirements{}), C: unsafe.Alignof(_MMAL_PARAMETER_BUFFER_REQUIREMENTS_T{})},
		{Go: unsafe.Alignof(_parameterSeek{}), C: unsafe.Alignof(_MMAL_PARAMETER_SEEK_T{})},
		{Go: unsafe.Alignof(_parameterStatistics{}), C: unsafe.Alignof(_MMAL_PARAMETER_STATISTICS_T{})},
		{Go: unsafe.Alignof(_parameterCoreStatistics{}), C: unsafe.Alignof(_MMAL_PARAMETER_CORE_STATISTICS_T{})},
		{Go: unsafe.Alignof(_parameterMemUsage{}), C: unsafe.Alignof(_MMAL_PARAMETER_MEM_USAGE_T{})},
		{Go: unsafe.Alignof(_parameterLogging{}), C: unsafe.Alignof(_MMAL_PARAMETER_LOGGING_T{})},
	})
}
