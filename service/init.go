package service

import "github.com/HankWang95/Gabriel/HAL"

var (
	HALObj HAL.HardInterface
)

func init() {
	HALObj = HAL.NewHALObject()
}
