package log

// toss up between logrus / klog v2
// generic imp until decided

import (
	"fmt"
	stdLog "log"
	"os"
)

// ansi color COLOR{F|B} f=foreground b=background
const (
	redF     = "31"
	greenF   = "32"
	yellowF  = "33"
	blueF    = "34"
	magentaF = "35"
	whiteF   = "37" // TODO: light mode terminal?
)

var logger = New()

func New() *stdLog.Logger {
	l := stdLog.New(os.Stdout, "", stdLog.Ldate|stdLog.Ltime|stdLog.LUTC)
	return l
}

func logMsg(message string, color string, lvl string) {
	// TODO: add flag to ignore based on lvl
	logger.Printf("\033[%sm [%s] %s \033[0m", color, lvl, message)
}

func Debug(message string) {
	logMsg(message, blueF, "Debug")
}

func Debugf(message string, v ...any) {
	message = fmt.Sprintf(message, v...)
	logMsg(message, blueF, "Debug")
}

func Info(message string) {
	logMsg(message, whiteF, "Info")
}

func Infof(message string, v ...any) {
	message = fmt.Sprintf(message, v...)
	logMsg(message, whiteF, "Info")
}

func Warning(message string) {
	logMsg(message, yellowF, "Warning")
}

func Warningf(message string, v ...any) {
	message = fmt.Sprintf(message, v...)
	logMsg(message, yellowF, "Warning")
}

func Error(message string) {
	logMsg(message, magentaF, "Error")
}

func Errorf(message string, v ...any) {
	message = fmt.Sprintf(message, v...)
	logMsg(message, magentaF, "Error")
}

func Fatal(message string) {
	logMsg(message, redF, "Fatal")
	os.Exit(69)
}

func Fatalf(message string, v ...any) {
	message = fmt.Sprintf(message, v...)
	logMsg(message, redF, "Fatal")
	os.Exit(69)
}

//func Panic(message string) {
//	logMsg(message, magentaF, "Panic")
// TODO: questionable, see uber go standards guide why not to use panic?
//	panic(message)
//}

/*
func HTTPLogger(param gin.LogFormatterParams) string {
	var color string
	switch {
	case param.StatusCode >= 500:
		color = redF
	case param.StatusCode >= 400:
		color = yellowF
	default:
		color = greenF
	}
	return fmt.Sprintf("[GIN] [%s] \033[%sm %d \033[0m %s %s %d %s %s %s\n",
		param.TimeStamp.Format("2006-01-02 15:04:05"),
		color,
		param.StatusCode,
		param.Method,
		param.Path,
		param.Latency,
		param.ClientIP,
		param.ErrorMessage,
		param.Request.UserAgent(),
	)
}
*/
