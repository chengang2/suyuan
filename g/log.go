package g

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

type Level int

var (
	F *os.File

	DefaultPrefix      = ""
	DefaultCallerDepth = 2

	logger     *log.Logger
	logPrefix  = ""
	levelFlags = []string{"DEBUG", "INFO", "WARN", "ERROR", "FATAL"}
)

const (
	DEBUG Level = iota
	INFO
	WARNING
	ERROR
	FATAL
)

func init() {
	var err error
	filePath := getLogFilePath()
	fileName := getLogFileName()
	F, err = MustOpen(fileName, filePath)
	if err != nil {
		log.Fatalf("logging.Setup err: %v", err)
	}

	logger = log.New(F, DefaultPrefix, log.LstdFlags)
}

func Debug(v ...interface{}) {
	setPrefix(DEBUG)
	logger.Println(v)
}

func Info(v ...interface{}) {
	setPrefix(INFO)
	logger.Println(v)
}

func Warn(v ...interface{}) {
	setPrefix(WARNING)
	logger.Println(v)
}

func Error(v ...interface{}) {
	setPrefix(ERROR)
	logger.Println(v)
}

func Fatal(v ...interface{}) {
	setPrefix(FATAL)
	logger.Fatalln(v)
}

func setPrefix(level Level) {
	_, file, line, ok := runtime.Caller(DefaultCallerDepth)
	if ok {
		logPrefix = fmt.Sprintf("[%s][%s:%d]", levelFlags[level], filepath.Base(file), line)
	} else {
		logPrefix = fmt.Sprintf("[%s]", levelFlags[level])
	}

	logger.SetPrefix(logPrefix)
}

func getLogFilePath() string {
	return fmt.Sprintf("%s%s", AppSetting.RuntimeRootPath, AppSetting.LogSavePath)
}

func getLogFileName() string {
	return fmt.Sprintf("%s%s.%s",
		AppSetting.LogSaveName,
		time.Now().Format(AppSetting.TimeFormat),
		AppSetting.LogFileExt,
	)
}

//func openLogFile(fileName, filePath string) (*os.File, error) {
//	dir, err := os.Getwd()
//	if err != nil {
//		return nil, fmt.Errorf("os.Getwd err: %v", err)
//	}
//
//	src := dir + "/" + filePath
//	perm := file.CheckPermission(src)
//	if perm == true {
//		return nil, fmt.Errorf("file.CheckPermission Permission denied src: %s", src)
//	}
//
//	err = file.IsNotExistMkDir(src)
//	if err != nil {
//		return nil, fmt.Errorf("file.IsNotExistMkDir src: %s, err: %v", src, err)
//	}
//
//	f, err := file.Open(src+fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
//	if err != nil {
//		return nil, fmt.Errorf("Fail to OpenFile :%v", err)
//	}
//
//	return f, nil
//}
