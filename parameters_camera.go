package mmal

/*
#include <stdlib.h>
#include <string.h>
#include <interface/mmal/mmal.h>
#include <interface/mmal/mmal_parameters.h>
*/
import "C"

import (
	"strings"
)

type ParameterThumbnailConfiguration _parameterThumbnailConfig
type ParameterRotation _parameterInt32
type ParameterEXIFDisable _parameterBoolean
type ParameterEXIF struct {
	raw parameterEXIF
}
type ParameterAWBMode _parameterAWBMode
type ParameterImageEffect _parameterImageFX
type ParameterColourEffect _parameterColourFX
type ParameterFlickerAvoid _parameterFlickerAvoid
type ParameterFlash _parameterFlash
type ParameterRedeye _parameterRedeye
type ParameterFocus _parameterFocus
type ParameterExposureCompInt _parameterInt32
type ParameterExposureCompRat _parameterRational
type ParameterZoom _parameterScalefactor
type ParameterMirror _parameterMirror
type ParameterCameraNum _parameterUint32
type ParameterCapture _parameterBoolean
type ParameterExposureMode _parameterExposureMode
type ParameterExpMeteringMode _parameterExposureMeteringMode
type ParameterFocusStatus _parameterFocusStatus
type ParameterCameraConfig _parameterCameraConfig
type ParameterCaptureStatus _parameterCaptureStatus
type ParameterFaceTrack _parameterFaceTrack
type ParameterDrawBoxFacesAndFocus _parameterBoolean
type ParameterJPEGQFactor _parameterUint32
type ParameterFrameRate _parameterFrameRate
type ParameterUseSTC _parameterCameraSTCMode
type ParameterCameraInfo _parameterCameraInfo
type ParameterVideoStabilisation _parameterBoolean

type ParameterFaceTrackResults struct {
	raw parameterFaceTrackResults
}

type ParameterEnableRawCapture _parameterBoolean

type ParameterDPFFile struct {
	raw parameterURI
}

type ParameterEnableDPFFile _parameterBoolean
type ParameterDPFFailIsFatal _parameterBoolean
type ParameterCaptureMode _parameterCaptureMode

type ParameterFocusRegions struct {
	raw parameterFocusRegions
}

type ParameterInputCrop _parameterInputCrop
type ParameterSensorInformation _parameterSensorInformation
type ParameterFlashSelect _parameterFlashSelect
type ParameterFieldOfView _parameterFieldOfView
type ParameterHighDynamicRange _parameterBoolean
type ParameterDynamicRangeCompression _parameterDRC
type ParameterAlgorithmControl _parameterAlgorithmControl
type ParameterSharpness _parameterRational
type ParameterContrast _parameterRational
type ParameterBrightness _parameterRational
type ParameterSaturation _parameterRational
type ParameterISO _parameterUint32
type ParameterAntishake _parameterBoolean
type ParameterImageEffectParameters _parameterImageFXParameters
type ParameterCameraBurstCapture _parameterBoolean
type ParameterCameraMinISO _parameterUint32
type ParameterCameraUseCase _parameterCameraUseCase
type ParameterCaptureStatsPass _parameterBoolean
type ParameterCameraCustomSensorConfig _parameterUint32
type ParameterEnableRegisterFile _parameterBoolean
type ParameterRegisterFailIsFatal _parameterBoolean
type ParameterConfigfileRegisters _parameterConfigfile
type ParameterConfigfileChunkRegisters _parameterConfigfileChunk
type ParameterJPEGAttachLog _parameterBoolean
type ParameterZeroShutterLag _parameterZeroshutterlag
type ParameterFPSRange _parameterFPSRange
type ParameterCaptureExposureComp _parameterInt32
type ParameterSwSharpenDisable _parameterBoolean
type ParameterFlashRequired _parameterBoolean
type ParameterSwSaturationDisable _parameterBoolean
type ParameterShutterSpeed _parameterUint32
type ParameterCustomAWBGains _parameterAWBGains
type ParameterCameraSettings _parameterCameraSettings
type ParameterPrivacyIndicator _parameterPrivacyIndicator
type ParameterVideoDenoise _parameterBoolean
type ParameterStillsDenoise _parameterBoolean
type ParameterAnnotate _parameterCameraAnnotate
type ParameterAnnotateV2 _parameterCameraAnnotateV2
type ParameterAnnotateV3 _parameterCameraAnnotateV3
type ParameterStereoscopicMode _parameterStereoscopicMode
type ParameterCameraInterface _parameterCameraInterface
type ParameterCameraClockingMode _parameterCameraClockingMode
type ParameterCameraRxConfig _parameterCameraRxConfig
type ParameterCameraRxTiming _parameterCameraRxTiming
type ParameterDPFConfig _parameterUint32
type ParameterJPEGRestartInterval _parameterUint32
type ParameterCameraIspBlockOverride _parameterUint32
type ParameterLensShadingOverride _parameterLensShading
type ParameterBlackLevel _parameterUint32
type ParameterResizeParams _parameterResize
type ParameterCrop _parameterCrop
type ParameterOutputShift _parameterInt32
type ParameterCCMShift _parameterInt32
type ParameterCustomCCM _parameterCustomCCM

