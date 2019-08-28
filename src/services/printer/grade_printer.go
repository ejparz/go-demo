package printer

import (
	"go-demo/src/services/printer/print_helper"
	"fmt"
	"sort"
)

type GradePrinterInterface interface {
	PrintAvgsRaw(avgMap map[string]float64)
	PrintAvgsByNameAsc(avgMap map[string]float64)
}


type GradePrinter struct {
	printHelper print_helper.PrintHelperInterface
}

func NewGradePrinter() *GradePrinter{
	return &GradePrinter{
		printHelper: print_helper.NewPrintHelper(),
	}
}

func NewMockGradePrinter(
	printHelper print_helper.PrintHelperInterface,
) *GradePrinter{
	return &GradePrinter{
		printHelper: printHelper,
	}
}

func (svc *GradePrinter) PrintAvgsRaw(avgMap map[string]float64){
	for name, avg := range avgMap{
		svc.printAvg(name, avg)
	}
}

func (svc *GradePrinter) PrintAvgsByNameAsc(avgMap map[string]float64){
	names := make([]string, 0, len(avgMap))
	for name := range avgMap {
		names = append(names, name)
	}
	sort.Strings(names)
	
	for _, name := range names {
		svc.printAvg(name, avgMap[name])
	}

}


func(svc *GradePrinter) printAvg(name string, avg float64){
	svc.printHelper.PrintLn(fmt.Sprintf("Name: %v, Avg: %0.2f", name, avg))
}
