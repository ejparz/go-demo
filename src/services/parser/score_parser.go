package parser

import (
	"go-demo/src/models"
	"strconv"
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
	
	scores := []models.Score{}
	for i, line := range lines{
		//Skip header line
		if(i == 0){
			continue;
		}
		name := line[0]
		val, err := svc.ParseInt64FromString(line[1])

		if(err != nil){
			return scores, err
		}

		score := models.Score{
			Name: name,
			Score: val,
		}

		scores = append(scores, score)


	}
	return scores, nil
}


func (svc *ScoreParserSvc) ParseInt64FromString(num string) (int64, error){
	return strconv.ParseInt(num, 10, 64)
}