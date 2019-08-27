package file

import (
	"os"
)


type FileHelperInterface interface {
	GetFileByAbsPath(absFilePath string) (*os.File, error)
}


type FileHelper struct {

}

func NewFileHelper() *FileHelper{
	return &FileHelper{}
}

//Wrapper Method for os.Open so that we can mock/test this particular functionality.
func (svc *FileHelper) GetFileByAbsPath(absFilePath string) (*os.File, error){
	return os.Open(absFilePath)
}