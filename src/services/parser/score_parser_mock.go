package parser

import (
	"go-demo/src/models"
)

type MockScoreParserSvc struct {
	Scores []models.Score
	Err error
}

func (svc *MockScoreParserSvc) ParseLinesToScores(lines [][]string) ([]models.Score, error) {
	return svc.Scores, svc.Err
}