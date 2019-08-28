package file

import (
	"fmt"
	"go-demo/src/services/processor"
	"go-demo/src/services/test_helper"
	"os"
	"testing"
)

var p *processor.Processor

func basicSetup() {
	p = processor.NewProcessor()
}

func TestPrintStudentAverages(t *testing.T) {
	basicSetup()

	wd, err := os.Getwd()
	test_helper.AssertEqual(err, nil)

	println(fmt.Sprintf("Working Dir %v", wd))

	err = p.PrintStudentAverages(fmt.Sprintf("%v\\..\\..\\..\\data\\grades.csv", wd))
	test_helper.AssertEqual(err, nil)
}
