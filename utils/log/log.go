package log

import "log"

func Errorf(format string, v ...interface{}) {
	log.Printf(format, v)
}

func Infof(format string, v ...interface{}) {
	log.Printf(format, v)
}