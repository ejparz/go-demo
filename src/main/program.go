package main

import (
	"go-demo/src/services/processor"
)

func main(){
	p := processor.NewProcessor()
	p.PrintStudentAverages("\\go\\src\\data\\grades.csv")
}