type ParameterAlgorithmControlAlgorithms uint32

const (
	ParameterAlgorithmControlAlgorithmsFacetracking ParameterAlgorithmControlAlgorithms = iota
	ParameterAlgorithmControlAlgorithmsRedeyeReduction
	ParameterAlgorithmControlAlgorithmsVideoStabilisation
	ParameterAlgorithmControlAlgorithmsWriteRaw
	ParameterAlgorithmControlAlgorithmsVideoDenoise
	ParameterAlgorithmControlAlgorithmsStillsDenoise
	ParameterAlgorithmControlAlgorithmsTemporalDenoise
	ParameterAlgorithmControlAlgorithmsAntishake
	ParameterAlgorithmControlAlgorithmsImageEffects
	ParameterAlgorithmControlAlgorithmsDynamicRangeCompression
	ParameterAlgorithmControlAlgorithmsFaceRecognition
	ParameterAlgorithmControlAlgorithmsFaceBeautification
	ParameterAlgorithmControlAlgorithmsSceneDetection
	ParameterAlgorithmControlAlgorithmsHighDynamicRange
	ParameterAlgorithmControlAlgorithmsMax = 0x7fffffff
)

type _parameterAlgorithmControl struct {
	hdr       _parameterHeader
	Algorithm ParameterAlgorithmControlAlgorithms
	Enabled   int32
}

type ParamAWBMode uint32

const (
	ParamAWBModeOff ParamAWBMode = iota
	ParamAWBModeAuto
	ParamAWBModeSunlight
	ParamAWBModeCloudy
	ParamAWBModeShade
	ParamAWBModeTungsten
	ParamAWBModeFluorescent
	ParamAWBModeIncandescent
	ParamAWBModeFlash
	ParamAWBModeHorizon
	ParamAWBModeMax = 0x7fffffff
)

type _parameterAWBMode struct {
	hdr   _parameterHeader
	Value ParamAWBMode
}

type _parameterAWBGains struct {
	hdr   _parameterHeader
	RGain Rational
	BGain Rational
}

type _parameterCameraAnnotate struct {
	hdr            _parameterHeader
	Enable         int32
	_text          [32]byte
	ShowShutter    int32
	ShowAnalogGain int32
	ShowLens       int32
	ShowCAF        int32
	ShowMotion     int32
}

func (p *_parameterCameraAnnotate) text() string {
	return cString(&p._text[0], len(p._text))
}

func (p *ParameterAnnotate) Text() string {
	return (*_parameterCameraAnnotate)(p).text()
}

type _parameterCameraAnnotateV2 struct {
	hdr                 _parameterHeader
	Enable              int32
	ShowShutter         int32
	ShowAnalogGain      int32
	ShowLens            int32
	ShowCAF             int32
	ShowMotion          int32
	ShowFrameNum        int32
	BlackTextBackground int32
	_text               [256]byte
}

func (p *_parameterCameraAnnotateV2) text() string {
	return cString(&p._text[0], len(p._text))
}

func (p *ParameterAnnotateV2) Text() string {
	return (*_parameterCameraAnnotateV2)(p).text()
}

type _parameterCameraAnnotateV3 struct {
	hdr                    _parameterHeader
	Enable                 int32
	ShowShutter            int32
	ShowAnalogGain         int32
	ShowLens               int32
	ShowCAF                int32
	ShowMotion             int32
	ShowFrameNum           int32
	EnableTextBackground   int32
	CustomBackgroundColour int32
	CustomBackgroundY      uint8
	CustomBackgroundU      uint8
	CustomBackgroundV      uint8
	dummy1                 uint8
	CustomTextColour       int32
	CustomTextY            uint8
	CustomTextU            uint8
	CustomTextV            uint8
	TextSize               uint8
	_text                  [256]byte
}

func (p *_parameterCameraAnnotateV3) text() string {
	return cString(&p._text[0], len(p._text))
}

