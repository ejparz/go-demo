package avg

import (
	"go-demo/src/models"
)

type MockAvgSvc struct {
	AvgMap map[string]float64
	Err error
}


func(svc *MockAvgSvc)CalculateAvgs(scores []models.Score) (map[string]float64, error){
	return svc.AvgMap, svc.Err
}