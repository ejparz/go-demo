package parser_test

import (
	"go-demo/src/services/test_helper"
	"go-demo/src/services/parser"
	"testing"
)


var scoreParser *parser.ScoreParserSvc

func basicSetup(){
	scoreParser = parser.NewScoreParserSvc()
}

func TestParseLinesToScores(t *testing.T){
	basicSetup()
	lines := [][]string{
		{"Name", "Score"},
		{"Bobby", "100"},
		{"Joann", "99"},
		{"Tom", "58"},
		{"Sally", "33"},
	}

	scores, err := scoreParser.ParseLinesToScores(lines)

	test_helper.AssertEqual(err, nil)

	test_helper.AssertEqual(len(scores), 4)
	test_helper.AssertEqual(scores[0].Name, "Bobby")
	test_helper.AssertEqual(scores[0].Score, int64(100))
}

//TODO Test Unhappy Path Parsing Lines to Scores

func TestParserInt64FromString_Success(t *testing.T){
	basicSetup()
	val, err := scoreParser.ParseInt64FromString("200")

	test_helper.AssertEqual(err, nil)
	test_helper.AssertEqual(val, int64(200))
}

func TestParserInt64FromString_Failure(t *testing.T){
	basicSetup()
	_, err := scoreParser.ParseInt64FromString("garbage")

	test_helper.AssertContains(err.Error(), "invalid syntax")
}