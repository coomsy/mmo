package logger

// toss up between logrus / klog v2
// generic imp until decided

import (
	stdLog "log"
	"os"
)

var log *stdLog.Logger

// ansi color COLOR{F|B} f=foreground b=background
const (
	redF     = "31"
	greenF   = "32"
	yellowF  = "33"
	blueF    = "34"
	magentaF = "35"
	whiteF   = "37" // TODO: light mode terminal?
)

func init() {
	log = stdLog.New(os.Stdout, "", stdLog.Ldate|stdLog.Ltime|stdLog.LUTC)
}

func logMsg(message string, color string, lvl string) {
	// TODO: add flag to ignore based on lvl
	log.Printf("\033[%sm [%s] %s \033[0m", color, lvl, message)
}

func Debug(message string) {
	logMsg(message, blueF, "Debug")
}

func Info(message string) {
	logMsg(message, whiteF, "Info")
}

func Warning(message string) {
	logMsg(message, yellowF, "Warning")
}

func Error(message string) {
	logMsg(message, magentaF, "Error")
}

func Panic(message string) {
	logMsg(message, magentaF, "Panic")
	panic(message)
}

func Fatal(message string) {
	logMsg(message, redF, "Fatal")
	os.Exit(69)
}
