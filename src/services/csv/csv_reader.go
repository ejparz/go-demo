package csv

import (
	"os"
	"encoding/csv"
)

type CsvReaderInterface interface {
	ReadCsv(f *os.File) ([][]string, error)
}

type CsvReader struct {

}

func NewCsvReader() *CsvReader{
	return &CsvReader{}
}


func (svc *CsvReader) ReadCsv(f *os.File) ([][]string, error){
	return csv.NewReader(f).ReadAll()
}