func (p *ParameterAnnotateV3) Text() string {
	return (*_parameterCameraAnnotateV3)(p).text()
}

type CameraClockingMode uint32

const (
	CameraClockingModeStrobe CameraClockingMode = iota
	CameraClockingModeClock
	CameraClockingModeMax = 0x7FFFFFFF
)

type _parameterCameraClockingMode struct {
	hdr  _parameterHeader
	mode CameraClockingMode
}

type ParameterCameraConfigTimestampMode uint32

const (
	ParamTimestampModeZero     ParameterCameraConfigTimestampMode = iota // Always timestamp frames as 0
	ParamTimestampModeRawSTC                                             // Use the raw STC value for the frame timestamp
	ParamTimestampModeResetSTC                                           // Use the STC timestamp but subtract the timestamp of the first frame sent to give a zero based timestamp.
	ParamTimestampModeMax      = 0x7FFFFFFF
)

type _parameterCameraConfig struct {
	hdr                               _parameterHeader
	MaxStillsW                        uint32
	MaxStillsH                        uint32
	StillsYUV422                      uint32
	OneShotStills                     uint32
	MaxPreviewVideoW                  uint32
	MaxPreviewVideoH                  uint32
	NumPreviewVideoFrames             uint32
	StillsCaptureCircularBufferHeight uint32
	FastPreviewResume                 uint32
	UseSTCTimestamp                   ParameterCameraConfigTimestampMode
}

type ParameterCameraInfoCamera struct {
	PortID      uint32
	MaxWidth    uint32
	MaxHeight   uint32
	LensPresent int32
	cameraName  [16]byte
}

func cString(cs *byte, n int) string {
	s := C.GoStringN((*C.char)(cs), C.int(n))
	if i := strings.IndexByte(s, 0); i >= 0 {
		return s[:i]
	}
	return s
}

func (p *ParameterCameraInfoCamera) CameraName() string {
	return cString(&p.cameraName[0], len(p.cameraName))
}

type ParameterCameraInfoFlashType uint32

const (
	ParameterCameraInfoFlashTypeXenon ParameterCameraInfoFlashType = iota // Make values explicit
	ParameterCameraInfoFlashTypeLed                                       // to ensure they match
	ParameterCameraInfoFlashTypeOther                                     // values in config ini
	ParameterCameraInfoFlashTypeMax   = 0x7FFFFFFF
)

type ParameterCameraInfoFlash struct {
	FlashType ParameterCameraInfoFlashType
}

type _parameterCameraInfo struct {
	hdr        _parameterHeader
	NumCameras uint32
	NumFlashes uint32
	Cameras    [4]ParameterCameraInfoCamera
	Flashes    [2]ParameterCameraInfoFlash
}

type CameraInterface uint32

const (
	CameraInterfaceCSI2 CameraInterface = iota
	CameraInterfaceCCP2
	CameraInterfaceCPI
	CameraInterfaceMax = 0x7FFFFFFF
)

type _parameterCameraInterface struct {
	hdr  _parameterHeader
	Mode CameraInterface
}

type CameraRxConfigDecode uint32

const (
	CameraRxConfigDecodeNone CameraRxConfigDecode = iota
	CameraRxConfigDecodeDPCM8To10
	CameraRxConfigDecodeDPCM7To10
	CameraRxConfigDecodeDPCM6To10
	CameraRxConfigDecodeDPCM8To12
	CameraRxConfigDecodeDPCM7To12
	CameraRxConfigDecodeDPCM6To12
	CameraRxConfigDecodeDPCM10To14
	CameraRxConfigDecodeDPCM8To14
	CameraRxConfigDecodeDPCM12To16
	CameraRxConfigDecodeDPCM10To16
	CameraRxConfigDecodeDPCM8To16
	CameraRxConfigDecodeMax = 0x7FFFFFFF
)

type CameraRxConfigEncode uint32

const (
	CameraRxConfigEncodeNone CameraRxConfigEncode = iota
	CameraRxConfigEncodeDPCM10To8
	CameraRxConfigEncodeDPCM12To8
	CameraRxConfigEncodeDPCM14To8
	CameraRxConfigEncodeMax = 0x7FFFFFFF
)

type CameraRxConfigUnpack uint32

const (
	CameraRxConfigUnpackNone CameraRxConfigUnpack = iota
	CameraRxConfigUnpack6
	CameraRxConfigUnpack7
	CameraRxConfigUnpack8
	CameraRxConfigUnpack10
	CameraRxConfigUnpack12
	CameraRxConfigUnpack14
	CameraRxConfigUnpack16
	CameraRxConfigUnpackMax = 0x7FFFFFFF
)

