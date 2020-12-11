package logs

import (
	"github.com/lestrrat-go/file-rotatelogs"
	"github.com/pkg/errors"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"os"
	"path/filepath"
	"raspberry/utils"
	"strings"
	"time"
)

const (
	infoPath  = "/log/raspberry/info/"
	errorPath = "/log/raspberry/error/"
)

// Setup
func Setup() {
	home, err := utils.Home()
	if err != nil {
		logrus.Errorf("Failed to get home dir: %s\n", err.Error())
	}
	if err := os.MkdirAll(home+infoPath, 0755); err != nil {
		logrus.Errorf("Failed to make info dir: %s\n", err.Error())
	}
	if err := os.MkdirAll(home+errorPath, 0755); err != nil {
		logrus.Errorf("Failed to make err dir: %s\n", err.Error())
	}
	writer, err := rotatelogs.New(
		home+infoPath+"%Y-%m-%d"+".log",
		rotatelogs.WithLinkName("log.log"),        // 生成软链，指向最新日志文件
		rotatelogs.WithMaxAge(7*24*time.Hour),     // 文件最大保存时间
		rotatelogs.WithRotationTime(24*time.Hour), // 日志切割时间间隔
	)
	if err != nil {
		logrus.Errorf("config local file system logger error. %+v", errors.WithStack(err))
	}

	errorWriter, err := rotatelogs.New(
		home+errorPath+"%Y-%m-%d"+".log",
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
