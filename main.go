package main
import (
    "fmt"
    "github.com/hefju/MyLogTest/appconfig"
    "github.com/donnie4w/go-logger/logger"
"time"
)

func main(){
    appconfig.LoadConfig()//初始化配置

    logger.Info("main log test")
    time.Sleep(time.Second*2)
//    logger.Warn("main log test")


    fmt.Println("end")
}
