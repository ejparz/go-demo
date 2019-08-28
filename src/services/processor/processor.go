package processor

import ()

type ProcessorInterface interface {

}

type Processor struct {


}

func NewProcessor() *Processor{
	return &Processor{}
}

func (svc *Processor) PrintStudentAverages() {
	//file svc fetch file
	//Parse CSV file into list of student 
	//Calculator - Fetch Map of student to avg.

	//GradePrinter - Print the averages in a pretty way.

}



