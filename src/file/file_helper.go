package file

import (
	"os"
	"strings"
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
	cleanAbsFilePath := svc.ReplaceSlashesWithOsSpecificPathSeparator(absFilePath)
	return os.Open(cleanAbsFilePath)
}

func (svc *FileHelper) ReplaceSlashesWithOsSpecificPathSeparator(path string) string {
	fixedPath1 := strings.Replace(path, "\\", string(os.PathSeparator), -1)
	fixedPath2 := strings.Replace(fixedPath1, "/", string(os.PathSeparator), -1)

	return fixedPath2
}