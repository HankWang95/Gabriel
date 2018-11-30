package HAL

import (
	"github.com/HankWang95/log"
	"sync"
)

type infraredSensor struct {
	*sync.Mutex
	on bool
}

func newInfraredSensor() (inS InfraredSensor) {
	inSStruct := new(infraredSensor)
	inSStruct.on = false
	inSStruct.Mutex = &sync.Mutex{}
	return inSStruct
}

func (inS *infraredSensor) TurnOnTrigger() (err error) {
	inS.Mutex.Lock()
	defer inS.Mutex.Unlock()
	if inS.on {
		return nil
	}
	inS.on = true
	// todo
	log.Info("Turn on trigger -")
	return
}

func (inS *infraredSensor) TurnOffTrigger() (err error) {
	inS.Mutex.Lock()
	defer inS.Mutex.Unlock()
	if !inS.on {
		return nil
	}
	inS.on = false
	log.Info("Turn off trigger -")
	return
}
