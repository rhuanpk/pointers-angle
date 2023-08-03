package logger

import (
	"log"
	"os"
)

// For all logs, the "Debug" log is the only that has your default output to stderr for in production the user has the possibility to supress it.
var (
	Debug = log.New(os.Stderr, "[DEBUG] ", log.Ltime|log.Lshortfile)
	Info  = log.New(os.Stdout, "[INFO] ", log.Ltime|log.Lshortfile)
	Warn  = log.New(os.Stdout, "[WARN] ", log.Ltime|log.Lshortfile)
	Error = log.New(os.Stdout, "[ERROR] ", log.Ltime|log.Lshortfile)
	Fatal = log.New(os.Stdout, "[FATAL] ", log.Ltime|log.Lshortfile)
)
