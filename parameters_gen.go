// GENERATED FILE; DO NOT EDIT

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
func (p *ParameterChangeEventRequest) header() unsafe.Pointer {
	p.hdr.id = uint32(ParameterIDChangeEventRequest)
	p.hdr.size = uint32(unsafe.Sizeof(ParameterChangeEventRequest{}))
	return unsafe.Pointer(&p.hdr)
}

func (p *ParameterZeroCopy) header() unsafe.Pointer {
	p.hdr.id = uint32(ParameterIDZeroCopy)
	p.hdr.size = uint32(unsafe.Sizeof(ParameterZeroCopy{}))
	return unsafe.Pointer(&p.hdr)
}

func (p *ParameterBufferRequirements) header() unsafe.Pointer {
	p.hdr.id = uint32(ParameterIDBufferRequirements)
	p.hdr.size = uint32(unsafe.Sizeof(ParameterBufferRequirements{}))
	return unsafe.Pointer(&p.hdr)
}

func (p *ParameterStatistics) header() unsafe.Pointer {
	p.hdr.id = uint32(ParameterIDStatistics)
	p.hdr.size = uint32(unsafe.Sizeof(ParameterStatistics{}))
	return unsafe.Pointer(&p.hdr)
}

func (p *ParameterCoreStatistics) header() unsafe.Pointer {
	p.hdr.id = uint32(ParameterIDCoreStatistics)
	p.hdr.size = uint32(unsafe.Sizeof(ParameterCoreStatistics{}))
	return unsafe.Pointer(&p.hdr)
}

func (p *ParameterMemUsage) header() unsafe.Pointer {
	p.hdr.id = uint32(ParameterIDMemUsage)
	p.hdr.size = uint32(unsafe.Sizeof(ParameterMemUsage{}))
	return unsafe.Pointer(&p.hdr)
}

func (p *ParameterBufferFlagFilter) header() unsafe.Pointer {
	p.hdr.id = uint32(ParameterIDBufferFlagFilter)
	p.hdr.size = uint32(unsafe.Sizeof(ParameterBufferFlagFilter{}))
	return unsafe.Pointer(&p.hdr)
}

func (p *ParameterSeek) header() unsafe.Pointer {
	p.hdr.id = uint32(ParameterIDSeek)
	p.hdr.size = uint32(unsafe.Sizeof(ParameterSeek{}))
	return unsafe.Pointer(&p.hdr)
}

func (p *ParameterPowermonEnable) header() unsafe.Pointer {
	p.hdr.id = uint32(ParameterIDPowermonEnable)
	p.hdr.size = uint32(unsafe.Sizeof(ParameterPowermonEnable{}))
	return unsafe.Pointer(&p.hdr)
}

func (p *ParameterLogging) header() unsafe.Pointer {
	p.hdr.id = uint32(ParameterIDLogging)
	p.hdr.size = uint32(unsafe.Sizeof(ParameterLogging{}))
	return unsafe.Pointer(&p.hdr)
}

func (p *ParameterSystemTime) header() unsafe.Pointer {
	p.hdr.id = uint32(ParameterIDSystemTime)
	p.hdr.size = uint32(unsafe.Sizeof(ParameterSystemTime{}))
	return unsafe.Pointer(&p.hdr)
}

func (p *ParameterNoImagePadding) header() unsafe.Pointer {
	p.hdr.id = uint32(ParameterIDNoImagePadding)
	p.hdr.size = uint32(unsafe.Sizeof(ParameterNoImagePadding{}))
	return unsafe.Pointer(&p.hdr)
}

func (p *ParameterLockstepEnable) header() unsafe.Pointer {
	p.hdr.id = uint32(ParameterIDLockstepEnable)
	p.hdr.size = uint32(unsafe.Sizeof(ParameterLockstepEnable{}))
	return unsafe.Pointer(&p.hdr)
}

func (p *ParameterThumbnailConfiguration) header() unsafe.Pointer {
	p.hdr.id = uint32(ParameterIDThumbnailConfiguration)
	p.hdr.size = uint32(unsafe.Sizeof(ParameterThumbnailConfiguration{}))
	return unsafe.Pointer(&p.hdr)
}

func (p *ParameterRotation) header() unsafe.Pointer {
	p.hdr.id = uint32(ParameterIDRotation)
	p.hdr.size = uint32(unsafe.Sizeof(ParameterRotation{}))
	return unsafe.Pointer(&p.hdr)
}

func (p *ParameterEXIFDisable) header() unsafe.Pointer {
	p.hdr.id = uint32(ParameterIDEXIFDisable)
	p.hdr.size = uint32(unsafe.Sizeof(ParameterEXIFDisable{}))
	return unsafe.Pointer(&p.hdr)
}

