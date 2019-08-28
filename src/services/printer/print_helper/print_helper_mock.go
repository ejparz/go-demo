package print_helper

type MockPrintHelper struct {
	PrintedLines []string
}

func (svc *MockPrintHelper) PrintLn(line string){
	svc.PrintedLines = append(svc.PrintedLines, line)
}