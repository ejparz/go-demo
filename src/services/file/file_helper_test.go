package file_test

import (
	"go-demo/src/services/file"
	"go-demo/src/services/test_helper"
	"os"
	"testing"
	"fmt"
)

var fileHelper *file.FileHelper

func basicSetup() {
	fileHelper = file.NewFileHelper()
}

func TestReplaceSlashesWithOsSpecificPathSeparator_FwdSlash(t *testing.T) {
	path := fileHelper.ReplaceSlashesWithOsSpecificPathSeparator("\\go\\src\\test")
	test_helper.AssertEqual(path, fmt.Sprintf("%[1]vgo%[1]vsrc%[1]vtest", string(os.PathSeparator)))
}

func TestReplaceSlashesWithOsSpecificPathSeparator_BackSlash(t *testing.T) {
	path := fileHelper.ReplaceSlashesWithOsSpecificPathSeparator("/go/src/test")
	test_helper.AssertEqual(path, fmt.Sprintf("%[1]vgo%[1]vsrc%[1]vtest", string(os.PathSeparator)))
}
