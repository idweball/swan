package log

import (
	"io"

	"gopkg.in/natefinch/lumberjack.v2"
)

type FileOutputOption struct {
	Filename  string
	MaxSize   int
	MaxAge    int
	MaxBackup int
	Compress  bool
}

func NewFileOutput(opt FileOutputOption) io.Writer {
	return &lumberjack.Logger{
		Filename:   opt.Filename,
		MaxSize:    opt.MaxSize,
		MaxAge:     opt.MaxAge,
		MaxBackups: opt.MaxBackup,
		LocalTime:  false,
		Compress:   opt.Compress,
	}
}
