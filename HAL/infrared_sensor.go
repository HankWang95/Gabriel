package HAL

import "github.com/HankWang95/log"

type infraredSensor struct {
	status int
}

func newInfraredSensor() (inS InfraredSensor) {
	return new(infraredSensor)
}

func (inS *infraredSensor) TurnOnTrigger() (err error) {
	log.Info("Turn on trigger -")
	return
}

func (inS *infraredSensor) TurnOffTrigger() (err error) {
	log.Info("Turn off trigger -")
	return
}
