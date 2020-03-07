package log

import "io"

type Logger interface {
	Print(v ...interface{})
	Println(v ...interface{})
	Printf(format string, v ...interface{})

	Info(v ...interface{})
	Infoln(v ...interface{})
	Infof(format string, v ...interface{})

	Warn(v ...interface{})
	Warnln(v ...interface{})
	Warnf(format string, v ...interface{})

	Error(v ...interface{})
	Errorln(v ...interface{})
	Errorf(format string, v ...interface{})

	Fatal(v ...interface{})
	Fatalln(v ...interface{})
	Fatalf(format string, v ...interface{})

	Debug(v ...interface{})
	Debugln(v ...interface{})
	Debugf(format string, v ...interface{})

	Trace(v ...interface{})
	Traceln(v ...interface{})
	Tracef(format string, v ...interface{})

	SetLevel(v string) error
	SetOutput(output io.Writer)
}

const (
	LoggerNameLogrus = "logrus"
)

var logger = map[string]func() Logger{
	LoggerNameLogrus: newLogrusLogger,
}

func New(name string) Logger {
	f, ok := logger[name]
	if !ok {
		panic("not found logger:" + name)
	}
	return f()
}