func (p *ParameterAWBMode) header() unsafe.Pointer {
	p.hdr.id = uint32(ParameterIDAWBMode)
	p.hdr.size = uint32(unsafe.Sizeof(ParameterAWBMode{}))
	return unsafe.Pointer(&p.hdr)
}

func (p *ParameterImageEffect) header() unsafe.Pointer {
	p.hdr.id = uint32(ParameterIDImageEffect)
	p.hdr.size = uint32(unsafe.Sizeof(ParameterImageEffect{}))
	return unsafe.Pointer(&p.hdr)
}

func (p *ParameterColourEffect) header() unsafe.Pointer {
	p.hdr.id = uint32(ParameterIDColourEffect)
	p.hdr.size = uint32(unsafe.Sizeof(ParameterColourEffect{}))
	return unsafe.Pointer(&p.hdr)
}

func (p *ParameterFlickerAvoid) header() unsafe.Pointer {
	p.hdr.id = uint32(ParameterIDFlickerAvoid)
	p.hdr.size = uint32(unsafe.Sizeof(ParameterFlickerAvoid{}))
	return unsafe.Pointer(&p.hdr)
}

func (p *ParameterFlash) header() unsafe.Pointer {
	p.hdr.id = uint32(ParameterIDFlash)
	p.hdr.size = uint32(unsafe.Sizeof(ParameterFlash{}))
	return unsafe.Pointer(&p.hdr)
}

func (p *ParameterRedeye) header() unsafe.Pointer {
	p.hdr.id = uint32(ParameterIDRedeye)
	p.hdr.size = uint32(unsafe.Sizeof(ParameterRedeye{}))
	return unsafe.Pointer(&p.hdr)
}

func (p *ParameterFocus) header() unsafe.Pointer {
	p.hdr.id = uint32(ParameterIDFocus)
	p.hdr.size = uint32(unsafe.Sizeof(ParameterFocus{}))
	return unsafe.Pointer(&p.hdr)
}

func (p *ParameterExposureCompInt) header() unsafe.Pointer {
	p.hdr.id = uint32(ParameterIDExposureCompInt)
	p.hdr.size = uint32(unsafe.Sizeof(ParameterExposureCompInt{}))
	return unsafe.Pointer(&p.hdr)
}

func (p *ParameterExposureCompRat) header() unsafe.Pointer {
	p.hdr.id = uint32(ParameterIDExposureCompRat)
	p.hdr.size = uint32(unsafe.Sizeof(ParameterExposureCompRat{}))
	return unsafe.Pointer(&p.hdr)
}

func (p *ParameterZoom) header() unsafe.Pointer {
	p.hdr.id = uint32(ParameterIDZoom)
	p.hdr.size = uint32(unsafe.Sizeof(ParameterZoom{}))
	return unsafe.Pointer(&p.hdr)
}

func (p *ParameterMirror) header() unsafe.Pointer {
	p.hdr.id = uint32(ParameterIDMirror)
	p.hdr.size = uint32(unsafe.Sizeof(ParameterMirror{}))
	return unsafe.Pointer(&p.hdr)
}

func (p *ParameterCameraNum) header() unsafe.Pointer {
	p.hdr.id = uint32(ParameterIDCameraNum)
	p.hdr.size = uint32(unsafe.Sizeof(ParameterCameraNum{}))
	return unsafe.Pointer(&p.hdr)
}

func (p *ParameterCapture) header() unsafe.Pointer {
	p.hdr.id = uint32(ParameterIDCapture)
	p.hdr.size = uint32(unsafe.Sizeof(ParameterCapture{}))
	return unsafe.Pointer(&p.hdr)
}

func (p *ParameterExposureMode) header() unsafe.Pointer {
	p.hdr.id = uint32(ParameterIDExposureMode)
	p.hdr.size = uint32(unsafe.Sizeof(ParameterExposureMode{}))
	return unsafe.Pointer(&p.hdr)
}

func (p *ParameterExpMeteringMode) header() unsafe.Pointer {
	p.hdr.id = uint32(ParameterIDExpMeteringMode)
	p.hdr.size = uint32(unsafe.Sizeof(ParameterExpMeteringMode{}))
	return unsafe.Pointer(&p.hdr)
}

func (p *ParameterFocusStatus) header() unsafe.Pointer {
	p.hdr.id = uint32(ParameterIDFocusStatus)
	p.hdr.size = uint32(unsafe.Sizeof(ParameterFocusStatus{}))
	return unsafe.Pointer(&p.hdr)
}

func (p *ParameterCameraConfig) header() unsafe.Pointer {
	p.hdr.id = uint32(ParameterIDCameraConfig)
	p.hdr.size = uint32(unsafe.Sizeof(ParameterCameraConfig{}))
	return unsafe.Pointer(&p.hdr)
}

