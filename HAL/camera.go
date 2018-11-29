package HAL

import "github.com/HankWang95/log"

type camera struct {
}

func newCamera() (c Camera) {
	return new(camera)
}

func (c *camera) TakePhoto() (path string, err error) {
	log.Info("Take photo - ", path)
	return
}

func (c *camera) StartPushVideoStream() (path string, err error) {
	log.Info("Start push video stream -", path)
	return
}

func (c *camera) StopPushVideoStream() (path string, err error) {
	log.Info("Stop push video stream -", path)
	return
}
