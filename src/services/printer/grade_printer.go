package printer

type GradePrinterInterface interface {
	PrintAvgsRaw(avgMap map[string]float64)
	PrintAvgsByNameAsc(avgMap map[string]float64)
	PrintAvgsByScoreDesc(avgMap map[string]float64)
}


type GradePrinter struct {

}

func NewGradePrinter() *GradePrinter{
	return &GradePrinter{}
}


func (svc *GradePrinter) PrintAvgsRaw(avgMap map[string]float64){

}

func (svc *GradePrinter) PrintAvgsByNameAsc(avgMap map[string]float64){

}

func (svc *GradePrinter) PrintAvgsByScoreDesc(avgMap map[string]float64){

}