func (p *ParameterCaptureStatus) header() unsafe.Pointer {
	p.hdr.id = uint32(ParameterIDCaptureStatus)
	p.hdr.size = uint32(unsafe.Sizeof(ParameterCaptureStatus{}))
	return unsafe.Pointer(&p.hdr)
}

func (p *ParameterFaceTrack) header() unsafe.Pointer {
	p.hdr.id = uint32(ParameterIDFaceTrack)
	p.hdr.size = uint32(unsafe.Sizeof(ParameterFaceTrack{}))
	return unsafe.Pointer(&p.hdr)
}

func (p *ParameterDrawBoxFacesAndFocus) header() unsafe.Pointer {
	p.hdr.id = uint32(ParameterIDDrawBoxFacesAndFocus)
	p.hdr.size = uint32(unsafe.Sizeof(ParameterDrawBoxFacesAndFocus{}))
	return unsafe.Pointer(&p.hdr)
}

func (p *ParameterJPEGQFactor) header() unsafe.Pointer {
	p.hdr.id = uint32(ParameterIDJPEGQFactor)
	p.hdr.size = uint32(unsafe.Sizeof(ParameterJPEGQFactor{}))
	return unsafe.Pointer(&p.hdr)
}

func (p *ParameterFrameRate) header() unsafe.Pointer {
	p.hdr.id = uint32(ParameterIDFrameRate)
	p.hdr.size = uint32(unsafe.Sizeof(ParameterFrameRate{}))
	return unsafe.Pointer(&p.hdr)
}

func (p *ParameterUseSTC) header() unsafe.Pointer {
	p.hdr.id = uint32(ParameterIDUseSTC)
	p.hdr.size = uint32(unsafe.Sizeof(ParameterUseSTC{}))
	return unsafe.Pointer(&p.hdr)
}

func (p *ParameterCameraInfo) header() unsafe.Pointer {
	p.hdr.id = uint32(ParameterIDCameraInfo)
	p.hdr.size = uint32(unsafe.Sizeof(ParameterCameraInfo{}))
	return unsafe.Pointer(&p.hdr)
}

func (p *ParameterVideoStabilisation) header() unsafe.Pointer {
	p.hdr.id = uint32(ParameterIDVideoStabilisation)
	p.hdr.size = uint32(unsafe.Sizeof(ParameterVideoStabilisation{}))
	return unsafe.Pointer(&p.hdr)
}

func (p *ParameterEnableRawCapture) header() unsafe.Pointer {
	p.hdr.id = uint32(ParameterIDEnableRawCapture)
	p.hdr.size = uint32(unsafe.Sizeof(ParameterEnableRawCapture{}))
	return unsafe.Pointer(&p.hdr)
}

func (p *ParameterEnableDPFFile) header() unsafe.Pointer {
	p.hdr.id = uint32(ParameterIDEnableDPFFile)
	p.hdr.size = uint32(unsafe.Sizeof(ParameterEnableDPFFile{}))
	return unsafe.Pointer(&p.hdr)
}

func (p *ParameterDPFFailIsFatal) header() unsafe.Pointer {
	p.hdr.id = uint32(ParameterIDDPFFailIsFatal)
	p.hdr.size = uint32(unsafe.Sizeof(ParameterDPFFailIsFatal{}))
	return unsafe.Pointer(&p.hdr)
}

func (p *ParameterCaptureMode) header() unsafe.Pointer {
	p.hdr.id = uint32(ParameterIDCaptureMode)
	p.hdr.size = uint32(unsafe.Sizeof(ParameterCaptureMode{}))
	return unsafe.Pointer(&p.hdr)
}

func (p *ParameterInputCrop) header() unsafe.Pointer {
	p.hdr.id = uint32(ParameterIDInputCrop)
	p.hdr.size = uint32(unsafe.Sizeof(ParameterInputCrop{}))
	return unsafe.Pointer(&p.hdr)
}

func (p *ParameterSensorInformation) header() unsafe.Pointer {
	p.hdr.id = uint32(ParameterIDSensorInformation)
	p.hdr.size = uint32(unsafe.Sizeof(ParameterSensorInformation{}))
	return unsafe.Pointer(&p.hdr)
}

func (p *ParameterFlashSelect) header() unsafe.Pointer {
	p.hdr.id = uint32(ParameterIDFlashSelect)
	p.hdr.size = uint32(unsafe.Sizeof(ParameterFlashSelect{}))
	return unsafe.Pointer(&p.hdr)
}

func (p *ParameterFieldOfView) header() unsafe.Pointer {
	p.hdr.id = uint32(ParameterIDFieldOfView)
	p.hdr.size = uint32(unsafe.Sizeof(ParameterFieldOfView{}))
	return unsafe.Pointer(&p.hdr)
}

