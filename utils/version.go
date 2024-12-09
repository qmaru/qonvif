package utils

import (
	"fmt"
)

const (
	DateVer   string = "COMMIT_DATE"
	CommitVer string = "COMMIT_VERSION"
	GoVer     string = "COMMIT_GOVER"
)

var Version string = fmt.Sprintf("%s (git-%s) (%s)", DateVer, CommitVer, GoVer)
