package service

// 控制空调
func CtrlAirConditioning(signalTemplateID int) (err error) {
	err = HALObj.SendInfraredSignal(signalTemplateID)
	if err != nil {
		return err
	}
	return nil
}

// 打开红外传感器
func TurnOnTrigger() (err error) {
	err = HALObj.TurnOnTrigger()
	if err != nil {
		return err
	}
	return nil
}

// 关闭红外传感器
func TurnOffTrigger() (err error) {
	err = HALObj.TurnOffTrigger()
	if err != nil {
		return err
	}
	return nil
}

// 拍照
func TakePhoto() (path string, err error) {
	path, err = HALObj.TakePhoto()
	if err != nil {
		return path, err
	}
	return
}

// 打开视频直播
func StartPushStream() (path string, err error) {
	path, err = HALObj.StartPushVideoStream()
	if err != nil {
		return path, err
	}
	return
}

// 关闭视频直播
func StopPushStream() (path string, err error) {
	path, err = HALObj.StopPushVideoStream()
	if err != nil {
		return path, err
	}
	return
}
