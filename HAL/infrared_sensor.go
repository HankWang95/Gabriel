package HAL

type infraredSensor struct{}

func newInfraredSensor() (inS InfraredSensor) {
	return new(infraredSensor)
}

func (inS *infraredSensor) TurnOnTrigger() (err error) {
	return
}

func (inS *infraredSensor) TurnOffTrigger() (err error) {
	return
}
