package model

import (
	"net/url"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Init() {
	host := viper.GetString("DB_HOST")
	port := viper.GetString("DB_PORT")
	charset := viper.GetString("DB_CHARSET")
	user := viper.GetString("DB_USER")
	pass := viper.GetString("DB_PASS")
	name := viper.GetString("DB_NAME")
	loc := viper.GetString("DB_LOC")

	dsn := user + ":" + pass + "@tcp(" + host + ":" + port + ")/" +
		name + "?charset=" + charset + "&parseTime=True&loc=" + url.QueryEscape(loc)

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         256,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{})

	if err != nil {
		panic("conn mysql fail: " + err.Error())
	}

	viper.Set("DB", db)
}