func (p *ParameterHighDynamicRange) header() unsafe.Pointer {
	p.hdr.id = uint32(ParameterIDHighDynamicRange)
	p.hdr.size = uint32(unsafe.Sizeof(ParameterHighDynamicRange{}))
	return unsafe.Pointer(&p.hdr)
}

func (p *ParameterDynamicRangeCompression) header() unsafe.Pointer {
	p.hdr.id = uint32(ParameterIDDynamicRangeCompression)
	p.hdr.size = uint32(unsafe.Sizeof(ParameterDynamicRangeCompression{}))
	return unsafe.Pointer(&p.hdr)
}

func (p *ParameterAlgorithmControl) header() unsafe.Pointer {
	p.hdr.id = uint32(ParameterIDAlgorithmControl)
	p.hdr.size = uint32(unsafe.Sizeof(ParameterAlgorithmControl{}))
	return unsafe.Pointer(&p.hdr)
}

func (p *ParameterSharpness) header() unsafe.Pointer {
	p.hdr.id = uint32(ParameterIDSharpness)
	p.hdr.size = uint32(unsafe.Sizeof(ParameterSharpness{}))
	return unsafe.Pointer(&p.hdr)
}

func (p *ParameterContrast) header() unsafe.Pointer {
	p.hdr.id = uint32(ParameterIDContrast)
	p.hdr.size = uint32(unsafe.Sizeof(ParameterContrast{}))
	return unsafe.Pointer(&p.hdr)
}

func (p *ParameterBrightness) header() unsafe.Pointer {
	p.hdr.id = uint32(ParameterIDBrightness)
	p.hdr.size = uint32(unsafe.Sizeof(ParameterBrightness{}))
	return unsafe.Pointer(&p.hdr)
}

func (p *ParameterSaturation) header() unsafe.Pointer {
	p.hdr.id = uint32(ParameterIDSaturation)
	p.hdr.size = uint32(unsafe.Sizeof(ParameterSaturation{}))
	return unsafe.Pointer(&p.hdr)
}

func (p *ParameterISO) header() unsafe.Pointer {
	p.hdr.id = uint32(ParameterIDISO)
	p.hdr.size = uint32(unsafe.Sizeof(ParameterISO{}))
	return unsafe.Pointer(&p.hdr)
}

func (p *ParameterAntishake) header() unsafe.Pointer {
	p.hdr.id = uint32(ParameterIDAntishake)
	p.hdr.size = uint32(unsafe.Sizeof(ParameterAntishake{}))
	return unsafe.Pointer(&p.hdr)
}

func (p *ParameterImageEffectParameters) header() unsafe.Pointer {
	p.hdr.id = uint32(ParameterIDImageEffectParameters)
	p.hdr.size = uint32(unsafe.Sizeof(ParameterImageEffectParameters{}))
	return unsafe.Pointer(&p.hdr)
}

func (p *ParameterCameraBurstCapture) header() unsafe.Pointer {
	p.hdr.id = uint32(ParameterIDCameraBurstCapture)
	p.hdr.size = uint32(unsafe.Sizeof(ParameterCameraBurstCapture{}))
	return unsafe.Pointer(&p.hdr)
}

func (p *ParameterCameraMinISO) header() unsafe.Pointer {
	p.hdr.id = uint32(ParameterIDCameraMinISO)
	p.hdr.size = uint32(unsafe.Sizeof(ParameterCameraMinISO{}))
	return unsafe.Pointer(&p.hdr)
}

func (p *ParameterCameraUseCase) header() unsafe.Pointer {
	p.hdr.id = uint32(ParameterIDCameraUseCase)
	p.hdr.size = uint32(unsafe.Sizeof(ParameterCameraUseCase{}))
	return unsafe.Pointer(&p.hdr)
}

func (p *ParameterCaptureStatsPass) header() unsafe.Pointer {
	p.hdr.id = uint32(ParameterIDCaptureStatsPass)
	p.hdr.size = uint32(unsafe.Sizeof(ParameterCaptureStatsPass{}))
	return unsafe.Pointer(&p.hdr)
}

func (p *ParameterCameraCustomSensorConfig) header() unsafe.Pointer {
	p.hdr.id = uint32(ParameterIDCameraCustomSensorConfig)
	p.hdr.size = uint32(unsafe.Sizeof(ParameterCameraCustomSensorConfig{}))
	return unsafe.Pointer(&p.hdr)
}

func (p *ParameterEnableRegisterFile) header() unsafe.Pointer {
	p.hdr.id = uint32(ParameterIDEnableRegisterFile)
	p.hdr.size = uint32(unsafe.Sizeof(ParameterEnableRegisterFile{}))
	return unsafe.Pointer(&p.hdr)
}

