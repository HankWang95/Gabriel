package HAL

type camera struct {
}

func newCamera() (c Camera) {
	return new(camera)
}

func (c *camera) TakePhoto() (path string, err error) {
	return
}

func (c *camera) RecordVideo(sec int) (path string, err error) {
	return
}
