package logger

import (
	"errors"
	"fmt"
	"os"
	"path"
	"runtime"
	"strings"
	"time"
)

type LogLevel uint16

const (
	UNKNOW LogLevel = iota
	DEBUG
	TRACE
	INFO
	WARNING
	ERROR
	FATAL
)

type Logger struct {
	OutPutPath   string
	OutPutMethod string
	FileObj      *os.File
	Level        LogLevel
	MaxSize      int
	LogChan      chan *string
}

func parseLevel(levelstr string) (LogLevel, error) {
	str := strings.ToUpper(levelstr)
	switch str {
	case "DEBUG":
		return DEBUG, nil
	case "TRACE":
		return TRACE, nil
	case "INFO":
		return INFO, nil
	case "WARNING":
		return WARNING, nil
	case "ERROR":
		return ERROR, nil
	case "FATAL":
		return FATAL, nil
	default:
		err := errors.New("未知的日志类型")
		return UNKNOW, err
	}
}

func NewLogger(OutPutPath, OutPutMethod, Levelstr string, MaxSize int) (*Logger, error) {
	LogChan := make(chan *string, 100)
	var err error
	var FileObj *os.File
	if OutPutMethod != "stdout" && OutPutMethod != "file" {
		return nil, errors.New("请输入正确的输出方式")
	}
	if OutPutMethod == "file" {
		FileObj, err = os.OpenFile(OutPutPath+time.Now().Format("20060102150405000"), os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
		if err != nil {
			fmt.Printf("open file error:%v", err)
			return nil, err
		}
	}
	Level, err := parseLevel(Levelstr)
	if err != nil {
		panic(err) //这样就不用返回任何东西了
	}
	result := Logger{
		OutPutPath,
		OutPutMethod,
		FileObj,
		Level,
		MaxSize,
		LogChan,
	}
	go result.writeBack()
	return &result, nil
}

func (l *Logger) FormaOutPut(level string, skip int, format string, a ...interface{}) {
	msg := fmt.Sprintf(format, a...)
	le, _ := parseLevel(level)
	if l.Level <= le {
		now := time.Now().Format("2006-02-02 15:04:05")
		//获取log的文件名 行号等信息 runtime.Caller()
		//这个函数的参数代表函数作用域的层级，层级越深参数从0越高
		//如果参数为0 则会输出FormaOutPut这个函数
		//如果参数为1 则会输出Debug Trace等函数
		//依次类推
		pc, file, line, ok := runtime.Caller(skip)
		if !ok {
			fmt.Println("runtime.Caller() failed")
			return
		}
		funcName := runtime.FuncForPC(pc).Name()
		fileName := path.Base(file)
		if l.OutPutMethod == "stdout" {
			fmt.Printf("[%s] [%s] [%s:%s:%d] %s\n", now, level, fileName, funcName, line, msg)
		} else {
			l.CheckSlice()
			s := fmt.Sprintf("[%s] [%s] [%s:%s:%d] %s\n", now, level, fileName, funcName, line, msg)
			select {
			case l.LogChan <- &s:
			default:
			}

			//fmt.Fprintf(l.FileObj, "[%s] [%s] [%s:%s:%d] %s\n", now, level, fileName, funcName, line, msg)
		}
	}
}

func (l *Logger) CheckSlice() {
	fileInfo, _ := l.FileObj.Stat()
	fmt.Println(fileInfo.Size())
	if fileInfo.Size() >= int64(l.MaxSize) {
		l.FileObj.Close()
		//打开一个新文件
		fileObj, err := os.OpenFile(l.OutPutPath+time.Now().Format("20060102150405000"), os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
		if err != nil {
			fmt.Println("新文件打开失败")
			return
		}
		l.FileObj = fileObj
	}
}

func (l *Logger) PrintAny(level string, skip int, format string, a ...interface{}) {
	l.FormaOutPut(level, skip, format, a...)
}

func (l *Logger) writeBack() {
	for s := range l.LogChan {
		fmt.Fprintf(l.FileObj, *s)
	}
}
