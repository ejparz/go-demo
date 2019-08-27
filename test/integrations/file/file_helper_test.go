package file

import (
	"go-demo/src/file"
	"testing"
	"go-demo/src/test_helper"
	"os"
	"fmt"
)

var fileHelper *file.FileHelper

func basicSetup(){
	fileHelper = file.NewFileHelper()
}

func TestReadFileFromAbsPath(t *testing.T){
	basicSetup()

	wd, err := os.Getwd()
	test_helper.AssertEqual(err, nil)

	println(fmt.Sprintf("Working Dir %v", wd))

	f, err := fileHelper.GetFileByAbsPath(fmt.Sprintf("%v\\..\\..\\data\\grades.csv", wd ))
	test_helper.AssertEqual(err, nil)

	defer f.Close()
}