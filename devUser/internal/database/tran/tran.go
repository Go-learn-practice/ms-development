package tran

import "test.com/devUser/internal/database"

// Transaction 事务的操作 与数据库有关 注入数据库的连接 gorm.db
type Transaction interface {
	Action(func(conn database.DbConn) error) error
}
