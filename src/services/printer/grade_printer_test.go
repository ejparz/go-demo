package printer_test

import (
	"fmt"
	"go-demo/src/services/printer"
	"go-demo/src/services/printer/print_helper"
	"go-demo/src/services/test_helper"
	"testing"
)

var printHelper *print_helper.MockPrintHelper
var avgMap map[string]float64

var gradePrinter *printer.GradePrinter

func setUpDependencies() {
	printHelper = &print_helper.MockPrintHelper{}

	avgMap = map[string]float64{
		"Bob":   72.0,
		"Rob":   99.89,
		"Janet": 80.0,
		"Alice": 88.0,
	}
}

func createSvc() {
	gradePrinter = printer.NewMockGradePrinter(
		printHelper,
	)
}

func basicSetup() {
	setUpDependencies()
	createSvc()
}

func TestPrintAvgsRaw(t *testing.T) {
	basicSetup()

	gradePrinter.PrintAvgsRaw(avgMap)

	test_helper.AssertEqual(len(printHelper.PrintedLines), 4)
	assertPrintedLinesContain(printHelper.PrintedLines, "Name: Bob, Avg: 72.00")
	assertPrintedLinesContain(printHelper.PrintedLines, "Name: Janet, Avg: 80.00")
	assertPrintedLinesContain(printHelper.PrintedLines, "Name: Rob, Avg: 99.89")
	assertPrintedLinesContain(printHelper.PrintedLines, "Name: Alice, Avg: 88.00")
}

func TestPrintAvgsByNameAsc(t *testing.T) {
	basicSetup()
	gradePrinter.PrintAvgsByNameAsc(avgMap)

	test_helper.AssertEqual(len(printHelper.PrintedLines), 4)
	test_helper.AssertEqual(printHelper.PrintedLines[0], "Name: Alice, Avg: 88.00")
	test_helper.AssertEqual(printHelper.PrintedLines[1], "Name: Bob, Avg: 72.00")
	test_helper.AssertEqual(printHelper.PrintedLines[2], "Name: Janet, Avg: 80.00")
	test_helper.AssertEqual(printHelper.PrintedLines[3], "Name: Rob, Avg: 99.89")
}


//Maps in Golang are unordered, so in order to assert content ov PrintAvgsRaw, we need a helper method to loop through lines
func assertPrintedLinesContain(printedLines []string, lineToAssert string) {
	for _, line := range printedLines {
		println(fmt.Sprintf("Checking Printed Line: %v", line))
		if line == lineToAssert {
			return
		}
	}

	panic(fmt.Sprintf("Unable to find a printed line with text: %v", lineToAssert))
}
