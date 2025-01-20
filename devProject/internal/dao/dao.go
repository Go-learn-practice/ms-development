package dao

import (
	"test.com/devProject/internal/database"
	"test.com/devProject/internal/database/gorms"
)

type Transaction struct {
	conn database.DbConn
}

func (t *Transaction) Action(fn func(conn database.DbConn) error) error {
	t.conn.Begin()
	err := fn(t.conn)
	if err != nil {
		// 回滚
		t.conn.Rollback()
		return err
	}
	// 提交
	t.conn.Commit()
	return nil
}

func NewTransaction() *Transaction {
	return &Transaction{
		conn: gorms.NewTran(),
	}
}
