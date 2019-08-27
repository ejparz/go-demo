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

	//Reuse arg in fmt sprintf 
	//https://stackoverflow.com/questions/26624699/is-there-a-way-to-reuse-an-argument-in-fmt-printf
	ps := string(os.PathSeparator)
	f, err := fileHelper.GetFileByAbsPath(fmt.Sprintf("%v%[2]v..%[2]v..%[2]vtest_data%[2]vgrades.csv", wd, ps ))
	test_helper.AssertEqual(err, nil)

	defer f.Close()
}