package config

import (
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"log"
	"os"
	"test.com/devCommon/logs"
)

var Conf = InitConfig()

type Config struct {
	viper *viper.Viper
	SC    *ServerConfig
	GC    *GrpcConfig
	Etcd  *EtcdConfig
	Mysql *MysqlConfig
	Jwt   *JwtConfig
}

type ServerConfig struct {
	Name string
	Addr string
}

type GrpcConfig struct {
	Name    string
	Addr    string
	Version string
	Weight  int64
}

type EtcdConfig struct {
	Addrs []string
}

type MysqlConfig struct {
	Username string
	Password string
	Host     string
	Port     int
	Db       string
}

type JwtConfig struct {
	AccessExp     int64
	RefreshExp    int64
	AccessSecret  string
	RefreshSecret string
}

func InitConfig() *Config {
	conf := &Config{viper: viper.New()}
	workDir, _ := os.Getwd()
	conf.viper.SetConfigName("config")
	conf.viper.SetConfigType("yaml")
	conf.viper.AddConfigPath(workDir + "/config")
	err := conf.viper.ReadInConfig()
	if err != nil {
		log.Fatalf("read config fail: %v \n", err)
	}
	conf.ReadServerConfig()
	// 日志
	conf.InitZapLog()
	conf.ReadGrpcConfig()
	// ReadEtcdConfig()
	conf.InitMysqlConfig()
	conf.InitJwtConfig()
	return conf
}

func (c *Config) InitZapLog() {
	//从配置中读取日志配置，初始化日志
	lc := &logs.LogConfig{
		DebugFileName: c.viper.GetString("zap.debugFileName"),
		InfoFileName:  c.viper.GetString("zap.infoFileName"),
		WarnFileName:  c.viper.GetString("zap.warnFileName"),
		MaxSize:       c.viper.GetInt("maxSize"),
		MaxAge:        c.viper.GetInt("maxAge"),
		MaxBackups:    c.viper.GetInt("maxBackups"),
	}
	err := logs.InitLogger(lc)
	if err != nil {
		log.Fatalln(err)
	}
}

// ReadServerConfig 读取服务配置
func (c *Config) ReadServerConfig() {
	sconf := &ServerConfig{}
	sconf.Name = c.viper.GetString("server.name")
	sconf.Addr = c.viper.GetString("server.addr")
	c.SC = sconf
}

// ReadRedisConfig 读取redis配置
func (c *Config) ReadRedisConfig() *redis.Options {
	redisConf := &redis.Options{
		Addr:     c.viper.GetString("redis.host") + ":" + c.viper.GetString("redis.port"),
		Password: c.viper.GetString("redis.password"),
		DB:       c.viper.GetInt("redis.db"),
	}
	return redisConf
}

func (c *Config) ReadGrpcConfig() {
	gconf := &GrpcConfig{}
	gconf.Name = c.viper.GetString("grpc.name")
	gconf.Addr = c.viper.GetString("grpc.addr")
	gconf.Version = c.viper.GetString("grpc.version")
	gconf.Weight = c.viper.GetInt64("grpc.weight")
	c.GC = gconf
}

// ReadEtcdConfig etcd 服务配置
func (c *Config) ReadEtcdConfig() {
	etcdConf := &EtcdConfig{}
	var addrs []string
	err := c.viper.UnmarshalKey("etcd.addrs", &addrs)
	if err != nil {
		log.Fatalln(err)
	}
	etcdConf.Addrs = addrs
	c.Etcd = etcdConf
}

func (c *Config) InitMysqlConfig() {
	mysqlConf := &MysqlConfig{
		Username: c.viper.GetString("mysql.username"),
		Password: c.viper.GetString("mysql.password"),
		Host:     c.viper.GetString("mysql.host"),
		Port:     c.viper.GetInt("mysql.port"),
		Db:       c.viper.GetString("mysql.dbname"),
	}
	c.Mysql = mysqlConf
}

func (c *Config) InitJwtConfig() {
	jwtConf := &JwtConfig{
		AccessExp:     c.viper.GetInt64("jwt.accessExp"),
		RefreshExp:    c.viper.GetInt64("jwt.refreshExp"),
		AccessSecret:  c.viper.GetString("jwt.accessSecret"),
		RefreshSecret: c.viper.GetString("jwt.refreshSecret"),
	}
	c.Jwt = jwtConf
}