type CameraRxConfigPack uint32

const (
	CameraRxConfigPackNone CameraRxConfigPack = iota
	CameraRxConfigPack8
	CameraRxConfigPack10
	CameraRxConfigPack12
	CameraRxConfigPack14
	CameraRxConfigPack16
	CameraRxConfigPackRaw10
	CameraRxConfigPackRaw12
	CameraRxConfigPackMax = 0x7FFFFFFF
)

type _parameterCameraRxConfig struct {
	hdr               _parameterHeader
	Decode            CameraRxConfigDecode
	Encode            CameraRxConfigEncode
	Unpack            CameraRxConfigUnpack
	Pack              CameraRxConfigPack
	DataLanes         uint32
	EncodeBlockLength uint32
	EmbeddedDataLines uint32
	ImageID           uint32
}

type _parameterCameraRxTiming struct {
	hdr        _parameterHeader
	Timing1    uint32
	Timing2    uint32
	Timing3    uint32
	Timing4    uint32
	Timing5    uint32
	Term1      uint32
	Term2      uint32
	CPITiming1 uint32
	CPITiming2 uint32
}

type _parameterCameraSettings struct {
	hdr           _parameterHeader
	Exposure      uint32
	AnalogGain    Rational
	DigitalGain   Rational
	AWBRedGain    Rational
	AWBBlueGain   Rational
	FocusPosition uint32
}

type CameraSTCMode uint32

const (
	ParamSTCModeOff    CameraSTCMode = iota // Frames do not have STCs, as needed in OpenMAX/IL
	ParamSTCModeRaw                         // Use raw clock STC, needed for true pause/resume support
	ParamSTCModeCooked                      // Start the STC from the start of capture, only for quick demo code
	ParamSTCModeMax    = 0x7fffffff
)

type _parameterCameraSTCMode struct {
	hdr   _parameterHeader
	Value CameraSTCMode
}

type ParamCameraUseCase uint32

const (
	ParamCameraUseCaseUnknown       ParamCameraUseCase = iota // Compromise on behaviour as use case totally unknown
	ParamCameraUseCaseStillsCapture                           // Stills capture use case
	ParamCameraUseCaseVideoCapture                            // Video encode (camcorder) use case
	ParamCameraUseCaseMax           = 0x7fffffff
)

type _parameterCameraUseCase struct {
	hdr     _parameterHeader
	UseCase ParamCameraUseCase
}

type ParameterCaptureModeMode uint32

const (
	ParamCaptureModeWaitForEnd        ParameterCaptureModeMode = iota // Resumes preview once capture is completed.
	ParamCaptureModeWaitForEndAndHold                                 // Resumes preview once capture is completed, and hold the image for subsequent reprocessing.
	/* Resumes preview as soon as possible once capture frame is received from the sensor.
	Requires fast_preview_resume to be set via MMAL_PARAMETER_CAMERA_CONFIG. */
	ParamCaptureModeResumeVFImmediately
)

type _parameterCaptureMode struct {
	hdr  _parameterHeader
	Mode ParameterCaptureModeMode
}

type ParamCaptureStatus uint32

const (
	ParamCaptureStatusNotCapturing ParamCaptureStatus = iota
	ParamCaptureStatusCaptureStarted
	ParamCaptureStatusCaptureEnded
	ParamCaptureStatusMax = 0x7FFFFFFF
)

type _parameterCaptureStatus struct {
	hdr    _parameterHeader
	Status ParamCaptureStatus
}

type ParameterCCM struct {
	CCM     [3][3]Rational
	Offsets [3]int32
}

type _parameterColourFX struct {
	hdr    _parameterHeader
	Enable int32
	U      uint32
	V      uint32
}

type _parameterCrop struct {
	hdr  _parameterHeader
	Rect Rect
}

type _parameterCustomCCM struct {
	hdr    _parameterHeader
	Enable int32
	CCM    ParameterCCM
}

type ParameterDRCStrength uint32

const (
	ParameterDRCStrengthOff ParameterDRCStrength = iota
	ParameterDRCStrengthLow
	ParameterDRCStrengthMedium
	ParameterDRCStrengthHigh
	ParameterDRCStrengthMax = 0x7fffffff
)

type _parameterDRC struct {
	hdr      _parameterHeader
	Strength ParameterDRCStrength
}

