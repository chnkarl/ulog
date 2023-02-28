package ulog

import (
	"encoding/json"
	"log"
	"os"
	"reflect"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type Log struct {
	C *gin.Context
}

func (l *Log) Info(logInfo interface{}) {
	logString := l.formatString(logInfo)
	l.writeLog("info", logString)
}

func (l *Log) Error(logInfo interface{}) {
	logString := l.formatString(logInfo)
	l.writeLog("error", logString)
}

func (l *Log) Warning(logInfo interface{}) {
	logString := l.formatString(logInfo)
	l.writeLog("warning", logString)
}

func (l *Log) Debug(logInfo interface{}) {
	logString := l.formatString(logInfo)
	l.writeLog("debug", logString)
}

func (l *Log) writeLog(level string, logInfo interface{}) {
	file := l.openFile()

	ip := l.getIp()

	_log := log.New(file, "["+level+"] "+ip+" ", log.Ldate|log.Ltime|log.Lshortfile)

	_log.Println(logInfo)

	defer file.Close()
}

/**
 * open log file
 */
func (l *Log) openFile() *os.File {

	// get log file path
	filePath := l.getLogPath()

	// call os.OpenFile
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)

	if err != nil {
		panic("open log file fail: " + err.Error())
	}

	return file
}

/**
 * format any type of logInfo to string for write in log file
 */
func (l *Log) formatString(logInfo interface{}) interface{} {

	// get logInfo type
	logType := reflect.TypeOf(logInfo).Name()

	var _logInfo interface{} = ""

	if logType == "string" {
		// if log type is string ,just output
		_logInfo = logInfo
	} else {
		// if log type is struct\map\slice convert to bytes string
		logBytes, _ := json.Marshal(&logInfo)

		_logInfo = string(logBytes)
	}

	return _logInfo
}

/**
 * get client ip to write in log file
 */
func (l *Log) getIp() string {
	ipPort := l.C.Request.RemoteAddr

	ip := strings.Split(ipPort, ":")[0]

	return ip
}

/**
 * get log path
 * create log file every day
 */
func (l *Log) getLogPath() string {

	const (
		YYYY = "2006"
		MM   = "01"
		DD   = "02"
	)

	now := time.Now().UTC()

	// get year 2006
	currentYear := now.Format(YYYY)

	// get month 01
	currentMonth := now.Format(MM)

	// get project path
	rootDir, err := os.Getwd()

	if err != nil {
		panic(err)
	}

	// format log dir path: logs/2006/01
	dir := rootDir + "/logs/" + currentYear + "/" + currentMonth

	// create dir if not exist
	mkDirErr := os.MkdirAll(dir, 0666)

	if mkDirErr != nil {
		panic("create log dir fail: " + mkDirErr.Error())
	}

	// return file path: logs/2006/01/02
	return dir + "/" + now.Format(DD) + ".log"
}
