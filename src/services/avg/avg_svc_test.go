package avg_test

import (
	"go-demo/src/services/test_helper"
	"go-demo/src/models"
	"go-demo/src/services/avg"
	"testing"
)

var avgSvc *avg.AvgSvc

func basicSetup(){
	avgSvc = avg.NewAvgSvc()
}

func TestCalculateAvgs(t *testing.T){
	basicSetup()
	scores := []models.Score{
		models.Score{
			Name: "Bob",
			Score: 100,
		},
		models.Score{
			Name: "Bob",
			Score: 95,
		},
		models.Score{
			Name: "Bob",
			Score: 90,
		},
		models.Score{
			Name: "Sally",
			Score: 88,
		},
		models.Score{
			Name: "Sally",
			Score: 87,
		},
	}

	avgMap, err := avgSvc.CalculateAvgs(scores)

	test_helper.AssertEqual(err, nil)
	test_helper.AssertEqual(avgMap["Bob"], 95.0)
	test_helper.AssertEqual(avgMap["Sally"], 87.5)
}
