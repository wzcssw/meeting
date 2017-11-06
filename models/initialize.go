package models

import (
	"time"
	"txmeeting/lib"
	// 载入mysql驱动
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// Orm gorm.DB
var Orm *gorm.DB

// InitializeDataBase 连接数据库
func InitializeDataBase() {
	databaseConfig := lib.GetDatabase()
	linkString := databaseConfig.Username + ":" + databaseConfig.Password + "@tcp(" + databaseConfig.Host + ":" +
		databaseConfig.Port + ")/" + databaseConfig.Database + "?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", linkString)
	if err != nil {
		panic(err)
	}
	Orm = db
	// defer db.Close()
}

// Model 基本模型的定义
type Model struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"update_at"`
}

// Result 通用的返回数据 函分页信息
type Result struct {
	Success  bool        `json:"success"`
	Page     int         `json:"page"`
	PageSize int         `json:"page_size"`
	AllCount int         `json:"all_count"`
	Data     interface{} `json:"data"`
}

// NewResult 得到实例，主要为了定义默认值
func NewResult(page int, size int) Result {
	return Result{Success: true, Page: page, PageSize: size}
}

// OffSet GetOffSet
func (result *Result) OffSet() int {
	return (result.Page - 1) * result.PageSize
}
