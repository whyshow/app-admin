package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	rotates "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"os"
	"path"
	"time"
)

// 日志
var (
	logFilePath = "./logs/"
	logFileName = "app.log"
)

var (
	LOGGER *logrus.Logger
)

func LogMiddleware() gin.HandlerFunc {
	// 日志文件
	fileName := path.Join(logFilePath, logFileName)
	// 写入文件
	src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println("err", err)
	}
	// 实例化
	LOGGER = logrus.New()
	//设置日志级别
	LOGGER.SetLevel(logrus.InfoLevel)
	//设置输出
	LOGGER.Out = src
	// 设置 rotates
	logWriter, err := rotates.New(
		// 分割后的文件名称
		fileName+".%Y%m%d.log",
		// 生成软链，指向最新日志文件
		rotates.WithLinkName(fileName),
		// 设置最大保存时间(7天)
		rotates.WithMaxAge(7*24*time.Hour),
		// 设置日志切割时间间隔(1天)
		rotates.WithRotationTime(24*time.Hour),
	)

	writerMap:= lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.FatalLevel: logWriter,
		logrus.DebugLevel: logWriter,
		logrus.WarnLevel:  logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.PanicLevel: logWriter,
	}

	LOGGER.AddHook(lfshook.NewHook(writerMap, &logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	}))

	return func(c *gin.Context) {
		//开始时间
		startTime := time.Now()
		//处理请求
		c.Next()
		//结束时间
		endTime := time.Now()
		// 执行时间
		latencyTime := endTime.Sub(startTime)
		//请求方式
		reqMethod := c.Request.Method
		//请求路由
		reqUrl := c.Request.RequestURI
		//状态码
		statusCode := c.Writer.Status()
		//请求ip
		clientIP := c.ClientIP()
		LOGGER.Fatalf("状态码:%3d|等待时间:%13v|请求IP:%15s|访问地址 %s|请求方式:%s",
			statusCode,
			latencyTime,
			clientIP,
			reqUrl,
			reqMethod,
		)
	}
}