func (p *ParameterRegisterFailIsFatal) header() unsafe.Pointer {
	p.hdr.id = uint32(ParameterIDRegisterFailIsFatal)
	p.hdr.size = uint32(unsafe.Sizeof(ParameterRegisterFailIsFatal{}))
	return unsafe.Pointer(&p.hdr)
}

func (p *ParameterConfigfileRegisters) header() unsafe.Pointer {
	p.hdr.id = uint32(ParameterIDConfigfileRegisters)
	p.hdr.size = uint32(unsafe.Sizeof(ParameterConfigfileRegisters{}))
	return unsafe.Pointer(&p.hdr)
}

func (p *ParameterConfigfileChunkRegisters) header() unsafe.Pointer {
	p.hdr.id = uint32(ParameterIDConfigfileChunkRegisters)
	p.hdr.size = uint32(unsafe.Sizeof(ParameterConfigfileChunkRegisters{}))
	return unsafe.Pointer(&p.hdr)
}

func (p *ParameterJPEGAttachLog) header() unsafe.Pointer {
	p.hdr.id = uint32(ParameterIDJPEGAttachLog)
	p.hdr.size = uint32(unsafe.Sizeof(ParameterJPEGAttachLog{}))
	return unsafe.Pointer(&p.hdr)
}

func (p *ParameterZeroShutterLag) header() unsafe.Pointer {
	p.hdr.id = uint32(ParameterIDZeroShutterLag)
	p.hdr.size = uint32(unsafe.Sizeof(ParameterZeroShutterLag{}))
	return unsafe.Pointer(&p.hdr)
}

func (p *ParameterFPSRange) header() unsafe.Pointer {
	p.hdr.id = uint32(ParameterIDFPSRange)
	p.hdr.size = uint32(unsafe.Sizeof(ParameterFPSRange{}))
	return unsafe.Pointer(&p.hdr)
}

func (p *ParameterCaptureExposureComp) header() unsafe.Pointer {
	p.hdr.id = uint32(ParameterIDCaptureExposureComp)
	p.hdr.size = uint32(unsafe.Sizeof(ParameterCaptureExposureComp{}))
	return unsafe.Pointer(&p.hdr)
}

func (p *ParameterSwSharpenDisable) header() unsafe.Pointer {
	p.hdr.id = uint32(ParameterIDSwSharpenDisable)
	p.hdr.size = uint32(unsafe.Sizeof(ParameterSwSharpenDisable{}))
	return unsafe.Pointer(&p.hdr)
}

func (p *ParameterFlashRequired) header() unsafe.Pointer {
	p.hdr.id = uint32(ParameterIDFlashRequired)
	p.hdr.size = uint32(unsafe.Sizeof(ParameterFlashRequired{}))
	return unsafe.Pointer(&p.hdr)
}

func (p *ParameterSwSaturationDisable) header() unsafe.Pointer {
	p.hdr.id = uint32(ParameterIDSwSaturationDisable)
	p.hdr.size = uint32(unsafe.Sizeof(ParameterSwSaturationDisable{}))
	return unsafe.Pointer(&p.hdr)
}

func (p *ParameterShutterSpeed) header() unsafe.Pointer {
	p.hdr.id = uint32(ParameterIDShutterSpeed)
	p.hdr.size = uint32(unsafe.Sizeof(ParameterShutterSpeed{}))
	return unsafe.Pointer(&p.hdr)
}

func (p *ParameterCustomAWBGains) header() unsafe.Pointer {
	p.hdr.id = uint32(ParameterIDCustomAWBGains)
	p.hdr.size = uint32(unsafe.Sizeof(ParameterCustomAWBGains{}))
	return unsafe.Pointer(&p.hdr)
}

func (p *ParameterCameraSettings) header() unsafe.Pointer {
	p.hdr.id = uint32(ParameterIDCameraSettings)
	p.hdr.size = uint32(unsafe.Sizeof(ParameterCameraSettings{}))
	return unsafe.Pointer(&p.hdr)
}

func (p *ParameterPrivacyIndicator) header() unsafe.Pointer {
	p.hdr.id = uint32(ParameterIDPrivacyIndicator)
	p.hdr.size = uint32(unsafe.Sizeof(ParameterPrivacyIndicator{}))
	return unsafe.Pointer(&p.hdr)
}

func (p *ParameterVideoDenoise) header() unsafe.Pointer {
	p.hdr.id = uint32(ParameterIDVideoDenoise)
	p.hdr.size = uint32(unsafe.Sizeof(ParameterVideoDenoise{}))
	return unsafe.Pointer(&p.hdr)
}

func (p *ParameterStillsDenoise) header() unsafe.Pointer {
	p.hdr.id = uint32(ParameterIDStillsDenoise)
	p.hdr.size = uint32(unsafe.Sizeof(ParameterStillsDenoise{}))
	return unsafe.Pointer(&p.hdr)
}

