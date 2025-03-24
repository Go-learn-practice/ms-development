package gorms

import (
	"context"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"test.com/devCommon/logs"
	"test.com/devProject/config"
	"test.com/devProject/internal/data/menu"
	"test.com/devProject/internal/data/pro"
	"test.com/devProject/internal/data/task"
)

var _db *gorm.DB

func init() {
	// 配置 mysql 连接参数
	username := config.Conf.Mysql.Username // 账号
	password := config.Conf.Mysql.Password // 密码
	host := config.Conf.Mysql.Host         // 数据库地址，可以是Ip或者域名
	port := config.Conf.Mysql.Port         // 数据库端口
	Dbname := config.Conf.Mysql.Db         // 数据库名
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, port, Dbname)
	// 连接数据库
	var err error
	_db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	logs.LG.Info("连接数据库成功")
	if err != nil {
		panic("连接数据库失败, error=" + err.Error())
	}
	// 自动迁移
	err = _db.AutoMigrate(
		&menu.MsProjectMenu{},
		&pro.MsProject{},
		&pro.MsProjectMember{},
		&pro.MsProjectCollection{},
		&pro.MsProjectTemplate{}, &task.MsTaskStagesTemplate{},
	)
	if err != nil {
		panic("自动迁移失败, error=" + err.Error())
	}
}

func GetDB() *gorm.DB {
	return _db
}

type GormConn struct {
	db *gorm.DB
	tx *gorm.DB
}

func (g *GormConn) Begin() {
	g.tx = GetDB().Begin()
}

func New() *GormConn {
	return &GormConn{
		db: GetDB(),
	}
}

func NewTran() *GormConn {
	return &GormConn{db: GetDB(), tx: GetDB()}
}

func (g *GormConn) Session(ctx context.Context) *gorm.DB {
	return g.db.Session(&gorm.Session{Context: ctx})
}

func (g *GormConn) Rollback() {
	g.tx.Rollback()
}

func (g *GormConn) Commit() {
	g.tx.Commit()
}

func (g *GormConn) Tx(ctx context.Context) *gorm.DB {
	return g.tx.WithContext(ctx)
}
