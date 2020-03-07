package log

import "io"

var defaultLog = New(LoggerNameLogrus)

func Print(v ...interface{}) {
	defaultLog.Print(v...)
}

func Println(v ...interface{}) {
	defaultLog.Println(v...)
}

func Printf(format string, v ...interface{}) {
	defaultLog.Printf(format, v...)
}

func Info(v ...interface{}) {
	defaultLog.Info(v...)
}

func Infoln(v ...interface{}) {
	defaultLog.Infoln(v...)
}

func Infof(format string, v ...interface{}) {
	defaultLog.Infof(format, v...)
}

func Error(v ...interface{}) {
	defaultLog.Error(v...)
}

func Errorln(v ...interface{}) {
	defaultLog.Errorln(v...)
}

func Errorf(format string, v ...interface{}) {
	defaultLog.Errorf(format, v...)
}

func Warn(v ...interface{}) {
	defaultLog.Warn(v...)
}

func Warnln(v ...interface{}) {
	defaultLog.Warnln(v...)
}

func Warnf(format string, v ...interface{}) {
	defaultLog.Warnf(format, v...)
}

func Fatal(v ...interface{}) {
	defaultLog.Fatal(v...)
}

func Fatalln(v ...interface{}) {
	defaultLog.Fatalln(v...)
}

func Fatalf(format string, v ...interface{}) {
	defaultLog.Fatalf(format, v...)
}

func Debug(v ...interface{}) {
	defaultLog.Debug(v...)
}

func Debugln(v ...interface{}) {
	defaultLog.Debugln(v...)
}

func Debugf(format string, v ...interface{}) {
	defaultLog.Debugf(format, v...)
}

func Trace(v ...interface{}) {
	defaultLog.Trace(v...)
}

func Traceln(v ...interface{}) {
	defaultLog.Traceln(v...)
}

func Tracef(format string, v ...interface{}) {
	defaultLog.Tracef(format, v...)
}

func SetLeveL(lv string) {
	defaultLog.SetLevel(lv)
}

func SetOuput(output io.Writer) {
	defaultLog.SetOutput(output)
}
