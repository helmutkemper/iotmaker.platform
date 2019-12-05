package genericTypes

type ImageDataCaptureMethod uint8

const (
	KImageDataCaptureMethodBooleanSensibility ImageDataCaptureMethod = iota
	KImageDataCaptureMethodCompleteData
	KImageDataCaptureMethodAlphaChannelOnly
)
