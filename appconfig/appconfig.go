package appconfig

import (
    "github.com/donnie4w/go-logger/logger"
)



func init(){
   // LoadConfig()
}
func LoadConfig(){

    //logger.SetConsole(false)//默认是输出到控制台的, 所以logger.SetConsole(true) 写不写都无所谓
    logger.SetRollingDaily("./log", "test.log")//如果没有log文件夹, 需要新增文件夹
    logger.SetLevel(logger.DEBUG)
    logger.Debug("loger init")

}
