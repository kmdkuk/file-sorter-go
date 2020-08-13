package log

import (
	"fmt"
	"github.com/spiegel-im-spiegel/logf"
	"os"
	"path/filepath"
	"runtime"
)

var logger *logf.Logger
func init() {
	logger = logf.New(
		logf.WithWriter(os.Stdout),
		logf.WithMinLevel(logf.DEBUG),
		// logf.WithFlags(logf.LstdFlags|logf.Lshortfile),
	)
}
func SetMinLevel(lv logf.Level){
	logger.SetMinLevel(lv)
}

func Info(v ...interface{}){
	logger.Print(format(v))
}

func Debug(v ...interface{}){
	logger.Debug(format(v))
}

func Error(v ...interface{}){
	logger.Error(format(v))
}

func format(v ...interface{}) string {
	_, file, line := caller()
	return fmt.Sprint(file,":", line, v)
}

func caller() (string, string, int){
	pc, file, line, _ := runtime.Caller(3)
	f := runtime.FuncForPC(pc)
	p, _ := os.Getwd()
	path, _ := filepath.Rel(p, file)
	return f.Name(), path, line
}
