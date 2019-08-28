package processor_test

import (
	"go-demo/src/models"
	"go-demo/src/services/avg"
	"go-demo/src/services/csv"
	"go-demo/src/services/file"
	"go-demo/src/services/parser"
	"go-demo/src/services/printer"
	"go-demo/src/services/processor"
	"go-demo/src/services/test_helper"

	"os"
	"testing"
	"errors"
)

var fileHelper *file.MockFileHelper
var csvReader *csv.MockCsvReader
var scoreParser *parser.MockScoreParserSvc
var avgSvc *avg.MockAvgSvc
var gradePrinter *printer.MockGradePrinter

var f *os.File

var p *processor.Processor

func setUpDependencies() {

	f := &os.File{}

	fileHelper = &file.MockFileHelper{
		File: f,
	}

	csvReader = &csv.MockCsvReader{
		Lines: [][]string{
			{"Bob", "33"},
			{"Joe", "100"},
		},
	}

	scoreParser = &parser.MockScoreParserSvc{
		Scores: []models.Score{
			models.Score{
				Name: "Bob",
				Score: 33,
			},
			models.Score{
				Name: "Joe",
				Score: 100,
			},
		},
	}

	avgSvc = &avg.MockAvgSvc{
		AvgMap: map[string]float64{
			"Bob": 33.00,
			"Joe": 100.00,
		},
	}
	gradePrinter = &printer.MockGradePrinter{}

}

func createSvc() {
	p = processor.NewMockProcessor(
		fileHelper,
		csvReader,
		scoreParser,
		avgSvc,
		gradePrinter,
	)
}

func basicSetup() {
	setUpDependencies()
	createSvc()
}

func TestPrintStudentAverages(t *testing.T) {
	basicSetup()
	err := p.PrintStudentAverages("grades.csv")

	test_helper.AssertEqual(err, nil)


	test_helper.AssertEqual(len(scoreParser.LinesToParse), len(csvReader.Lines))
	test_helper.AssertEqual(len(avgSvc.ScoresFromWhichToCalc), len(scoreParser.Scores))
	test_helper.AssertEqual(gradePrinter.PrintAvgsByNameAscCount, 1)

	test_helper.AssertEqual(fileHelper.CloseFileCount, 1)
}

func TestPrintStudentAverages_CsvParseError(t *testing.T) {
	setUpDependencies()
	errMsg := "Error Parsing grade!"
	csvReader.Err = errors.New(errMsg)

	createSvc()

	err := p.PrintStudentAverages("grades.csv")

	test_helper.AssertEqual(err.Error(), errMsg)
}
