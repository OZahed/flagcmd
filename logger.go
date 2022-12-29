package flagcmd

import (
	"log"
	"os"
)

var (
	lg *log.Logger
)

func init() {
	lg = log.New(os.Stderr, "Error: ", log.Ltime|log.LUTC)
}

func LogErrorf(s string, i ...interface{}) {
	lg.Printf("%s\n", i)
}