type _parameterEXIF struct {
	hdr         _parameterHeader
	KeyLen      uint32
	ValueOffset uint32
	ValueLen    uint32
	data        [1]uint8
	_           [3]byte
}

type ParamExposureMeteringMode uint32

const (
	ParamExposureMeteringModeAverage ParamExposureMeteringMode = iota
	ParamExposureMeteringModeSpot
	ParamExposureMeteringModeBacklit
	ParamExposureMeteringModeMatrix
	ParamExposureMeteringModeMax = 0x7fffffff
)

type _parameterExposureMeteringMode struct {
	hdr   _parameterHeader
	Value ParamExposureMeteringMode
}

type ParamExposureMode uint32

const (
	ParamExposureModeOff ParamExposureMode = iota
	ParamExposureModeAuto
	ParamExposureModeNight
	ParamExposureModeNightpreview
	ParamExposureModeBacklight
	ParamExposureModeSpotlight
	ParamExposureModeSports
	ParamExposureModeSnow
	ParamExposureModeBeach
	ParamExposureModeVerylong
	ParamExposureModeFixedFPS
	ParamExposureModeAntishake
	ParamExposureModeFireworks
	ParamExposureModeMax = 0x7fffffff
)

type _parameterExposureMode struct {
	hdr   _parameterHeader
	Value ParamExposureMode
}

type ParameterFaceTrackFace struct {
	FaceID    int32
	Score     int32
	FaceRect  Rect
	EyeRect   [2]Rect
	MouthRect Rect
}

type _parameterFaceTrackResults struct {
	hdr         _parameterHeader
	NumFaces    uint32
	FrameWidth  uint32
	FrameHeight uint32
	faces       [1]ParameterFaceTrackFace
}

type ParamFaceTrackMode uint32

const (
	ParamFaceDetectNone ParamFaceTrackMode = iota // Disables face detection
	ParamFaceDetectOn                             // Enables face detection
	ParamFaceDetectMax  = 0x7FFFFFFF
)

type _parameterFaceTrack struct {
	hdr        _parameterHeader
	Mode       ParamFaceTrackMode
	MaxRegions uint32
	Frames     uint32
	Quality    uint32
}

type _parameterFieldOfView struct {
	hdr  _parameterHeader
	FovH Rational
	FovV Rational
}

type _parameterFlashSelect struct {
	hdr       _parameterHeader
	FlashType ParameterCameraInfoFlashType
}

type ParamFlash uint32

const (
	ParamFlashOff ParamFlash = iota
	ParamFlashAuto
	ParamFlashOn
	ParamFlashRedeye
	ParamFlashFillin
	ParamFlashTorch
	ParamFlashMax = 0x7FFFFFFF
)

type _parameterFlash struct {
	hdr   _parameterHeader
	Value ParamFlash
}

type ParamFlickerAvoid uint32

const (
	ParamFlickerAvoidOff ParamFlickerAvoid = iota
	ParamFlickerAvoidAuto
	ParamFlickerAvoid50Hz
	ParamFlickerAvoid60Hz
	ParamFlickerAvoidMax = 0x7FFFFFFF
)

type _parameterFlickerAvoid struct {
	hdr   _parameterHeader
	Value ParamFlickerAvoid
}

type _parameterFocusRegions struct {
	hdr         _parameterHeader
	NumRegions  uint32
	LockToFaces int32
	regions     [1]ParameterFocusRegion
}

type ParameterFocusRegionType uint32

const (
	ParameterFocusRegionTypeNormal ParameterFocusRegionType = iota // Region defines a generic region
	ParameterFocusRegionTypeFace                                   // Region defines a face
	ParameterFocusRegionTypeMax
)

type ParameterFocusRegion struct {
	Rect   Rect
	Weight uint32
	Mask   uint32
	Type   ParameterFocusRegionType
}

type ParamFocusStatus uint32

const (
	ParamFocusStatusOff ParamFocusStatus = iota
	ParamFocusStatusRequest
	ParamFocusStatusReached
	ParamFocusStatusUnableToReach
	ParamFocusStatusLost
	ParamFocusStatusCAFMoving
	ParamFocusStatusCAFSuccess
	ParamFocusStatusCAFFailed
	ParamFocusStatusManualMoving
	ParamFocusStatusManualReached
	ParamFocusStatusCAFWatching
	ParamFocusStatusCAFSceneChanged
	ParamFocusStatusMax = 0x7FFFFFFF
)

type _parameterFocusStatus struct {
	hdr    _parameterHeader
	Status ParamFocusStatus
}

type ParamFocus uint32

