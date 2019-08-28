package parser

import (
	"go-demo/src/models"
)

type ScoreParserInterface interface {

}

type ScoreParserSvc struct {

}

func NewScoreParserSvc() *ScoreParserSvc{
	return &ScoreParserSvc{}
}

func(svc *ScoreParserSvc) ParseLinesToScores(lines [][]string) ([]models.Score) {
	return []models.Score{}
}