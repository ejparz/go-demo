package file

import "os"

type MockFileHelper struct {
	File *os.File
	Err error

	CloseFileCount int
	
}

func(svc *MockFileHelper) GetFileByAbsPath(absFilePath string) (*os.File, error){
	return svc.File, svc.Err
}

func (svc *MockFileHelper) CloseFile(f *os.File){
	svc.CloseFileCount ++
}