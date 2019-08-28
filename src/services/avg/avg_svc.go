package avg

import (
	"go-demo/src/models"
)

type AvgSvcInterface interface {

}

type AvgSvc struct {

}

func NewAvgSvc() *AvgSvc {
	return &AvgSvc{}
}

func (svc *AvgSvc) CalculateAvgs(scores []models.Score) map[string]float64 {
	return map[string]float64{}
}