package main

import (
	"fmt"
	"go-demo/src/services/processor"
)

func main(){
	p := processor.NewProcessor()
	p.PrintStudentAverages()
}

