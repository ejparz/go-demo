package print_helper

//Helper interface to help with mocking/testing print functionality
type PrintHelperInterface interface {
	PrintLn(line string)
}

type PrintHelper struct {

}

func NewPrintHelper() *PrintHelper {
	return &PrintHelper{}
}

func (svc *PrintHelper)PrintLn(line string){
	println(line)
}

