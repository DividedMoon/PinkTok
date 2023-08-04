package config

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	hertzzap "github.com/hertz-contrib/logger/zap"
	"io"
	"log"
	"os"
	"path"
	"time"
)

const zapConfigPath = "./configs/zap-config.yaml"

type LogConfig struct {
	OutputDir  string `yaml:"outputDir"`
	MaxSize    int    `yaml:"maxSize"`
	MaxBackups int    `yaml:"maxBackups"`
	MaxAge     int    `yaml:"maxAge"`
	Compress   bool   `yaml:"compress"`
	LogLevel   string `yaml:"logLevel"`
}

func initLogger() error {
	var config LogConfig
	if err := ReadConfigFromYAML(zapConfigPath, &config); err != nil {
		log.Println(err.Error())
		return err
	}

	// Customizable output directory.
	var logFilePath string
	logFilePath = config.OutputDir + "/logs/"
	if err := os.MkdirAll(logFilePath, 0o777); err != nil {
		log.Println(err.Error())
		return err
	}

	// Set filename to date
	logFileName := time.Now().Format("2006-01-02") + ".log"
	fileName := path.Join(logFilePath, logFileName)
	if _, err := os.Stat(fileName); err != nil {
		if _, err := os.Create(fileName); err != nil {
			log.Println(err.Error())
			return err
		}
	}

	// For zap detailed settings, please refer to https://github.com/hertz-contrib/logger/tree/main/zap and https://github.com/uber-go/zap
	logger := hertzzap.NewLogger()
	// Provides compression and deletion
	lumberjackLogger := &lumberjack.Logger{
		Filename:   fileName,
		MaxSize:    config.MaxSize,    // A file can be up to 20M.
		MaxBackups: config.MaxBackups, // Save up to 5 files at the same time.
		MaxAge:     config.MaxAge,     // A file can exist for a maximum of 10 days.
		Compress:   config.Compress,   // Compress with gzip.
	}

	//logger.SetOutput(lumberjackLogger)
	switch config.LogLevel {
	case "debug":
		logger.SetLevel(hlog.LevelDebug)
	case "warn":
		logger.SetLevel(hlog.LevelWarn)
	case "info":
		logger.SetLevel(hlog.LevelInfo)
	case "fatal":
		logger.SetLevel(hlog.LevelFatal)
	case "error":
		logger.SetLevel(hlog.LevelError)
	case "notice":
		logger.SetLevel(hlog.LevelNotice)
	case "trace":
		logger.SetLevel(hlog.LevelTrace)
	}

	// if you want to output the log to the file and the stdout at the same time, you can use the following codes

	fileWriter := io.MultiWriter(lumberjackLogger, os.Stdout)
	logger.SetOutput(fileWriter)
	hlog.SetLogger(logger)
	return nil

}
