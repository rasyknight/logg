package logg

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/user"
	"runtime"
	"strconv"
	"strings"
)

type DebugLevel int

var DebugDesc = []string{"FORCE", "ERROR", "WARNING", "NOTICE", "INFO", "DEBUG"}

const (
	FORCE = iota
	ERROR
	WARNING
	NOTICE
	INFO
	DEBUG
)

type logInfo struct {
	LogLevel int
}

var logInf logInfo

func intMax(i, j int) int {
	if i > j {
		return i
	} else {
		return j
	}
}
func InitLog() {
	usr, _ := user.Current()
	homeDir := usr.HomeDir
	logFileName := homeDir + "/log/service.log"
	logFile, err := os.OpenFile(logFileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	log.SetOutput(logFile)
	SetDebugLevel(5)
}

func GetGid() (gid uint64) {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, err := strconv.ParseUint(string(b), 10, 64)
	if err != nil {
		panic(err)
	}
	return n
}
func logWrite(Level, str string) {
	_, file, lineNo, _ := runtime.Caller(2)
	Files := strings.Split(file, "/")
	FileLen := len(Files)
	var LogFileName string
	begin := intMax(FileLen-2, 0)
	for i := begin; i < FileLen; i++ {
		if i == begin {
			LogFileName = Files[i]
		} else {
			LogFileName = LogFileName + "/" + Files[i]
		}
	}
	len := len(str)
	if str[len-1] == 10 {
		log.Printf("[%s] F[%s] L[%d] T[%d]: %s", Level, LogFileName, lineNo, GetGid(), str)
	} else {
		log.Printf("[%s] F[%s] L[%d] T[%d]: %s\n", Level, LogFileName, lineNo, GetGid(), str)
	}
}
func Debug(message string, args ...interface{}) {
	if logInf.LogLevel < DEBUG {
		return
	}
	logWrite("DEBUG", fmt.Sprintf(message, args...))
}

func Info(message string, args ...interface{}) {
	if logInf.LogLevel < INFO {
		return
	}
	logWrite("INFO", fmt.Sprintf(message, args...))
}
func Notice(message string, args ...interface{}) {
	if logInf.LogLevel < NOTICE {
		return
	}
	logWrite("NOTICE", fmt.Sprintf(message, args...))
}
func Warning(message string, args ...interface{}) {
	if logInf.LogLevel < WARNING {
		return
	}
	logWrite("WARNING", fmt.Sprintf(message, args...))
}
func Error(message string, args ...interface{}) {
	if logInf.LogLevel < ERROR {
		return
	}
	logWrite("ERROR", fmt.Sprintf(message, args...))
}
func Force(message string, args ...interface{}) {
	str := fmt.Sprintf(message, args...)
	logWrite("FORCE", str)
}
func Debugln(val ...any) {
	if logInf.LogLevel < DEBUG {
		return
	}
	str := fmt.Sprintln(val...)
	logWrite("DEBUG", str)
}
func Infoln(val ...any) {
	if logInf.LogLevel < INFO {
		return
	}
	str := fmt.Sprintln(val...)
	logWrite("INFO", str)
}
func Noticeln(val ...any) {
	if logInf.LogLevel < NOTICE {
		return
	}
	str := fmt.Sprintln(val...)
	logWrite("NOTICE", str)
}
func Warningln(val ...any) {
	if logInf.LogLevel < WARNING {
		return
	}
	str := fmt.Sprintln(val...)
	logWrite("WARNING", str)
}
func Errorln(val ...any) {
	if logInf.LogLevel < ERROR {
		return
	}
	str := fmt.Sprintln(val...)
	logWrite("ERROR", str)
}
func Forceln(val ...any) {
	str := fmt.Sprintln(val...)
	logWrite("FORCE", str)
}
func SetDebugLevel(level int) int {
	OldLevel := logInf.LogLevel
	logInf.LogLevel = level
	str := fmt.Sprintf("Debug Level change from %s to %s", GetDebugLevel(OldLevel), GetDebugLevel(level))
	logWrite("Force", str)
	return OldLevel
}
func GetDebugLevel(level int) string {
	return DebugDesc[level]
}
func GetDebug() int {
	return logInf.LogLevel
}

