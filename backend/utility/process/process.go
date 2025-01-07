package process

import (
	"os"
)

func GetProcessId() int {
	return os.Getpid()
}
