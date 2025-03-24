package database

// DbConn 定义事务接口
type DbConn interface {
	Begin()
	Rollback()
	Commit()
}
