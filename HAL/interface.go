package HAL

// 摄像头
type Camera interface {
	TakePhoto() (path string, err error)
	StartPushVideoStream() (path string, err error)
	StopPushVideoStream() (path string, err error)
}

// 红外传感器
type InfraredSensor interface {
	TurnOnTrigger() (err error)
	TurnOffTrigger() (err error)
}

// 红外线发射器
type InfraredEmitter interface {
	SendInfraredSignal(signalTemplateID int) (err error)
}

// todo 红外线接收

type HardInterface interface {
	Camera
	InfraredSensor
	InfraredEmitter
}

type hardObject struct {
	Camera
	InfraredSensor
	InfraredEmitter
}

// 获取硬件层抽象对象
func NewHALObject() (hardInterface HardInterface) {
	h := new(hardObject)
	h.Camera = newCamera()
	h.InfraredSensor = newInfraredSensor()
	h.InfraredEmitter = newInfraredEmitter()
	return h
}
