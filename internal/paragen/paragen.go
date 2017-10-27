package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

type Parameter struct {
	GoType string
	ID     string
}

type Array struct {
	GoType      string
	ElemType    string
	ElemLiteral string
	OffsetExpr  string
	HeaderType  string
}

type TemplateData struct {
	Structs   []Parameter
	Allocated []Parameter
	Arrays    []Array
}

var parameters = TemplateData{
	Structs: []Parameter{
		{"ParameterChangeEventRequest", "ParameterIDChangeEventRequest"},
		{"ParameterZeroCopy", "ParameterIDZeroCopy"},
		{"ParameterBufferRequirements", "ParameterIDBufferRequirements"},
		{"ParameterStatistics", "ParameterIDStatistics"},
		{"ParameterCoreStatistics", "ParameterIDCoreStatistics"},
		{"ParameterMemUsage", "ParameterIDMemUsage"},
		{"ParameterBufferFlagFilter", "ParameterIDBufferFlagFilter"},
		{"ParameterSeek", "ParameterIDSeek"},
		{"ParameterPowermonEnable", "ParameterIDPowermonEnable"},
		{"ParameterLogging", "ParameterIDLogging"},
		{"ParameterSystemTime", "ParameterIDSystemTime"},
		{"ParameterNoImagePadding", "ParameterIDNoImagePadding"},
		{"ParameterLockstepEnable", "ParameterIDLockstepEnable"},
		{"ParameterThumbnailConfiguration", "ParameterIDThumbnailConfiguration"},
		{"ParameterRotation", "ParameterIDRotation"},
		{"ParameterEXIFDisable", "ParameterIDEXIFDisable"},
		{"ParameterAWBMode", "ParameterIDAWBMode"},
		{"ParameterImageEffect", "ParameterIDImageEffect"},
		{"ParameterColourEffect", "ParameterIDColourEffect"},
		{"ParameterFlickerAvoid", "ParameterIDFlickerAvoid"},
		{"ParameterFlash", "ParameterIDFlash"},
		{"ParameterRedeye", "ParameterIDRedeye"},
		{"ParameterFocus", "ParameterIDFocus"},
		{"ParameterExposureCompInt", "ParameterIDExposureCompInt"},
		{"ParameterExposureCompRat", "ParameterIDExposureCompRat"},
		{"ParameterZoom", "ParameterIDZoom"},
		{"ParameterMirror", "ParameterIDMirror"},
		{"ParameterCameraNum", "ParameterIDCameraNum"},
		{"ParameterCapture", "ParameterIDCapture"},
		{"ParameterExposureMode", "ParameterIDExposureMode"},
		{"ParameterExpMeteringMode", "ParameterIDExpMeteringMode"},
		{"ParameterFocusStatus", "ParameterIDFocusStatus"},
		{"ParameterCameraConfig", "ParameterIDCameraConfig"},
		{"ParameterCaptureStatus", "ParameterIDCaptureStatus"},
		{"ParameterFaceTrack", "ParameterIDFaceTrack"},
		{"ParameterDrawBoxFacesAndFocus", "ParameterIDDrawBoxFacesAndFocus"},
		{"ParameterJPEGQFactor", "ParameterIDJPEGQFactor"},
		{"ParameterFrameRate", "ParameterIDFrameRate"},
		{"ParameterUseSTC", "ParameterIDUseSTC"},
		{"ParameterCameraInfo", "ParameterIDCameraInfo"},
		{"ParameterVideoStabilisation", "ParameterIDVideoStabilisation"},
		{"ParameterEnableRawCapture", "ParameterIDEnableRawCapture"},
		{"ParameterEnableDPFFile", "ParameterIDEnableDPFFile"},
		{"ParameterDPFFailIsFatal", "ParameterIDDPFFailIsFatal"},
		{"ParameterCaptureMode", "ParameterIDCaptureMode"},
		{"ParameterInputCrop", "ParameterIDInputCrop"},
		{"ParameterSensorInformation", "ParameterIDSensorInformation"},
		{"ParameterFlashSelect", "ParameterIDFlashSelect"},
		{"ParameterFieldOfView", "ParameterIDFieldOfView"},
		{"ParameterHighDynamicRange", "ParameterIDHighDynamicRange"},
		{"ParameterDynamicRangeCompression", "ParameterIDDynamicRangeCompression"},
		{"ParameterAlgorithmControl", "ParameterIDAlgorithmControl"},
		{"ParameterSharpness", "ParameterIDSharpness"},
		{"ParameterContrast", "ParameterIDContrast"},
		{"ParameterBrightness", "ParameterIDBrightness"},
		{"ParameterSaturation", "ParameterIDSaturation"},
		{"ParameterISO", "ParameterIDISO"},
		{"ParameterAntishake", "ParameterIDAntishake"},
		{"ParameterImageEffectParameters", "ParameterIDImageEffectParameters"},
		{"ParameterCameraBurstCapture", "ParameterIDCameraBurstCapture"},
		{"ParameterCameraMinISO", "ParameterIDCameraMinISO"},
		{"ParameterCameraUseCase", "ParameterIDCameraUseCase"},
		{"ParameterCaptureStatsPass", "ParameterIDCaptureStatsPass"},
		{"ParameterCameraCustomSensorConfig", "ParameterIDCameraCustomSensorConfig"},
		{"ParameterEnableRegisterFile", "ParameterIDEnableRegisterFile"},
		{"ParameterRegisterFailIsFatal", "ParameterIDRegisterFailIsFatal"},
		{"ParameterConfigfileRegisters", "ParameterIDConfigfileRegisters"},
		{"ParameterConfigfileChunkRegisters", "ParameterIDConfigfileChunkRegisters"},
		{"ParameterJPEGAttachLog", "ParameterIDJPEGAttachLog"},
		{"ParameterZeroShutterLag", "ParameterIDZeroShutterLag"},
		{"ParameterFPSRange", "ParameterIDFPSRange"},
		{"ParameterCaptureExposureComp", "ParameterIDCaptureExposureComp"},
		{"ParameterSwSharpenDisable", "ParameterIDSwSharpenDisable"},
		{"ParameterFlashRequired", "ParameterIDFlashRequired"},
		{"ParameterSwSaturationDisable", "ParameterIDSwSaturationDisable"},
		{"ParameterShutterSpeed", "ParameterIDShutterSpeed"},
		{"ParameterCustomAWBGains", "ParameterIDCustomAWBGains"},
		{"ParameterCameraSettings", "ParameterIDCameraSettings"},
		{"ParameterPrivacyIndicator", "ParameterIDPrivacyIndicator"},
		{"ParameterVideoDenoise", "ParameterIDVideoDenoise"},
		{"ParameterStillsDenoise", "ParameterIDStillsDenoise"},
		{"ParameterAnnotate", "ParameterIDAnnotate"},
		{"ParameterAnnotateV2", "ParameterIDAnnotateV2"},
		{"ParameterAnnotateV3", "ParameterIDAnnotateV3"},
		{"ParameterStereoscopicMode", "ParameterIDStereoscopicMode"},
		{"ParameterCameraInterface", "ParameterIDCameraInterface"},
		{"ParameterCameraClockingMode", "ParameterIDCameraClockingMode"},
		{"ParameterCameraRxConfig", "ParameterIDCameraRxConfig"},
		{"ParameterCameraRxTiming", "ParameterIDCameraRxTiming"},
		{"ParameterDPFConfig", "ParameterIDDPFConfig"},
		{"ParameterJPEGRestartInterval", "ParameterIDJPEGRestartInterval"},
		{"ParameterCameraIspBlockOverride", "ParameterIDCameraIspBlockOverride"},
		{"ParameterLensShadingOverride", "ParameterIDLensShadingOverride"},
		{"ParameterBlackLevel", "ParameterIDBlackLevel"},
		{"ParameterResizeParams", "ParameterIDResizeParams"},
		{"ParameterCrop", "ParameterIDCrop"},
		{"ParameterOutputShift", "ParameterIDOutputShift"},
		{"ParameterCCMShift", "ParameterIDCCMShift"},
		{"ParameterCustomCCM", "ParameterIDCustomCCM"},
	},
	Allocated: []Parameter{
		{"ParameterURI", "ParameterIDURI"},
		{"ParameterSupportedEncodings", "ParameterIDSupportedEncodings"},
		{"ParameterEXIF", "ParameterIDEXIF"},
		{"ParameterFaceTrackResults", "ParameterIDFaceTrackResults"},
		{"ParameterFocusRegions", "ParameterIDFocusRegions"},
	},
	Arrays: []Array{
		{
			GoType:      "parameterURI",
			ElemType:    "byte",
			ElemLiteral: "byte(0)",
			OffsetExpr:  "_parameterURI{}.uri",
		},
		{
			GoType:      "parameterEncoding",
			ElemType:    "uint32",
			ElemLiteral: "uint32(0)",
			OffsetExpr:  "_parameterEncoding{}.encoding",
		},
		{
			GoType:      "parameterEXIF",
			ElemType:    "uint8",
			ElemLiteral: "uint8(0)",
			OffsetExpr:  "_parameterEXIF{}.data",
			HeaderType:  "ParameterEXIFHeader",
		},
		{
			GoType:      "parameterFaceTrackResults",
			ElemType:    "ParameterFaceTrackFace",
			ElemLiteral: "ParameterFaceTrackFace{}",
			OffsetExpr:  "_parameterFaceTrackResults{}.faces",
			HeaderType:  "ParameterFaceTrackResultsHeader",
		},
		{
			GoType:      "parameterFocusRegions",
			ElemType:    "ParameterFocusRegion",
			ElemLiteral: "ParameterFocusRegion{}",
			OffsetExpr:  "_parameterFocusRegions{}.regions",
			HeaderType:  "ParameterFocusRegionsHeader",
		},
	},
}

