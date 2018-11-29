package HAL

import (
	"errors"
	"github.com/HankWang95/log"
)

const (
	// 红外线发送模版
	MorningAirConditioningTemplate = 1
)

type infraredEmitter struct{}

type InfraredTemplate struct{}

func newInfraredEmitter() (inE InfraredEmitter) {
	return new(infraredEmitter)
}

// 发送红外信号
func (inE *infraredEmitter) SendInfraredSignal(signalTemplateID int) (err error) {
	_, err = inE.findTemplate(signalTemplateID)
	if err != nil {
		return err
	}

	log.Info("Send Infrared Signal ID-", signalTemplateID)

	return
}

// 通过id 寻找相对应模版
func (inE *infraredEmitter) findTemplate(signalTemplateID int) (template *InfraredTemplate, err error) {
	switch signalTemplateID {
	case MorningAirConditioningTemplate:
		// todo 找模版
		return nil, nil
	default:
		return nil, errors.New("signal Template err :no template ")
	}
}