package helper

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

func InitMysql(args string) (conn sqlx.SqlConn, err error) {

	conn = sqlx.NewMysql(args)

	var count int
	query := "select count(*) from user"
	err = conn.QueryRowCtx(context.Background(), &count, query)
	if err != nil {
		fmt.Println("MYSQL ERROR", err)
	}
	fmt.Println("MYSQL OK", count)
	fmt.Println("MYSQL INFO", args)

	return conn, err
}
