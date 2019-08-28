package avg

import (
	"go-demo/src/models"
)

type MockAvgSvc struct {
	AvgMap map[string]float64
	Err error

	ScoresFromWhichToCalc []models.Score
}


func(svc *MockAvgSvc)CalculateAvgs(scores []models.Score) (map[string]float64, error){
	svc.ScoresFromWhichToCalc = scores
	return svc.AvgMap, svc.Err
}