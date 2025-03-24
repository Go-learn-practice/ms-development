package dao

import (
	"test.com/devProject/internal/database"
	"test.com/devProject/internal/database/gorms"
)

// Transaction 事务
type Transaction struct {
	conn database.DbConn
}

func (tx *Transaction) Action(fn func(conn database.DbConn) error) error {
	tx.conn.Begin()
	err := fn(tx.conn)
	if err != nil {
		// 回滚
		tx.conn.Rollback()
		return err
	}
	// 提交
	tx.conn.Commit()
	return nil
}

func NewTransaction() *Transaction {
	return &Transaction{
		conn: gorms.NewTran(),
	}
}