func (p *ParameterAnnotate) header() unsafe.Pointer {
	p.hdr.id = uint32(ParameterIDAnnotate)
	p.hdr.size = uint32(unsafe.Sizeof(ParameterAnnotate{}))
	return unsafe.Pointer(&p.hdr)
}

func (p *ParameterAnnotateV2) header() unsafe.Pointer {
	p.hdr.id = uint32(ParameterIDAnnotateV2)
	p.hdr.size = uint32(unsafe.Sizeof(ParameterAnnotateV2{}))
	return unsafe.Pointer(&p.hdr)
}

func (p *ParameterAnnotateV3) header() unsafe.Pointer {
	p.hdr.id = uint32(ParameterIDAnnotateV3)
	p.hdr.size = uint32(unsafe.Sizeof(ParameterAnnotateV3{}))
	return unsafe.Pointer(&p.hdr)
}

func (p *ParameterStereoscopicMode) header() unsafe.Pointer {
	p.hdr.id = uint32(ParameterIDStereoscopicMode)
	p.hdr.size = uint32(unsafe.Sizeof(ParameterStereoscopicMode{}))
	return unsafe.Pointer(&p.hdr)
}

func (p *ParameterCameraInterface) header() unsafe.Pointer {
	p.hdr.id = uint32(ParameterIDCameraInterface)
	p.hdr.size = uint32(unsafe.Sizeof(ParameterCameraInterface{}))
	return unsafe.Pointer(&p.hdr)
}

func (p *ParameterCameraClockingMode) header() unsafe.Pointer {
	p.hdr.id = uint32(ParameterIDCameraClockingMode)
	p.hdr.size = uint32(unsafe.Sizeof(ParameterCameraClockingMode{}))
	return unsafe.Pointer(&p.hdr)
}

func (p *ParameterCameraRxConfig) header() unsafe.Pointer {
	p.hdr.id = uint32(ParameterIDCameraRxConfig)
	p.hdr.size = uint32(unsafe.Sizeof(ParameterCameraRxConfig{}))
	return unsafe.Pointer(&p.hdr)
}

func (p *ParameterCameraRxTiming) header() unsafe.Pointer {
	p.hdr.id = uint32(ParameterIDCameraRxTiming)
	p.hdr.size = uint32(unsafe.Sizeof(ParameterCameraRxTiming{}))
	return unsafe.Pointer(&p.hdr)
}

func (p *ParameterDPFConfig) header() unsafe.Pointer {
	p.hdr.id = uint32(ParameterIDDPFConfig)
	p.hdr.size = uint32(unsafe.Sizeof(ParameterDPFConfig{}))
	return unsafe.Pointer(&p.hdr)
}

func (p *ParameterJPEGRestartInterval) header() unsafe.Pointer {
	p.hdr.id = uint32(ParameterIDJPEGRestartInterval)
	p.hdr.size = uint32(unsafe.Sizeof(ParameterJPEGRestartInterval{}))
	return unsafe.Pointer(&p.hdr)
}

func (p *ParameterCameraIspBlockOverride) header() unsafe.Pointer {
	p.hdr.id = uint32(ParameterIDCameraIspBlockOverride)
	p.hdr.size = uint32(unsafe.Sizeof(ParameterCameraIspBlockOverride{}))
	return unsafe.Pointer(&p.hdr)
}

func (p *ParameterLensShadingOverride) header() unsafe.Pointer {
	p.hdr.id = uint32(ParameterIDLensShadingOverride)
	p.hdr.size = uint32(unsafe.Sizeof(ParameterLensShadingOverride{}))
	return unsafe.Pointer(&p.hdr)
}

func (p *ParameterBlackLevel) header() unsafe.Pointer {
	p.hdr.id = uint32(ParameterIDBlackLevel)
	p.hdr.size = uint32(unsafe.Sizeof(ParameterBlackLevel{}))
	return unsafe.Pointer(&p.hdr)
}

func (p *ParameterResizeParams) header() unsafe.Pointer {
	p.hdr.id = uint32(ParameterIDResizeParams)
	p.hdr.size = uint32(unsafe.Sizeof(ParameterResizeParams{}))
	return unsafe.Pointer(&p.hdr)
}

func (p *ParameterCrop) header() unsafe.Pointer {
	p.hdr.id = uint32(ParameterIDCrop)
	p.hdr.size = uint32(unsafe.Sizeof(ParameterCrop{}))
	return unsafe.Pointer(&p.hdr)
}

func (p *ParameterOutputShift) header() unsafe.Pointer {
	p.hdr.id = uint32(ParameterIDOutputShift)
	p.hdr.size = uint32(unsafe.Sizeof(ParameterOutputShift{}))
	return unsafe.Pointer(&p.hdr)
}

