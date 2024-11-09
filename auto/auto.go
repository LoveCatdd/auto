package auto

import (
	"os"

	"github.com/LoveCatdd/util/pkg/lib/core/config"
	"github.com/LoveCatdd/util/pkg/lib/core/log"
	"github.com/LoveCatdd/webctx/pkg/lib/core/web/auth"
	"github.com/LoveCatdd/webctx/pkg/lib/core/web/server"
	"github.com/joho/godotenv"
)

// 用于初始化文件配置
func init() {

	// 相对工作文件路径
	godotenv.Load("../.env")

	environ := os.Getenv("environment")
	config.SetEnviro(environ)

	// get jwt conf
	config.Yaml(auth.JwtConfig)

	// get server name and server port conf
	config.Yaml(server.AppConf)

	// get zap conf
	config.Yaml(log.Config)
	if log.Config.Zap.Enable { // 开启

		log.InitZap()

		defer log.Sync()
	}
}
