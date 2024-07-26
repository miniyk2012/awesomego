package main

import (
	"context"
	"database/sql"
	"testing"

	// 注册 mysql 数据库驱动
	_ "github.com/go-sql-driver/mysql"
)

type user struct {
	UserID int64
}

func Test_sql(t *testing.T) {
	// 创建 db 实例
	db, err := sql.Open("mysql", "yangkai.04:Yk@5078231@(localhost:3306)/demo1")
	if err != nil {
		t.Error(err)
		return
	}

	// 执行 sql
	ctx := context.Background()
	row := db.QueryRowContext(ctx, "SELECT user_id FROM user ORDER BY created_at DESC limit 1")
	if row.Err() != nil {
		t.Error(row.Err())
		return
	}
	// 解析结果
	var u user
	if err = row.Scan(&u.UserID); err != nil {
		t.Error(err)
		return
	}
	t.Log(u.UserID)
}
