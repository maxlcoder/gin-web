package setting

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"os"
)

// 配置
type Mysql struct {
	Host string
	Port string
	DB string
	User string
	Password string
}

type Server struct {
	HTTTPort string
	PageSize int
	JwtSecret string
}

var MysqlSetting = &Mysql{}
var ServerSetting = &Server{}

func init() {
	mode := os.Getenv(gin.EnvGinMode)
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./conf")
	viper.SetConfigName(mode)

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file : %w \n", err))
	}

	MysqlSetting.Host = viper.GetString("mysql.host")
	MysqlSetting.Port = viper.GetString("mysql.port")
	MysqlSetting.DB = viper.GetString("mysql.db")
	MysqlSetting.User = viper.GetString("mysql.user")
	MysqlSetting.Password = viper.GetString("mysql.password")

	ServerSetting.HTTTPort = ":" + viper.GetString("app.http_port")
	ServerSetting.PageSize = viper.GetInt("app.page_size")
	ServerSetting.JwtSecret = viper.GetString("app.jwt_secret")

}