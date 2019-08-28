package csv

import (
	"os"
)

type MockCsvReader struct {
	Lines [][]string
	Err error
}

func (svc *MockCsvReader) ReadCsv(f *os.File) ([][]string, error){
	return svc.Lines, svc.Err
}