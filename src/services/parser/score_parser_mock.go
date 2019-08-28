package parser

import (
	"go-demo/src/models"
)

type MockScoreParserSvc struct {
	Scores []models.Score
	Err error

	LinesToParse [][]string
}

func (svc *MockScoreParserSvc) ParseLinesToScores(lines [][]string) ([]models.Score, error) {
	svc.LinesToParse = lines
	return svc.Scores, svc.Err
}