func (p *ParameterCCMShift) header() unsafe.Pointer {
	p.hdr.id = uint32(ParameterIDCCMShift)
	p.hdr.size = uint32(unsafe.Sizeof(ParameterCCMShift{}))
	return unsafe.Pointer(&p.hdr)
}

func (p *ParameterCustomCCM) header() unsafe.Pointer {
	p.hdr.id = uint32(ParameterIDCustomCCM)
	p.hdr.size = uint32(unsafe.Sizeof(ParameterCustomCCM{}))
	return unsafe.Pointer(&p.hdr)
}

// Types requiring memory preallocation
func (p ParameterURI) header() unsafe.Pointer {
	hdr := (*_parameterHeader)(unsafe.Pointer(&p.raw[0]))
	hdr.id = uint32(ParameterIDURI)
	hdr.size = uint32(len(p.raw))
	return unsafe.Pointer(hdr)
}

func (p ParameterSupportedEncodings) header() unsafe.Pointer {
	hdr := (*_parameterHeader)(unsafe.Pointer(&p.raw[0]))
	hdr.id = uint32(ParameterIDSupportedEncodings)
	hdr.size = uint32(len(p.raw))
	return unsafe.Pointer(hdr)
}

func (p ParameterEXIF) header() unsafe.Pointer {
	hdr := (*_parameterHeader)(unsafe.Pointer(&p.raw[0]))
	hdr.id = uint32(ParameterIDEXIF)
	hdr.size = uint32(len(p.raw))
	return unsafe.Pointer(hdr)
}

func (p ParameterFaceTrackResults) header() unsafe.Pointer {
	hdr := (*_parameterHeader)(unsafe.Pointer(&p.raw[0]))
	hdr.id = uint32(ParameterIDFaceTrackResults)
	hdr.size = uint32(len(p.raw))
	return unsafe.Pointer(hdr)
}

func (p ParameterFocusRegions) header() unsafe.Pointer {
	hdr := (*_parameterHeader)(unsafe.Pointer(&p.raw[0]))
	hdr.id = uint32(ParameterIDFocusRegions)
	hdr.size = uint32(len(p.raw))
	return unsafe.Pointer(hdr)
}

// Array types
func (p parameterURI) data() []byte {
	hdr := (*_parameterHeader)(unsafe.Pointer(&p[0]))
	ln := (int(hdr.size) - int(unsafe.Offsetof(_parameterURI{}.uri))) / int(unsafe.Sizeof(byte(0)))
	ptr := unsafe.Pointer(uintptr(unsafe.Pointer(&p[0])) + unsafe.Offsetof(_parameterURI{}.uri))
	return (*[math.MaxInt32 / unsafe.Sizeof(byte(0))]byte)(ptr)[:int(ln):int(ln)]
}

func newParameterURI(ln int) parameterURI {
	l := ln*int(unsafe.Sizeof(byte(0))) + int(unsafe.Offsetof(_parameterURI{}.uri))
	return parameterURI(make([]byte, l, l))
}

func newParameterURISlice(data []byte) parameterURI {
	ln := len(data)
	p := newParameterURI(ln)
	ptr := unsafe.Pointer(uintptr(unsafe.Pointer(&p[0])) + unsafe.Offsetof(_parameterURI{}.uri))
	src := unsafe.Pointer(&data[0])
	C.memcpy(ptr, src, C.size_t(ln*int(unsafe.Sizeof(byte(0)))))
	return p
}

func (p parameterEncoding) data() []uint32 {
	hdr := (*_parameterHeader)(unsafe.Pointer(&p[0]))
	ln := (int(hdr.size) - int(unsafe.Offsetof(_parameterEncoding{}.encoding))) / int(unsafe.Sizeof(uint32(0)))
	ptr := unsafe.Pointer(uintptr(unsafe.Pointer(&p[0])) + unsafe.Offsetof(_parameterEncoding{}.encoding))
	return (*[math.MaxInt32 / unsafe.Sizeof(uint32(0))]uint32)(ptr)[:int(ln):int(ln)]
}

func newParameterEncoding(ln int) parameterEncoding {
	l := ln*int(unsafe.Sizeof(uint32(0))) + int(unsafe.Offsetof(_parameterEncoding{}.encoding))
	return parameterEncoding(make([]byte, l, l))
}

func newParameterEncodingSlice(data []uint32) parameterEncoding {
	ln := len(data)
	p := newParameterEncoding(ln)
	ptr := unsafe.Pointer(uintptr(unsafe.Pointer(&p[0])) + unsafe.Offsetof(_parameterEncoding{}.encoding))
	src := unsafe.Pointer(&data[0])
	C.memcpy(ptr, src, C.size_t(ln*int(unsafe.Sizeof(uint32(0)))))
	return p
}

