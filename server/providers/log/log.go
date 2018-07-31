package log

import (
	"io"
	"log"
)

var (
	//Trace logs messages at trace level
	Trace *log.Logger
	//Info logs messages at info level
	Info *log.Logger
	//Warning logs messages at warn level
	Warning *log.Logger
	//Error logs messages at error level
	Error *log.Logger
)

//InitLog initializes the logger
func InitLog(
	traceHandle io.Writer,
	infoHandle io.Writer,
	warningHandle io.Writer,
	errorHandle io.Writer) {

	Trace = log.New(traceHandle,
		"TRACE: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Info = log.New(infoHandle,
		"INFO: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Warning = log.New(warningHandle,
		"WARNING: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Error = log.New(errorHandle,
		"ERROR: ",
		log.Ldate|log.Ltime|log.Lshortfile)
}
