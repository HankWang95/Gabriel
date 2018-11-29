package HAL

import (
	"testing"
	"time"
)

var Hobj HardInterface

func init() {
	Hobj = NewHALObject()
}

func TestInfraredEmitter_SendInfraredSignal(t *testing.T) {
	for {
		err := Hobj.SendInfraredSignal(1)
		if err != nil {
			t.Error(err)
		}
		time.Sleep(1 * time.Second)
	}

}

func TestInfraredEmitter_SendInfraredSignalAlarm(t *testing.T) {
	err := Hobj.SendInfraredSignal(1)
	if err != nil {
		t.Error(err)
	}
	time.Sleep(10 * time.Second)
}
