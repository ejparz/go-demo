package printer

type MockGradePrinter struct {
	PrintAvgsRawCount int
	PrintAvgsByNameAscCount int
}

func (svc *MockGradePrinter) PrintAvgsRaw(avgMap map[string]float64){
	svc.PrintAvgsRawCount++
}

func (svc *MockGradePrinter) PrintAvgsByNameAsc(avgMap map[string]float64){
	svc.PrintAvgsByNameAscCount++
}


