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


type CountAndTotal struct {
	Count int64
	Total int64
}

func (svc *AvgSvc) CalculateAvgs(scores []models.Score) (map[string]float64, error) {

	countAndTotalMap := map[string]CountAndTotal{}

	for _, s := range scores {
		if val, ok := countAndTotalMap[s.Name]; ok {
			countAndTotalMap[s.Name] = CountAndTotal{
				Count: val.Count + 1,
				Total: val.Total + s.Score,
			}
		} else{
			countAndTotalMap[s.Name] = CountAndTotal{
				Count: 1,
				Total: s.Score,
			}
		}
	}

	return svc.CalculateAvgsFromCountAndTotalMap(countAndTotalMap)
}

func (svc *AvgSvc) CalculateAvgsFromCountAndTotalMap(countAndTotals map[string]CountAndTotal) (map[string]float64, error){

	avgMap := map[string]float64{}
	for name, ct := range countAndTotals{
		avgMap[name] = float64(ct.Total) / float64(ct.Count)
	}

	return avgMap, nil
}