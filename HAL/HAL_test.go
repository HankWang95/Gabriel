package HAL

import (
	"testing"
	"time"
)

var Hobj HardInterface

func init() {
	Hobj = NewHALObject()
}

// --------------------------红外发射器------------------------------

func TestInfraredEmitter_SendInfraredSignal(t *testing.T) {
	var i int
	var err error
	for {
		i++
		if i%5 == 0 {
			err = Hobj.SendInfraredSignal(0)

		} else {
			err = Hobj.SendInfraredSignal(1)

		}
		if err != nil {
			t.Error(err)
		}
		time.Sleep(4 * time.Second)
	}

}

func TestInfraredEmitter_SendInfraredSignalAlarm(t *testing.T) {
	err := Hobj.SendInfraredSignal(1)
	if err != nil {
		t.Error(err)
	}
	time.Sleep(10 * time.Second)
}

// ------------------------红外传感器-----------------------------

func TestInfraredSensor_TurnOnTrigger(t *testing.T) {
	err := Hobj.TurnOnTrigger()
	if err != nil {
		t.Error(err)
	}
}

func TestInfraredSensor_TurnOffTrigger(t *testing.T) {
	err := Hobj.TurnOffTrigger()
	if err != nil {
		t.Error(err)
	}
}

func TestInfraredSensor_TurnOffAndOn(t *testing.T) {
	for i := 0; i < 100; i++ {
		go TestInfraredSensor_TurnOffTrigger(t)

		go TestInfraredSensor_TurnOnTrigger(t)
	}

	time.Sleep(time.Second * 20)

}
