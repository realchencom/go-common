package realsmile

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Datasource struct {
	Mysql    string `mapstructure:"mysql"`
	Sqlite   string `mapstructure:"sqlite"`
	LogLevel logger.LogLevel
}

func GetMylSqlDB() (*gorm.DB, error) {
	var this *Datasource
	if err := Con.Sub("datasource").Unmarshal(&this); err != nil {
		panic(fmt.Errorf("unmarshal datasource conf failed, err:%s \n", err))
	}
	//返回不使用默认的日志文件
	DB, err := gorm.Open(mysql.Open(this.Mysql), &gorm.Config{Logger: nil})
	return DB, err
}
