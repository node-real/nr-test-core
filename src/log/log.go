package log

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	path2 "path"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"
)

const (
	Blue   = "0;34"
	Red    = "0;31"
	Green  = "0;32"
	Yellow = "0;33"
	Cyan   = "0;36"
	Pink   = "1;35"
)

func Color(code, msg string) string {
	return fmt.Sprintf("\033[%sm%s\033[m", code, msg)
}

const (
	TraceLog = iota
	DebugLog
	InfoLog
	WarnLog
	ErrorLog
	FatalLog
	MaxLevelLog
)

var (
	levels = map[int]string{
		DebugLog: Color(Green, "[DEBUG]"),
		InfoLog:  Color(Cyan, "[INFO]"),
		WarnLog:  Color(Yellow, "[WARN]"),
		ErrorLog: Color(Red, "[ERROR]"),
		FatalLog: Color(Red, "[FATAL]"),
		TraceLog: Color(Pink, "[TRACE]"),
	}
	Stdout = os.Stdout
)

const (
	NAME_PREFIX          = "LEVEL"
	CALL_DEPTH           = 2
	DEFAULT_MAX_LOG_SIZE = 20
	BYTE_TO_MB           = 1024 * 1024
	PATH                 = "./Log/"
)

func GetGID() uint64 {
	var buf [64]byte
	b := buf[:runtime.Stack(buf[:], false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.ParseUint(string(b), 10, 64)
	return n
}

var Log *Logger

func init() {
	logFile := newOutputLogFile()
	InitLog(InfoLog, logFile)
}

func LevelName(level int) string {
	if name, ok := levels[level]; ok {
		return name
	}
	return NAME_PREFIX + strconv.Itoa(level)
}

func newOutputLogFile() *os.File {
	path, _ := os.Getwd()
	for i := 0; i < 10; i++ {
		fileEnum, _ := os.ReadDir(path)
		hasMod := false
		for _, f := range fileEnum {
			if f.Name() == "go.mod" {
				hasMod = true
				break
			}
		}
		if hasMod {
			break
		} else {
			path = substr(path, 0, strings.LastIndex(path, "/"))
		}
	}
	outputPath := path2.Join(path, "output")
	_, err := os.Stat(outputPath)
	if err != nil {
		os.Mkdir(outputPath, 0755)
	}

	logPath := path2.Join(outputPath, getLogName())
	logFile, err := os.Create(logPath)
	if err != nil {
		fmt.Println(err)
	}
	return logFile
}

func getLogName() string {
	timeStr := time.Now().Format("2006-01-02_15:04:05")
	return fmt.Sprintf("%s.log", timeStr)
}

func substr(s string, pos, length int) string {
	runes := []rune(s)
	l := pos + length
	if l > len(runes) {
		l = len(runes)
	}
	return string(runes[pos:l])
}

func NameLevel(name string) int {
	for k, v := range levels {
		if v == name {
			return k
		}
	}
	var level int
	if strings.HasPrefix(name, NAME_PREFIX) {
		level, _ = strconv.Atoi(name[len(NAME_PREFIX):])
	}
	return level
}

func FileOpen(path string) (*os.File, error) {
	if fi, err := os.Stat(path); err == nil {
		if !fi.IsDir() {
			return nil, fmt.Errorf("open %s: not a directory", path)
		}
	} else if os.IsNotExist(err) {
		if err := os.MkdirAll(path, 0766); err != nil {
			return nil, err
		}
	} else {
		return nil, err
	}

	var currenttime = time.Now().Format("2006-01-02_15.04.05")

	logfile, err := os.OpenFile(path+currenttime+"_LOG.log0", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return nil, err
	}
	return logfile, nil
}

func Trace(a ...interface{}) {
	if TraceLog < Log.level {
		return
	}

	pc := make([]uintptr, 10)
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	file, line := f.FileLine(pc[0])
	fileName := filepath.Base(file)

	nameFull := f.Name()
	nameEnd := filepath.Ext(nameFull)
	funcName := strings.TrimPrefix(nameEnd, ".")

	a = append([]interface{}{funcName + "()", fileName + ":" + strconv.Itoa(line)}, a...)

	Log.Trace(a...)
}

func Tracef(format string, a ...interface{}) {
	if TraceLog < Log.level {
		return
	}

	pc := make([]uintptr, 10)
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	file, line := f.FileLine(pc[0])
	fileName := filepath.Base(file)

	nameFull := f.Name()
	nameEnd := filepath.Ext(nameFull)
	funcName := strings.TrimPrefix(nameEnd, ".")

	a = append([]interface{}{funcName, fileName, line}, a...)

	Log.Tracef("%s() %s:%d "+format, a...)
}

func Debug(a ...interface{}) {
	if DebugLog < Log.level {
		return
	}
	Log.Debug(a...)
}

func Debugf(format string, a ...interface{}) {
	if DebugLog < Log.level {
		return
	}
	Log.Debugf(format, a...)
}

func Info(a ...interface{}) {
	Log.Info(a...)
}

func Warn(a ...interface{}) {
	Log.Warn(a...)
}

func Error(a ...interface{}) {
	Log.Error(a...)
}

func Fatal(a ...interface{}) {
	Log.Fatal(a...)
}

func Infof(format string, a ...interface{}) {
	Log.Infof(format, a...)
}

func Warnf(format string, a ...interface{}) {
	Log.Warnf(format, a...)
}

func Errorf(format string, a ...interface{}) {
	Log.Errorf(format, a...)
}

func Failed(format string, a ...interface{}) {
	Log.Errorf(format, a...)
}

func Fatalf(format string, a ...interface{}) {
	Log.Fatalf(format, a...)
}

// Init deprecated, use InitLog instead
//func Init(a ...interface{}) {
//	os.Stderr.WriteString("warning: use of deprecated Init. Use InitLog instead\n")
//	InitLog(InfoLog, a...)
//}

func /**/ InitLog(logLevel int, a ...interface{}) *Logger {
	writers := []io.Writer{}
	var logFile *os.File
	var err error
	if len(a) == 0 {
		writers = append(writers, ioutil.Discard)
	} else {
		for _, o := range a {
			switch o.(type) {
			case string:
				logFile, err = FileOpen(o.(string))
				if err != nil {
					fmt.Println("error: open log0 file failed")
					os.Exit(1)
				}
				writers = append(writers, logFile)
			case *os.File:
				writers = append(writers, o.(*os.File))
			default:
				fmt.Println("error: invalid log0 location")
				os.Exit(1)
			}
		}
	}
	fileAndStdoutWrite := io.MultiWriter(writers...)
	Log = New(fileAndStdoutWrite, "", log.LUTC|log.Ldate|log.Lmicroseconds, logLevel, logFile)
	return Log
}

func GetLogFileSize() (int64, error) {
	f, e := Log.logFile.Stat()
	if e != nil {
		return 0, e
	}
	return f.Size(), nil
}

func GetMaxLogChangeInterval(maxLogSize int64) int64 {
	if maxLogSize != 0 {
		return (maxLogSize * BYTE_TO_MB)
	} else {
		return (DEFAULT_MAX_LOG_SIZE * BYTE_TO_MB)
	}
}

func CheckIfNeedNewFile() bool {
	logFileSize, err := GetLogFileSize()
	maxLogFileSize := GetMaxLogChangeInterval(0)
	if err != nil {
		return false
	}
	if logFileSize > maxLogFileSize {
		return true
	} else {
		return false
	}
}

func ClosePrintLog() error {
	var err error
	if Log.logFile != nil {
		err = Log.logFile.Close()
	}
	return err
}