const (
	ParamFocusAuto ParamFocus = iota
	ParamFocusAutoNear
	ParamFocusAutoMacro
	ParamFocusCAF
	ParamFocusCAFNear
	ParamFocusFixedInfinity
	ParamFocusFixedHyperfocal
	ParamFocusFixedNear
	ParamFocusFixedMacro
	ParamFocusEDOF
	ParamFocusCAFMacro
	ParamFocusCAFFast
	ParamFocusCAFNearFast
	ParamFocusCAFMacroFast
	ParamFocusFixedCurrent
	ParamFocusMax = 0x7FFFFFFF
)

type _parameterFocus struct {
	hdr   _parameterHeader
	Value ParamFocus
}

type _parameterFPSRange struct {
	hdr     _parameterHeader
	FPSLow  Rational
	FPSHigh Rational
}

type ParamImageFX uint32

const (
	ParamImageFXNone ParamImageFX = iota
	ParamImageFXNegative
	ParamImageFXSolarize
	ParamImageFXPosterize
	ParamImageFXWhiteboard
	ParamImageFXBlackboard
	ParamImageFXSketch
	ParamImageFXDenoise
	ParamImageFXEmboss
	ParamImageFXOilpaint
	ParamImageFXHatch
	ParamImageFXGpen
	ParamImageFXPastel
	ParamImageFXWatercolour
	ParamImageFXFilm
	ParamImageFXBlur
	ParamImageFXSaturation
	ParamImageFXColourswap
	ParamImageFXWashedout
	ParamImageFXPosterise
	ParamImageFXColourpoint
	ParamImageFXColourbalance
	ParamImageFXCartoon
	ParamImageFXDeinterlaceDouble
	ParamImageFXDeinterlaceAdv
	ParamImageFXDeinterlaceFast
	ParamImageFXMax = 0x7fffffff
)

type _parameterImageFXParameters struct {
	hdr             _parameterHeader
	Effect          ParamImageFX
	NumEffectParams uint32
	EffectParameter [6]uint32
}

type _parameterImageFX struct {
	hdr   _parameterHeader
	Value ParamImageFX
}

type _parameterInputCrop struct {
	hdr  _parameterHeader
	Rect Rect
}

type _parameterLensShading struct {
	hdr            _parameterHeader
	Enabled        int32
	GridCellSize   uint32
	GridWidth      uint32
	GridStride     uint32
	GridHeight     uint32
	MemHandleTable uint32
	RefTransform   uint32
}

type ParamPrivacyIndicator uint32

const (
	ParameterPrivacyIndicatorOff ParamPrivacyIndicator = iota // Indicator will be off.
	/* Indicator will come on just after a stills capture and
	and remain on for 2seconds, or will be on whilst output[1]
	is actively producing images. */
	ParameterPrivacyIndicatorOn
	/* Turns indicator of for 2s independent of capture status.
	Set this mode repeatedly to keep the indicator on for a
	longer period. */
	ParameterPrivacyIndicatorForceOn
	ParameterPrivacyIndicatorMax = 0x7fffffff
)

type _parameterPrivacyIndicator struct {
	hdr  _parameterHeader
	Mode ParamPrivacyIndicator
}

type ParamRedeye uint32

const (
	ParamRedeyeOff ParamRedeye = iota
	ParamRedeyeOn
	ParamRedeyeSimple
	ParamRedeyeMax = 0x7FFFFFFF
)

type _parameterRedeye struct {
	hdr   _parameterHeader
	Value ParamRedeye
}

type ResizeMode uint32

const (
	ResizeNone ResizeMode = iota
	ResizeCrop
	ResizeBox
	ResizeBytes
	ResizeDummy = 0x7FFFFFFF
)

type _parameterResize struct {
	hdr                 _parameterHeader
	Mode                ResizeMode
	MaxWidth            uint32
	MaxHeight           uint32
	MaxBytes            uint32
	PreserveAspectRatio int32
	AllowUpscaling      int32
}

type _parameterSensorInformation struct {
	hdr            _parameterHeader
	FNumber        Rational
	FocalLength    Rational
	ModelID        uint32
	ManufacturerID uint32
	Revision       uint32
}

type StereoscopicMode uint32

const (
	StereoscopicModeNone StereoscopicMode = iota
	StereoscopicModeSideBySide
	StereoscopicModeTopBottom
	StereoscopicModeMax = 0x7FFFFFFF
)

type _parameterStereoscopicMode struct {
	hdr      _parameterHeader
	Mode     StereoscopicMode
	Decimate int32
	SwapEyes int32
}

