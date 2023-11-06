package svc

import (
	"fmt"
	"greet/common/helper"
	"greet/internal/config"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config
	DB     sqlx.SqlConn
}

func NewServiceContext(c config.Config) *ServiceContext {
	mysql := c.Mysql
	args := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=true&writeTimeout=%s",
		mysql.Username,
		mysql.Password,
		mysql.Host,
		mysql.Port,
		mysql.Database,
		mysql.Charset,
		mysql.Timeout,
	)

	db, _ := helper.InitMysql(args)

	return &ServiceContext{
		Config: c,
		DB:     db,
	}
}
