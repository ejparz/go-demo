package file

import (
	"go-demo/src/services/processor"
	"testing"
	"go-demo/src/services/test_helper"
	"os"
	"fmt"
)

var p *processor.Processor

func basicSetup(){
	p = processor.NewProcessor()
}

func TestReadFileFromAbsPath(t *testing.T){
	basicSetup()

	wd, err := os.Getwd()
	test_helper.AssertEqual(err, nil)

	println(fmt.Sprintf("Working Dir %v", wd))

	 err = p.PrintStudentAverages(fmt.Sprintf("%v\\..\\..\\..\\data\\grades.csv", wd ))
	test_helper.AssertEqual(err, nil)
}