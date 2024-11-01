package auto

import (
	"os"

	"github.com/LoveCatdd/util/pkg/lib/core/config"
	"github.com/LoveCatdd/webctx/pkg/lib/core/web/auth"
	"github.com/joho/godotenv"
)

// 用于初始化文件配置
func init() {

	// 相对工作文件路径
	godotenv.Load("../.env")

	environ := os.Getenv("environment")

	config.SetEnviro(environ)

	config.Yaml(auth.JwtConfig)

}
