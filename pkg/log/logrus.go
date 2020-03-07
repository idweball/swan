package log

import (
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"time"
)

type logrusLogger struct {
	engine *logrus.Logger
}

func newLogrusLogger() Logger {
	engine := logrus.New()
	engine.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:true,
		DisableColors:true,
		TimestampFormat:time.RFC3339,
	})
	engine.SetOutput(os.Stdout)
	engine.SetLevel(logrus.DebugLevel)
	return &logrusLogger{engine: engine}
}

func (l *logrusLogger) Print(v ...interface{}) {
	l.engine.Print(v...)
}

func (l *logrusLogger) Println(v ...interface{}) {
	l.engine.Println(v...)
}

func (l *logrusLogger) Printf(format string, v ...interface{}) {
	l.engine.Printf(format, v...)
}

func (l *logrusLogger) Info(v ...interface{}) {
	l.engine.Info(v...)
}

func (l *logrusLogger) Infoln(v ...interface{}) {
	l.engine.Infoln(v...)
}

func (l *logrusLogger) Infof(format string, v ...interface{}) {
	l.engine.Infof(format, v...)
}

func (l *logrusLogger) Warn(v ...interface{}) {
	l.engine.Warn(v...)
}

func (l *logrusLogger) Warnln(v ...interface{}) {
	l.engine.Warnln(v...)
}

func (l *logrusLogger) Warnf(format string, v ...interface{}) {
	l.engine.Warnf(format, v...)
}

func (l *logrusLogger) Error(v ...interface{}) {
	l.engine.Error(v...)
}

func (l *logrusLogger) Errorln(v ...interface{}) {
	l.engine.Errorln(v...)
}

func (l *logrusLogger) Errorf(format string, v ...interface{}) {
	l.engine.Errorf(format, v...)
}

func (l *logrusLogger) Fatal(v ...interface{}) {
	l.engine.Fatal(v...)
}

func (l *logrusLogger) Fatalln(v ...interface{}) {
	l.engine.Fatalln(v...)
}

func (l *logrusLogger) Fatalf(format string, v ...interface{}) {
	l.engine.Fatalf(format, v...)
}

func (l *logrusLogger) Debug(v ...interface{}) {
	l.engine.Debug(v...)
}

func (l *logrusLogger) Debugln(v ...interface{}) {
	l.engine.Debugln(v...)
}

func (l *logrusLogger) Debugf(format string, v ...interface{}) {
	l.engine.Debugf(format, v...)
}

func (l *logrusLogger) Trace(v ...interface{}) {
	l.engine.Trace(v...)
}

func (l *logrusLogger) Traceln(v ...interface{}) {
	l.engine.Traceln(v...)
}

func (l *logrusLogger) Tracef(format string, v ...interface{}) {
	l.engine.Tracef(format, v...)
}

func (l *logrusLogger) SetLevel(v string) error {
	lv, err := logrus.ParseLevel(v)
	if err != nil {
		return err
	}
	l.engine.SetLevel(lv)
	return nil
}

func (l *logrusLogger) SetOutput(output io.Writer) {
	l.engine.SetOutput(output)
}
