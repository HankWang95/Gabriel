package HAL

type infraredEmitter struct{}

type InfraredTemplate struct{}

func newInfraredEmitter() (inE InfraredEmitter) {
	return new(infraredEmitter)
}

func (inE *infraredEmitter) SendInfraredSignal(signalTemplate InfraredTemplate) (err error) {
	return
}
