package avg

import (
	"go-demo/src/models"
)

type AvgSvcInterface interface {
	CalculateAvgs(scores []models.Score) (map[string]float64, error)
}

type AvgSvc struct {

}

func NewAvgSvc() *AvgSvc {
	return &AvgSvc{}
}

func (svc *AvgSvc) CalculateAvgs(scores []models.Score) (map[string]float64, error) {
	return map[string]float64{}, nil
}