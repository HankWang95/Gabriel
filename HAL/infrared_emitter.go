package HAL

import (
	"errors"
	"github.com/HankWang95/log"
	"sync"
	"time"
)

const (
	// 红外线发送模版
	TurnOff                        = 0
	MorningAirConditioningTemplate = 1
)

type infraredEmitter struct {
	lastSignal int
	*sync.RWMutex
	alarmTimer chan int
}

type InfraredTemplate struct{}

func newInfraredEmitter() (inE InfraredEmitter) {
	inEStruct := new(infraredEmitter)
	inEStruct.RWMutex = &sync.RWMutex{}
	inEStruct.lastSignal = -1
	inEStruct.alarmTimer = make(chan int)
	go inEStruct.alarmForInfraredEmitter()
	return inEStruct
}

// 发送红外信号
func (inE *infraredEmitter) SendInfraredSignal(signalTemplateID int) (err error) {
	inE.RWMutex.RLock()
	defer inE.RWMutex.RUnlock()
	_, err = inE.findTemplate(signalTemplateID)
	if err != nil {
		log.Error(err)
		return err
	}
	// todo 硬件交互
	inE.alarmTimer <- signalTemplateID
	log.Info("Send Infrared Signal ID-", signalTemplateID)

	return
}

// 通过id 寻找相对应模版
func (inE *infraredEmitter) findTemplate(signalTemplateID int) (template *InfraredTemplate, err error) {
	switch signalTemplateID {
	case MorningAirConditioningTemplate:
		// todo 找模版
		return nil, nil
	case TurnOff:
		//todo
		return nil, nil
	default:
		return nil, errors.New("signal Template err :no template ")
	}
}

// 告警系统- 在空调开起 8小时 以后 自动关闭
// todo 运行时改回 8hours
func (inE *infraredEmitter) alarmForInfraredEmitter() {
	ticker := time.NewTicker(5 * time.Second)
	for {
		select {
		case signal := <-inE.alarmTimer:
			if signal == TurnOff {
				ticker.Stop()
			} else {
				ticker = time.NewTicker(5 * time.Second)
			}

		case <-ticker.C:
			// 不使用异步方式会导致死锁
			go inE.SendInfraredSignal(TurnOff)
			<-inE.alarmTimer

			log.Info("空调告警系统：8小时自动关闭")
			ticker.Stop()

		}
	}
}
