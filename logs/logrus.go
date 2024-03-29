package logs

import (
	"os"

	log "github.com/Sirupsen/logrus"
)

func init() {
	// 设置日志格式化为JSON而不是默认的ASCII
	log.SetFormatter(&log.JSONFormatter{})

	// 设置输出stdout而不是默认的stderr，也可以是一个文件
	// 为当前logrus实例设置消息的输出，同样地，
	// 可以设置logrus实例的输出到任意io.writer
	log.SetOutput(os.Stdout)

	// 设置只记录严重或以上警告,这里如果要设置其他级别,直接点进去看就行了,错误级别这种东西都很直观
	log.SetLevel(log.WarnLevel)
}

func main() {
	//下面全是用来记录日志的
	log.WithFields(log.Fields{
		"animal": "walrus",
		"size":   10,
	}).Info("A group of walrus emerges from the ocean")

	log.WithFields(log.Fields{
		"omg":    true,
		"number": 122,
	}).Warn("The group's number increased tremendously!")

	log.WithFields(log.Fields{
		"omg":    true,
		"number": 100,
	}).Fatal("The ice breaks!")

	// 通过日志语句重用字段
	// logrus.Entry返回自WithFields()
	contextLogger := log.WithFields(log.Fields{
		"common": "this is a common field",
		"other":  "I also should be logged always",
	})

	contextLogger.Info("I'll be logged with common and other field")
	contextLogger.Info("Me too")
}
