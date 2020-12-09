package Logs

import (
	"github.com/lestrrat-go/file-rotatelogs"
	"github.com/pkg/errors"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const (
	infoPath  = "/var/log/raspberry/info/"
	errorPath = "/var/log/raspberry/error/"
)

// Setup
func Setup() {
	if err := os.MkdirAll(infoPath, 0666); err != nil {
		logrus.Errorf("Failed to make info dir: %s\n", err.Error())
	}
	if err := os.MkdirAll(errorPath, 0666); err != nil {
		logrus.Errorf("Failed to make err dir: %s\n", err.Error())
	}
	writer, err := rotatelogs.New(
		infoPath+"%Y-%m-%d"+".log",
		rotatelogs.WithLinkName("log.log"),        // 生成软链，指向最新日志文件
		rotatelogs.WithMaxAge(7*24*time.Hour),     // 文件最大保存时间
		rotatelogs.WithRotationTime(24*time.Hour), // 日志切割时间间隔
	)
	if err != nil {
		logrus.Errorf("config local file system logger error. %+v", errors.WithStack(err))
	}

	errorWriter, err := rotatelogs.New(
		errorPath+"%Y-%m-%d"+".log",
		rotatelogs.WithLinkName("error.log"),      // 生成软链，指向最新日志文件
		rotatelogs.WithMaxAge(7*24*time.Hour),     // 文件最大保存时间
		rotatelogs.WithRotationTime(24*time.Hour), // 日志切割时间间隔
	)
	if err != nil {
		logrus.Errorf("config local file system logger error. %+v", errors.WithStack(err))
	}
	lfHook := lfshook.NewHook(lfshook.WriterMap{
		logrus.DebugLevel: writer, // 为不同级别设置不同的输出目的
		logrus.InfoLevel:  writer,
		logrus.WarnLevel:  writer,
		logrus.ErrorLevel: errorWriter,
		logrus.FatalLevel: writer,
		logrus.PanicLevel: writer,
	}, &logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
		PrettyPrint:     false,
		//FieldMap: logrus.FieldMap{
		//	"host": setting.GlobalSetting.LocalHost,
		//},
	})
	logrus.SetReportCaller(true)
	logrus.AddHook(lfHook)
	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors: true,
	})
	//logrus.SetOutput()
}

func getCurrentDirectory() string {
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	return strings.Replace(dir, "\\", "/", -1)
}
