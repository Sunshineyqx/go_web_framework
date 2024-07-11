package conf

import (
	"giligili/cache"
	"giligili/model"
	"giligili/util"
	"os"

	"github.com/joho/godotenv"
)

// Init 初始化配置项
func Init() {
	// 从本地读取环境变量
	godotenv.Load()

	// 设置日志级别
	util.BuildLogger(os.Getenv("LOG_LEVEL"))

	// 读取翻译文件
	if err := util.InitTrans(os.Getenv("TRANSLATE_LANGUAGE")); err != nil {
		util.Log().Panic("Translator init error: ", err)
	}

	// 连接数据库
	model.Database(os.Getenv("MYSQL_DSN"))
	cache.Redis()
}
