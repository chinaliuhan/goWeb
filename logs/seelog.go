package logs

import (
	// "errors"
	"fmt"
	// "io"

	"github.com/cihub/seelog"
)

var Logger seelog.LoggerInterface

func loadAppConfig() {
	appConfig := `
<seelog minlevel="warn">
    <outputs formatid="common">
        <rollingfile type="size" filename="/data/logs/roll.log" maxsize="100000" maxrolls="5"/>
		<filter levels="critical">
            <file path="/data/logs/critical.log" formatid="critical"/>
            <smtp formatid="criticalemail" senderaddress="astaxie@gmail.com" sendername="ShortUrl API" hostname="smtp.gmail.com" hostport="587" username="mailusername" password="mailpassword">
                <recipient address="xiemengjun@gmail.com"/>
            </smtp>
        </filter>
    </outputs>
    <formats>
        <format id="common" format="%Date/%Time [%LEV] %Msg%n" />
	    <format id="critical" format="%File %FullPath %Func %Msg%n" />
	    <format id="criticalemail" format="Critical error on our server!\n    %Time %Date %RelFile %Func %Msg \nSent by Seelog"/>
    </formats>
</seelog>
`
	logger, err := seelog.LoggerFromConfigAsBytes([]byte(appConfig))
	if err != nil {
		fmt.Println(err)
		return
	}
	UseLogger(logger)
}

func init() {
	DisableLog()
	loadAppConfig()
}

// DisableLog禁用所有库日志输出
func DisableLog() {
	Logger = seelog.Disabled
}

// UseLogger使用指定的seelog。输出库日志的LoggerInterface。
//如果在应用程序中使用Seelog日志系统，请使用这个函数。
func UseLogger(newLogger seelog.LoggerInterface) {
	Logger = newLogger
}

//调用错误处理
func main() {
	err := "Info: 错误信息"
	Logger.Info("Start server at:%v", err)
	err = "Critical: 错误信息"
	Logger.Critical("Server err:%v", err)
}
