package config

import (
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
}

type ServerConfig struct {
	Name string
	Addr string
}

type GrpcConfig struct {
	Name string
	Addr string
}

type EtcdConfig struct {
	Addrs []string
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
	// 注册etcd配置
	//conf.ReadEtcdConfig()
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
