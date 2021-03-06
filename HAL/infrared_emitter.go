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
	signal chan int
}

type InfraredTemplate struct{}

func newInfraredEmitter() (inE InfraredEmitter) {
	inEStruct := new(infraredEmitter)
	inEStruct.RWMutex = &sync.RWMutex{}
	inEStruct.lastSignal = -1
	inEStruct.signal = make(chan int)
	go inEStruct.listenSignal()
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
	inE.signal <- signalTemplateID
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

// -------------------------- 空调开启超时告警系统 ----------------------------

var stopToken = make(chan struct{})
var alarmIsRunning = struct {
	*sync.Mutex
	running bool
}{
	Mutex:   &sync.Mutex{},
	running: false,
}
var ticker = make(chan *time.Ticker)

// 告警系统- 在空调开起 8小时 以后 自动关闭
func (inE *infraredEmitter) listenSignal() {
	for {
		select {
		case signal := <-inE.signal:
			if signal == TurnOff {
				if alarmIsRunning.running {
					stopToken <- struct{}{}
					log.Info("send stop token")
				}
			} else {
				go inE.alarmRun()
			}
		}
	}
}

// 警报器
func (inE *infraredEmitter) alarmRun() {
	if alarmIsRunning.running {
		ticker <- time.NewTicker(3 * time.Second)
		return
	}

	alarmIsRunning.Lock()
	alarmIsRunning.running = true
	alarmIsRunning.Unlock()

	var t *time.Ticker
	t = time.NewTicker(3 * time.Second)

Loop:
	for {
		select {
		case t = <-ticker:
			break
		case <-t.C:
			log.Info("告警：空调运行时间过长，自动关闭")
			alarmIsRunning.Lock()
			alarmIsRunning.running = false
			alarmIsRunning.Unlock()
			err := inE.SendInfraredSignal(TurnOff)
			if err != nil {
				log.Error(err)
			}
			break Loop
		case <-stopToken:
			alarmIsRunning.Lock()
			alarmIsRunning.running = false
			alarmIsRunning.Unlock()
			break Loop
		}
	}

}
