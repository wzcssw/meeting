package lib

import (
	"github.com/jinzhu/configor"
)

// DatabaseConfigFile 数据库配置对象
var DatabaseConfigFile DatabaseConfigFileObject

// ENV 运行环境
var ENV string

// DatabaseObject s
type DatabaseObject struct {
	Database string
	Pool     int
	Username string
	Password string
	Host     string
	Port     string
	Socket   string
}

// DatabaseConfigFileObject  DatabaseConfigFileObject
type DatabaseConfigFileObject struct {
	Develop    DatabaseObject
	Staging    DatabaseObject
	Production DatabaseObject
}

// GetDatabase GetDatabase
func (dcfo *DatabaseConfigFileObject) GetDatabase(env string) DatabaseObject {
	switch env {
	case "develop":
		return dcfo.Develop
	case "staging":
		return dcfo.Staging
	case "production":
		return dcfo.Production
	default:
		return DatabaseObject{}
	}
}

// GetDatabase 获得当前环境下的数据库配置
func GetDatabase() DatabaseObject {
	return DatabaseConfigFile.GetDatabase(ENV)
}

func init() {
	configor.Load(&DatabaseConfigFile, "database.yml")
}