func (p parameterEXIF) data() []uint8 {
	hdr := (*_parameterHeader)(unsafe.Pointer(&p[0]))
	ln := (int(hdr.size) - int(unsafe.Offsetof(_parameterEXIF{}.data))) / int(unsafe.Sizeof(uint8(0)))
	ptr := unsafe.Pointer(uintptr(unsafe.Pointer(&p[0])) + unsafe.Offsetof(_parameterEXIF{}.data))
	return (*[math.MaxInt32 / unsafe.Sizeof(uint8(0))]uint8)(ptr)[:int(ln):int(ln)]
}

func newParameterEXIF(ln int) parameterEXIF {
	l := ln*int(unsafe.Sizeof(uint8(0))) + int(unsafe.Offsetof(_parameterEXIF{}.data))
	return parameterEXIF(make([]byte, l, l))
}

func newParameterEXIFSlice(data []uint8) parameterEXIF {
	ln := len(data)
	p := newParameterEXIF(ln)
	ptr := unsafe.Pointer(uintptr(unsafe.Pointer(&p[0])) + unsafe.Offsetof(_parameterEXIF{}.data))
	src := unsafe.Pointer(&data[0])
	C.memcpy(ptr, src, C.size_t(ln*int(unsafe.Sizeof(uint8(0)))))
	return p
}

func (p parameterEXIF) header() *ParameterEXIFHeader {
	return (*ParameterEXIFHeader)(unsafe.Pointer(&p[0]))
}

func (p parameterFaceTrackResults) data() []ParameterFaceTrackFace {
	hdr := (*_parameterHeader)(unsafe.Pointer(&p[0]))
	ln := (int(hdr.size) - int(unsafe.Offsetof(_parameterFaceTrackResults{}.faces))) / int(unsafe.Sizeof(ParameterFaceTrackFace{}))
	ptr := unsafe.Pointer(uintptr(unsafe.Pointer(&p[0])) + unsafe.Offsetof(_parameterFaceTrackResults{}.faces))
	return (*[math.MaxInt32 / unsafe.Sizeof(ParameterFaceTrackFace{})]ParameterFaceTrackFace)(ptr)[:int(ln):int(ln)]
}

func newParameterFaceTrackResults(ln int) parameterFaceTrackResults {
	l := ln*int(unsafe.Sizeof(ParameterFaceTrackFace{})) + int(unsafe.Offsetof(_parameterFaceTrackResults{}.faces))
	return parameterFaceTrackResults(make([]byte, l, l))
}

func newParameterFaceTrackResultsSlice(data []ParameterFaceTrackFace) parameterFaceTrackResults {
	ln := len(data)
	p := newParameterFaceTrackResults(ln)
	ptr := unsafe.Pointer(uintptr(unsafe.Pointer(&p[0])) + unsafe.Offsetof(_parameterFaceTrackResults{}.faces))
	src := unsafe.Pointer(&data[0])
	C.memcpy(ptr, src, C.size_t(ln*int(unsafe.Sizeof(ParameterFaceTrackFace{}))))
	return p
}

func (p parameterFaceTrackResults) header() *ParameterFaceTrackResultsHeader {
	return (*ParameterFaceTrackResultsHeader)(unsafe.Pointer(&p[0]))
}

func (p parameterFocusRegions) data() []ParameterFocusRegion {
	hdr := (*_parameterHeader)(unsafe.Pointer(&p[0]))
	ln := (int(hdr.size) - int(unsafe.Offsetof(_parameterFocusRegions{}.regions))) / int(unsafe.Sizeof(ParameterFocusRegion{}))
	ptr := unsafe.Pointer(uintptr(unsafe.Pointer(&p[0])) + unsafe.Offsetof(_parameterFocusRegions{}.regions))
	return (*[math.MaxInt32 / unsafe.Sizeof(ParameterFocusRegion{})]ParameterFocusRegion)(ptr)[:int(ln):int(ln)]
}

func newParameterFocusRegions(ln int) parameterFocusRegions {
	l := ln*int(unsafe.Sizeof(ParameterFocusRegion{})) + int(unsafe.Offsetof(_parameterFocusRegions{}.regions))
	return parameterFocusRegions(make([]byte, l, l))
}

func newParameterFocusRegionsSlice(data []ParameterFocusRegion) parameterFocusRegions {
	ln := len(data)
	p := newParameterFocusRegions(ln)
	ptr := unsafe.Pointer(uintptr(unsafe.Pointer(&p[0])) + unsafe.Offsetof(_parameterFocusRegions{}.regions))
	src := unsafe.Pointer(&data[0])
	C.memcpy(ptr, src, C.size_t(ln*int(unsafe.Sizeof(ParameterFocusRegion{}))))
	return p
}

func (p parameterFocusRegions) header() *ParameterFocusRegionsHeader {
	return (*ParameterFocusRegionsHeader)(unsafe.Pointer(&p[0]))
}

