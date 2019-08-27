package file

import "os"

type MockFileHelper struct {
	File *os.File
	Err error
	
}

func(svc *MockFileHelper) GetFileByAbsPath(absFilePath string) (*os.File, error){
	return svc.File, svc.Err
}