type _parameterThumbnailConfig struct {
	hdr     _parameterHeader
	Enable  uint32
	Width   uint32
	Height  uint32
	Quality uint32
}

type _parameterZeroshutterlag struct {
	hdr                _parameterHeader
	ZeroShutterLagMode int32
	ConcurrentCapture  int32
}

type parameterEXIF []byte

func (p *ParameterEXIF) Data() []uint8 {
	return p.raw.data()
}

func NewParameterEXIF(ln int) *ParameterEXIF {
	return &ParameterEXIF{newParameterEXIF(ln)}
}

func NewParameterEXIFSlice(data []uint8) *ParameterEXIF {
	return &ParameterEXIF{newParameterEXIFSlice(data)}
}

func NewParameterEXIFString(s string) *ParameterEXIF {
	return &ParameterEXIF{newParameterEXIFSlice(append([]uint8(s), 0))}
}

type ParameterEXIFHeader _parameterEXIF

func (p *ParameterEXIF) Header() *ParameterEXIFHeader {
	return p.raw.header()
}

type parameterFaceTrackResults []byte

func (p *ParameterFaceTrackResults) Data() []ParameterFaceTrackFace {
	return p.raw.data()
}

func NewParameterFaceTrackResults(ln int) *ParameterFaceTrackResults {
	return &ParameterFaceTrackResults{newParameterFaceTrackResults(ln)}
}

func NewParameterFaceTrackResultsSlice(data []ParameterFaceTrackFace) *ParameterFaceTrackResults {
	return &ParameterFaceTrackResults{newParameterFaceTrackResultsSlice(data)}
}

type ParameterFaceTrackResultsHeader _parameterFaceTrackResults

func (p *ParameterFaceTrackResults) Header() *ParameterFaceTrackResultsHeader {
	return p.raw.header()
}

type parameterFocusRegions []byte

func (p *ParameterFocusRegions) Data() []ParameterFocusRegion {
	return p.raw.data()
}

func NewParameterFocusRegions(ln int) *ParameterFocusRegions {
	return &ParameterFocusRegions{newParameterFocusRegions(ln)}
}

func NewParameterFocusRegionsSlice(data []ParameterFocusRegion) *ParameterFocusRegions {
	return &ParameterFocusRegions{newParameterFocusRegionsSlice(data)}
}

type ParameterFocusRegionsHeader _parameterFocusRegions

func (p *ParameterFocusRegions) Header() *ParameterFocusRegionsHeader {
	return p.raw.header()
}

