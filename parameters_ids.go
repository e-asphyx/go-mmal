package mmal

/*
#include <stdlib.h>
#include <string.h>
#include <interface/mmal/mmal.h>
#include <interface/mmal/mmal_parameters.h>
*/
import "C"

const (
	ParameterIDChangeEventRequest       = C.MMAL_PARAMETER_CHANGE_EVENT_REQUEST
	ParameterIDZeroCopy                 = C.MMAL_PARAMETER_ZERO_COPY
	ParameterIDBufferRequirements       = C.MMAL_PARAMETER_BUFFER_REQUIREMENTS
	ParameterIDStatistics               = C.MMAL_PARAMETER_STATISTICS
	ParameterIDCoreStatistics           = C.MMAL_PARAMETER_CORE_STATISTICS
	ParameterIDMemUsage                 = C.MMAL_PARAMETER_MEM_USAGE
	ParameterIDBufferFlagFilter         = C.MMAL_PARAMETER_BUFFER_FLAG_FILTER
	ParameterIDSeek                     = C.MMAL_PARAMETER_SEEK
	ParameterIDPowermonEnable           = C.MMAL_PARAMETER_POWERMON_ENABLE
	ParameterIDLogging                  = C.MMAL_PARAMETER_LOGGING
	ParameterIDSystemTime               = C.MMAL_PARAMETER_SYSTEM_TIME
	ParameterIDNoImagePadding           = C.MMAL_PARAMETER_NO_IMAGE_PADDING
	ParameterIDLockstepEnable           = C.MMAL_PARAMETER_LOCKSTEP_ENABLE
	ParameterIDThumbnailConfiguration   = C.MMAL_PARAMETER_THUMBNAIL_CONFIGURATION
	ParameterIDRotation                 = C.MMAL_PARAMETER_ROTATION
	ParameterIDEXIFDisable              = C.MMAL_PARAMETER_EXIF_DISABLE
	ParameterIDAWBMode                  = C.MMAL_PARAMETER_AWB_MODE
	ParameterIDImageEffect              = C.MMAL_PARAMETER_IMAGE_EFFECT
	ParameterIDColourEffect             = C.MMAL_PARAMETER_COLOUR_EFFECT
	ParameterIDFlickerAvoid             = C.MMAL_PARAMETER_FLICKER_AVOID
	ParameterIDFlash                    = C.MMAL_PARAMETER_FLASH
	ParameterIDRedeye                   = C.MMAL_PARAMETER_REDEYE
	ParameterIDFocus                    = C.MMAL_PARAMETER_FOCUS
	ParameterIDExposureCompInt          = C.MMAL_PARAMETER_EXPOSURE_COMP
	ParameterIDExposureCompRat          = C.MMAL_PARAMETER_EXPOSURE_COMP
	ParameterIDZoom                     = C.MMAL_PARAMETER_ZOOM
	ParameterIDMirror                   = C.MMAL_PARAMETER_MIRROR
	ParameterIDCameraNum                = C.MMAL_PARAMETER_CAMERA_NUM
	ParameterIDCapture                  = C.MMAL_PARAMETER_CAPTURE
	ParameterIDExposureMode             = C.MMAL_PARAMETER_EXPOSURE_MODE
	ParameterIDExpMeteringMode          = C.MMAL_PARAMETER_EXP_METERING_MODE
	ParameterIDFocusStatus              = C.MMAL_PARAMETER_FOCUS_STATUS
	ParameterIDCameraConfig             = C.MMAL_PARAMETER_CAMERA_CONFIG
	ParameterIDCaptureStatus            = C.MMAL_PARAMETER_CAPTURE_STATUS
	ParameterIDFaceTrack                = C.MMAL_PARAMETER_FACE_TRACK
	ParameterIDDrawBoxFacesAndFocus     = C.MMAL_PARAMETER_DRAW_BOX_FACES_AND_FOCUS
	ParameterIDJPEGQFactor              = C.MMAL_PARAMETER_JPEG_Q_FACTOR
	ParameterIDFrameRate                = C.MMAL_PARAMETER_FRAME_RATE
	ParameterIDUseSTC                   = C.MMAL_PARAMETER_USE_STC
	ParameterIDCameraInfo               = C.MMAL_PARAMETER_CAMERA_INFO
	ParameterIDVideoStabilisation       = C.MMAL_PARAMETER_VIDEO_STABILISATION
	ParameterIDEnableRawCapture         = C.MMAL_PARAMETER_ENABLE_RAW_CAPTURE
	ParameterIDEnableDPFFile            = C.MMAL_PARAMETER_ENABLE_DPF_FILE
	ParameterIDDPFFailIsFatal           = C.MMAL_PARAMETER_DPF_FAIL_IS_FATAL
	ParameterIDCaptureMode              = C.MMAL_PARAMETER_CAPTURE_MODE
	ParameterIDInputCrop                = C.MMAL_PARAMETER_INPUT_CROP
	ParameterIDSensorInformation        = C.MMAL_PARAMETER_SENSOR_INFORMATION
	ParameterIDFlashSelect              = C.MMAL_PARAMETER_FLASH_SELECT
	ParameterIDFieldOfView              = C.MMAL_PARAMETER_FIELD_OF_VIEW
	ParameterIDHighDynamicRange         = C.MMAL_PARAMETER_HIGH_DYNAMIC_RANGE
	ParameterIDDynamicRangeCompression  = C.MMAL_PARAMETER_DYNAMIC_RANGE_COMPRESSION
	ParameterIDAlgorithmControl         = C.MMAL_PARAMETER_ALGORITHM_CONTROL
	ParameterIDSharpness                = C.MMAL_PARAMETER_SHARPNESS
	ParameterIDContrast                 = C.MMAL_PARAMETER_CONTRAST
	ParameterIDBrightness               = C.MMAL_PARAMETER_BRIGHTNESS
	ParameterIDSaturation               = C.MMAL_PARAMETER_SATURATION
	ParameterIDISO                      = C.MMAL_PARAMETER_ISO
	ParameterIDAntishake                = C.MMAL_PARAMETER_ANTISHAKE
	ParameterIDImageEffectParameters    = C.MMAL_PARAMETER_IMAGE_EFFECT_PARAMETERS
	ParameterIDCameraBurstCapture       = C.MMAL_PARAMETER_CAMERA_BURST_CAPTURE
	ParameterIDCameraMinISO             = C.MMAL_PARAMETER_CAMERA_MIN_ISO
	ParameterIDCameraUseCase            = C.MMAL_PARAMETER_CAMERA_USE_CASE
	ParameterIDCaptureStatsPass         = C.MMAL_PARAMETER_CAPTURE_STATS_PASS
	ParameterIDCameraCustomSensorConfig = C.MMAL_PARAMETER_CAMERA_CUSTOM_SENSOR_CONFIG
	ParameterIDEnableRegisterFile       = C.MMAL_PARAMETER_ENABLE_REGISTER_FILE
	ParameterIDRegisterFailIsFatal      = C.MMAL_PARAMETER_REGISTER_FAIL_IS_FATAL
	ParameterIDConfigfileRegisters      = C.MMAL_PARAMETER_CONFIGFILE_REGISTERS
	ParameterIDConfigfileChunkRegisters = C.MMAL_PARAMETER_CONFIGFILE_CHUNK_REGISTERS
	ParameterIDJPEGAttachLog            = C.MMAL_PARAMETER_JPEG_ATTACH_LOG
	ParameterIDZeroShutterLag           = C.MMAL_PARAMETER_ZERO_SHUTTER_LAG
	ParameterIDFPSRange                 = C.MMAL_PARAMETER_FPS_RANGE
	ParameterIDCaptureExposureComp      = C.MMAL_PARAMETER_CAPTURE_EXPOSURE_COMP
	ParameterIDSwSharpenDisable         = C.MMAL_PARAMETER_SW_SHARPEN_DISABLE
	ParameterIDFlashRequired            = C.MMAL_PARAMETER_FLASH_REQUIRED
	ParameterIDSwSaturationDisable      = C.MMAL_PARAMETER_SW_SATURATION_DISABLE
	ParameterIDShutterSpeed             = C.MMAL_PARAMETER_SHUTTER_SPEED
	ParameterIDCustomAWBGains           = C.MMAL_PARAMETER_CUSTOM_AWB_GAINS
	ParameterIDCameraSettings           = C.MMAL_PARAMETER_CAMERA_SETTINGS
	ParameterIDPrivacyIndicator         = C.MMAL_PARAMETER_PRIVACY_INDICATOR
	ParameterIDVideoDenoise             = C.MMAL_PARAMETER_VIDEO_DENOISE
	ParameterIDStillsDenoise            = C.MMAL_PARAMETER_STILLS_DENOISE
	ParameterIDAnnotate                 = C.MMAL_PARAMETER_ANNOTATE
	ParameterIDAnnotateV2               = C.MMAL_PARAMETER_ANNOTATE
	ParameterIDAnnotateV3               = C.MMAL_PARAMETER_ANNOTATE
	ParameterIDStereoscopicMode         = C.MMAL_PARAMETER_STEREOSCOPIC_MODE
	ParameterIDCameraInterface          = C.MMAL_PARAMETER_CAMERA_INTERFACE
	ParameterIDCameraClockingMode       = C.MMAL_PARAMETER_CAMERA_CLOCKING_MODE
	ParameterIDCameraRxConfig           = C.MMAL_PARAMETER_CAMERA_RX_CONFIG
	ParameterIDCameraRxTiming           = C.MMAL_PARAMETER_CAMERA_RX_TIMING
	ParameterIDDPFConfig                = C.MMAL_PARAMETER_DPF_CONFIG
	ParameterIDJPEGRestartInterval      = C.MMAL_PARAMETER_JPEG_RESTART_INTERVAL
	ParameterIDCameraIspBlockOverride   = C.MMAL_PARAMETER_CAMERA_ISP_BLOCK_OVERRIDE
	ParameterIDLensShadingOverride      = C.MMAL_PARAMETER_LENS_SHADING_OVERRIDE
	ParameterIDBlackLevel               = C.MMAL_PARAMETER_BLACK_LEVEL
	ParameterIDResizeParams             = C.MMAL_PARAMETER_RESIZE_PARAMS
	ParameterIDCrop                     = C.MMAL_PARAMETER_CROP
	ParameterIDOutputShift              = C.MMAL_PARAMETER_OUTPUT_SHIFT
	ParameterIDCCMShift                 = C.MMAL_PARAMETER_CCM_SHIFT
	ParameterIDCustomCCM                = C.MMAL_PARAMETER_CUSTOM_CCM
	ParameterIDURI                      = C.MMAL_PARAMETER_URI
	ParameterIDSupportedEncodings       = C.MMAL_PARAMETER_SUPPORTED_ENCODINGS
	ParameterIDEXIF                     = C.MMAL_PARAMETER_EXIF
	ParameterIDFaceTrackResults         = C.MMAL_PARAMETER_FACE_TRACK_RESULTS
	ParameterIDFocusRegions             = C.MMAL_PARAMETER_FOCUS_REGIONS
)
