package parser

import (
	"go-demo/src/models"
)

type ScoreParserInterface interface {
	ParseLinesToScores(lines [][]string) ([]models.Score, error)
}

type ScoreParserSvc struct {

}

func NewScoreParserSvc() *ScoreParserSvc{
	return &ScoreParserSvc{}
}

func(svc *ScoreParserSvc) ParseLinesToScores(lines [][]string) ([]models.Score, error) {
	return []models.Score{}, nil
}