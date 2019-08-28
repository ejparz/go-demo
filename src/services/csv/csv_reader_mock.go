package csv

import (
	"os"
)

type MockCsvReader struct {
	Lines [][]string
	Err error

	ReadFile *os.File
}

func (svc *MockCsvReader) ReadCsv(f *os.File) ([][]string, error){
	svc.ReadFile = f
	return svc.Lines, svc.Err
}