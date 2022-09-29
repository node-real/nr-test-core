package log

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

type Logger struct {
	level   int
	logger  *log.Logger
	logFile *os.File
}

func New(out io.Writer, prefix string, flag, level int, file *os.File) *Logger {
	return &Logger{
		level:   level,
		logger:  log.New(out, prefix, flag),
		logFile: file,
	}
}

func (l *Logger) SetDebugLevel(level int) error {
	if level > MaxLevelLog || level < 0 {
		return errors.New("Invalid Debug Level")
	}

	l.level = level
	return nil
}

func (l *Logger) Output(level int, a ...interface{}) error {
	if level >= l.level {
		gid := GetGID()
		gidStr := strconv.FormatUint(gid, 10)

		a = append([]interface{}{LevelName(level), "GID",
			gidStr + ","}, a...)

		return l.logger.Output(CALL_DEPTH, fmt.Sprintln(a...))
	}
	return nil
}

func (l *Logger) Outputf(level int, format string, v ...interface{}) error {
	if level >= l.level {
		gid := GetGID()
		v = append([]interface{}{LevelName(level), "GID",
			gid}, v...)

		return l.logger.Output(CALL_DEPTH, fmt.Sprintf("%s %s %d, "+format+"\n", v...))
	}
	return nil
}

func (l *Logger) Trace(a ...interface{}) {
	l.Output(TraceLog, a...)
}

func (l *Logger) Tracef(format string, a ...interface{}) {
	l.Outputf(TraceLog, format, a...)
}

func (l *Logger) Debug(a ...interface{}) {
	l.Output(DebugLog, a...)
}

func (l *Logger) Debugf(format string, a ...interface{}) {
	l.Outputf(DebugLog, format, a...)
}

func (l *Logger) Info(a ...interface{}) {
	l.Output(InfoLog, a...)
}

func (l *Logger) Infof(format string, a ...interface{}) {
	l.Outputf(InfoLog, format, a...)
}

func (l *Logger) Warn(a ...interface{}) {
	l.Output(WarnLog, a...)
}

func (l *Logger) Warnf(format string, a ...interface{}) {
	l.Outputf(WarnLog, format, a...)
}

func (l *Logger) Error(a ...interface{}) {
	l.Output(ErrorLog, a...)
}

func (l *Logger) Errorf(format string, a ...interface{}) {
	l.Outputf(ErrorLog, format, a...)
}

func (l *Logger) Fatal(a ...interface{}) {
	l.Output(FatalLog, a...)
}

func (l *Logger) Fatalf(format string, a ...interface{}) {
	l.Outputf(FatalLog, format, a...)
}