const outputFile = "parameters_gen.go"

const funcTmpl = `// GENERATED FILE; DO NOT EDIT

package mmal

/*
#include <string.h>
#include <interface/mmal/mmal.h>
#include <interface/mmal/mmal_parameters.h>
*/
import "C"

import (
	"math"
	"unsafe"
)

// Struct types
{{range .Structs -}}
func (p *{{.GoType}}) header() unsafe.Pointer {
	p.hdr.id = uint32({{.ID}})
	p.hdr.size = uint32(unsafe.Sizeof({{.GoType}}{}))
	return unsafe.Pointer(&p.hdr)
}

{{end -}}

// Types requiring memory preallocation
{{range .Allocated -}}
func (p {{.GoType}}) header() unsafe.Pointer {
	hdr := (*_parameterHeader)(unsafe.Pointer(&p.raw[0]))
	hdr.id = uint32({{.ID}})
	hdr.size = uint32(len(p.raw))
	return unsafe.Pointer(hdr)
}

{{end -}}

// Array types
{{range .Arrays -}}

func (p {{.GoType}}) data() []{{.ElemType}} {
	hdr := (*_parameterHeader)(unsafe.Pointer(&p[0]))
	ln := (int(hdr.size) - int(unsafe.Offsetof({{.OffsetExpr}}))) / int(unsafe.Sizeof({{.ElemLiteral}}))
	ptr := unsafe.Pointer(uintptr(unsafe.Pointer(&p[0])) + unsafe.Offsetof({{.OffsetExpr}}))
	return (*[math.MaxInt32 / unsafe.Sizeof({{.ElemLiteral}})]{{.ElemType}})(ptr)[:int(ln):int(ln)]
}

func new{{.GoType | title}}(ln int) {{.GoType}} {
	l := ln*int(unsafe.Sizeof({{.ElemLiteral}})) + int(unsafe.Offsetof({{.OffsetExpr}}))
	return {{.GoType}}(make([]byte, l, l))
}

func new{{.GoType | title}}Slice(data []{{.ElemType}}) {{.GoType}} {
	ln := len(data)
	p := new{{.GoType | title}}(ln)
	ptr := unsafe.Pointer(uintptr(unsafe.Pointer(&p[0])) + unsafe.Offsetof({{.OffsetExpr}}))
	src := unsafe.Pointer(&data[0])
	C.memcpy(ptr, src, C.size_t(ln*int(unsafe.Sizeof({{.ElemLiteral}}))))
	return p
}

{{if .HeaderType -}}

func (p {{.GoType}}) header() *{{.HeaderType}} {
	return (*{{.HeaderType}})(unsafe.Pointer(&p[0]))
}

{{end -}}
{{end -}}
`

func main() {
	f, err := os.Create(outputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	t := template.Must(template.New("funcs").Funcs(template.FuncMap{
		"title": func(s string) string {
			return strings.Title(s)
		},
	}).Parse(funcTmpl))

	if err := t.Execute(f, parameters); err != nil {
		log.Fatal(err)
	}
}
