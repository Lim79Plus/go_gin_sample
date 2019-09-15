package logger

import (
	"log"
	"os"

	"github.com/comail/colog"
)

var logger *log.Logger

// LogInit initialize log configuration
func LogInit() {
	colog.Register()
	cl := colog.NewCoLog(os.Stdout, "gin_web_app ", log.LstdFlags)
	cl.SetMinLevel(colog.LTrace)
	cl.SetDefaultLevel(colog.LInfo)

	logger = cl.NewLogger()
}

// Trace trace
func Trace(v ...interface{}) {
	logger.Println(convertArgs("trace: ", v...)...)
}

func convertArgs(level string, v ...interface{}) []interface{} {
	var l interface{} = level
	variables := []interface{}{}
	variables = append(variables, l)
	variables = append(variables, v...)
	return variables
}
