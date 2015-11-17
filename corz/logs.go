// Shows how basic Log4J like library could be implemented.
//
// Author: Dmitri Krasnenko

package corz

import (
	"os"
	"flag"
	gol "log"
)

type (
	severity string
	blackhole struct{
	}
)

type memento struct {
	info *gol.Logger
	warn *gol.Logger
	erno *gol.Logger
}

var (
	std * memento
	nul = blackhole{
	}
)

const (
	warn = "WARN"
	info = "INFO"
	erro = "ERROR"
)

func (blackhole) Write(p []byte) (int, error) {
	return len(p), nil
}

func init() {
	//Set log severity using flag.Value API
	flag.Var(new(severity),"stderrthreshold","Log level")

	//Set logger to ERROR level  by default
	std = &memento{
		info:gol.New(nul, "", gol.LstdFlags),
		warn:gol.New(nul, "", gol.LstdFlags),
		erno:gol.New(os.Stdout,"", gol.LstdFlags),
	}
}

func (s *severity) String() string {
	return string(*s)
}

func (s *severity) Set(level string) error {
	//Init logger
	switch level {
	case info:
		std = &memento{
			info:gol.New(os.Stdout, "", gol.LstdFlags),
			warn:gol.New(os.Stdout, "", gol.LstdFlags),
			erno:gol.New(os.Stdout, "", gol.LstdFlags),
		}
	case warn:
		std = &memento{
			info:gol.New(nul, "", gol.LstdFlags),
			warn:gol.New(os.Stdout, "", gol.LstdFlags),
			erno:gol.New(os.Stdout, "", gol.LstdFlags),
		}
	default:
		std = &memento{
			info:gol.New(nul, "", gol.LstdFlags),
			warn:gol.New(nul, "", gol.LstdFlags),
			erno:gol.New(os.Stdout,"", gol.LstdFlags),
		}
	}

	return nil
}

func logger(level string) (*gol.Logger) {
	switch level {
	case info:
		return std.info
	case warn:
		return std.warn
	}

	return std.erno
}

func Infoln(v ...interface{}) {
	logger(info).Println(
		append([]interface{}{"INFO: "},v...)...,
	)
}

func Warnln(v ...interface{}) {
	logger(warn).Println(
		v...,
	)
}

func Errnln(v ...interface{}) {
	logger(erro).Println(
		v...,
	)
}

func Infof(format string, v ...interface{}) {
	logger(info).Printf(
		format,
		v...,
	)
}

func Warnf(format string, v ...interface{}) {
	logger(warn).Printf(
		format,
		v...,
	)
}

func Errorf(format string, v ...interface{}) {
	logger(erro).Printf(
		format,
		v...,
	)
}

func Info(v ...interface{}) {
	logger(info).Print(
		v...,
	)
}

func Warn(v ...interface{}) {
	logger(warn).Print(
		v...,
	)
}

func Error(v ...interface{}) {
	logger(erro).Print(
		v...,
	)
}

