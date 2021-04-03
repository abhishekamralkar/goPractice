package golang

import (
	"runtime"
)

func Version() string {
	return runtime.Version()

}
