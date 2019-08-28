package processor

import (
	"go-demo/src/services/file"
	"go-demo/src/services/csv"
	"go-demo/src/services/parser"
	"go-demo/src/services/avg"
	"go-demo/src/services/printer"
)

type ProcessorInterface interface {

}

type Processor struct {
	fileHelper file.FileHelperInterface
	csvReader csv.CsvReaderInterface
	scoreParser parser.ScoreParserInterface
	avgSvc avg.AvgSvcInterface
	gradePrinter printer.GradePrinterInterface
}

func NewProcessor() *Processor{
	return &Processor{
		fileHelper: file.NewFileHelper(),
		csvReader: csv.NewCsvReader(),
		scoreParser: parser.NewScoreParserSvc(),
		avgSvc: avg.NewAvgSvc(),
		gradePrinter: printer.NewGradePrinter(),
	}
}

func NewMockProcessor(
	fileHelper file.FileHelperInterface,
	csvReader csv.CsvReaderInterface,
	scoreParser parser.ScoreParserInterface,
	avgSvc avg.AvgSvcInterface,
	gradePrinter printer.GradePrinterInterface,
) *Processor{
	return &Processor{
		fileHelper: fileHelper,
		csvReader: csvReader,
		scoreParser: scoreParser,
		avgSvc: avgSvc,
		gradePrinter: gradePrinter,
	}
}

func (svc *Processor) PrintStudentAverages(filePath string) error{
	file, err := svc.fileHelper.GetFileByAbsPath(filePath)
	if(err != nil){
		return err
	}

	defer svc.fileHelper.CloseFile(file)

	lines, err := svc.csvReader.ReadCsv(file)
	if(err != nil){
		return err
	}

	scores, err := svc.scoreParser.ParseLinesToScores(lines)
	if(err != nil){
		return err
	}

	avgMap, err := svc.avgSvc.CalculateAvgs(scores)
	if(err != nil){
		return err
	}

	svc.gradePrinter.PrintAvgsByNameAsc(avgMap)

	return nil
}



