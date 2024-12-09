package utils

import (
	"github.com/qmaru/minitools/v2/file"
)

// FileSuite 初始化
var FileSuite *file.FileSuiteBasic

func init() {
	FileSuite = file.New()
}