// for sizeof and alignof test
type _MMAL_PARAMETER_THUMBNAIL_CONFIG C.MMAL_PARAMETER_THUMBNAIL_CONFIG_T
type _MMAL_PARAMETER_EXIF C.MMAL_PARAMETER_EXIF_T
type _MMAL_PARAMETER_EXPOSUREMODE C.MMAL_PARAMETER_EXPOSUREMODE_T
type _MMAL_PARAMETER_EXPOSUREMETERINGMODE C.MMAL_PARAMETER_EXPOSUREMETERINGMODE_T
type _MMAL_PARAMETER_AWBMODE C.MMAL_PARAMETER_AWBMODE_T
type _MMAL_PARAMETER_IMAGEFX C.MMAL_PARAMETER_IMAGEFX_T
type _MMAL_PARAMETER_IMAGEFX_PARAMETERS C.MMAL_PARAMETER_IMAGEFX_PARAMETERS_T
type _MMAL_PARAMETER_COLOURFX C.MMAL_PARAMETER_COLOURFX_T
type _MMAL_PARAMETER_CAMERA_STC_MODE C.MMAL_PARAMETER_CAMERA_STC_MODE_T
type _MMAL_PARAMETER_FLICKERAVOID C.MMAL_PARAMETER_FLICKERAVOID_T
type _MMAL_PARAMETER_FLASH C.MMAL_PARAMETER_FLASH_T
type _MMAL_PARAMETER_REDEYE C.MMAL_PARAMETER_REDEYE_T
type _MMAL_PARAMETER_FOCUS C.MMAL_PARAMETER_FOCUS_T
type _MMAL_PARAMETER_CAPTURE_STATUS C.MMAL_PARAMETER_CAPTURE_STATUS_T
type _MMAL_PARAMETER_FOCUS_STATUS C.MMAL_PARAMETER_FOCUS_STATUS_T
type _MMAL_PARAMETER_FACE_TRACK C.MMAL_PARAMETER_FACE_TRACK_T
type _MMAL_PARAMETER_FACE_TRACK_FACE C.MMAL_PARAMETER_FACE_TRACK_FACE_T
type _MMAL_PARAMETER_FACE_TRACK_RESULTS C.MMAL_PARAMETER_FACE_TRACK_RESULTS_T
type _MMAL_PARAMETER_CAMERA_CONFIG C.MMAL_PARAMETER_CAMERA_CONFIG_T
type _MMAL_PARAMETER_CAMERA_INFO_CAMERA C.MMAL_PARAMETER_CAMERA_INFO_CAMERA_T
type _MMAL_PARAMETER_CAMERA_INFO_FLASH C.MMAL_PARAMETER_CAMERA_INFO_FLASH_T
type _MMAL_PARAMETER_CAMERA_INFO C.MMAL_PARAMETER_CAMERA_INFO_T
type _MMAL_PARAMETER_CAPTUREMODE C.MMAL_PARAMETER_CAPTUREMODE_T
type _MMAL_PARAMETER_FOCUS_REGION C.MMAL_PARAMETER_FOCUS_REGION_T
type _MMAL_PARAMETER_FOCUS_REGIONS C.MMAL_PARAMETER_FOCUS_REGIONS_T
type _MMAL_PARAMETER_INPUT_CROP C.MMAL_PARAMETER_INPUT_CROP_T
type _MMAL_PARAMETER_SENSOR_INFORMATION C.MMAL_PARAMETER_SENSOR_INFORMATION_T
type _MMAL_PARAMETER_FLASH_SELECT C.MMAL_PARAMETER_FLASH_SELECT_T
type _MMAL_PARAMETER_FIELD_OF_VIEW C.MMAL_PARAMETER_FIELD_OF_VIEW_T
type _MMAL_PARAMETER_DRC C.MMAL_PARAMETER_DRC_T
type _MMAL_PARAMETER_ALGORITHM_CONTROL C.MMAL_PARAMETER_ALGORITHM_CONTROL_T
type _MMAL_PARAMETER_CAMERA_USE_CASE C.MMAL_PARAMETER_CAMERA_USE_CASE_T
type _MMAL_PARAMETER_FPS_RANGE C.MMAL_PARAMETER_FPS_RANGE_T
type _MMAL_PARAMETER_ZEROSHUTTERLAG C.MMAL_PARAMETER_ZEROSHUTTERLAG_T
type _MMAL_PARAMETER_AWB_GAINS C.MMAL_PARAMETER_AWB_GAINS_T
type _MMAL_PARAMETER_CAMERA_SETTINGS C.MMAL_PARAMETER_CAMERA_SETTINGS_T
type _MMAL_PARAMETER_PRIVACY_INDICATOR C.MMAL_PARAMETER_PRIVACY_INDICATOR_T
type _MMAL_PARAMETER_CAMERA_ANNOTATE C.MMAL_PARAMETER_CAMERA_ANNOTATE_T
type _MMAL_PARAMETER_CAMERA_ANNOTATE_V2 C.MMAL_PARAMETER_CAMERA_ANNOTATE_V2_T
type _MMAL_PARAMETER_CAMERA_ANNOTATE_V3 C.MMAL_PARAMETER_CAMERA_ANNOTATE_V3_T
type _MMAL_PARAMETER_STEREOSCOPIC_MODE C.MMAL_PARAMETER_STEREOSCOPIC_MODE_T
type _MMAL_PARAMETER_CAMERA_INTERFACE C.MMAL_PARAMETER_CAMERA_INTERFACE_T
type _MMAL_PARAMETER_CAMERA_CLOCKING_MODE C.MMAL_PARAMETER_CAMERA_CLOCKING_MODE_T
type _MMAL_PARAMETER_CAMERA_RX_CONFIG C.MMAL_PARAMETER_CAMERA_RX_CONFIG_T
type _MMAL_PARAMETER_CAMERA_RX_TIMING C.MMAL_PARAMETER_CAMERA_RX_TIMING_T
type _MMAL_PARAMETER_LENS_SHADING C.MMAL_PARAMETER_LENS_SHADING_T
type _MMAL_PARAMETER_RESIZE C.MMAL_PARAMETER_RESIZE_T
type _MMAL_PARAMETER_CROP C.MMAL_PARAMETER_CROP_T
type _MMAL_PARAMETER_CCM C.MMAL_PARAMETER_CCM_T
type _MMAL_PARAMETER_CUSTOM_CCM C.MMAL_PARAMETER_CUSTOM_CCM_T
