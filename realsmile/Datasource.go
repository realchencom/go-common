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

	if DB, err := gorm.Open(mysql.Open(this.Mysql), &gorm.Config{Logger: nil}); err != nil {
		Log.Errorf("Failed to get MySQL database: %v !errorMsg:%v", this.Mysql, err)
		panic(fmt.Errorf("failed to get MySQL database:  %v !errorMsg:%v", this.Mysql, err))
	} else {
		return DB, nil
	}
}
func GetMylSqlByConStr(conStr string) (*gorm.DB, error) {
	if DB, err := gorm.Open(mysql.Open(conStr), &gorm.Config{Logger: nil}); err != nil {
		Log.Errorf("Failed to get MySQL database: %v !errorMsg:%v", conStr, err)
		panic(fmt.Errorf("failed to get MySQL database: %v !errorMsg:%v", conStr, err))
	} else {
		return DB, nil
	}
}
func Saves[T any](DB *gorm.DB, slices []T, size int) error {
	//每组大小
	amount := len(slices)    //总数
	group := amount / size   //分组
	surplus := amount % size //余数
	if surplus != 0 {
		group++
	}
	for i := 0; i < group; i++ {
		var element interface{}
		if (i + 1) == group {
			element = slices[i*size : i*size+surplus]
		} else {
			element = slices[i*size : (i+1)*size]
		}
		tx := DB.Save(element)
		if tx.Error != nil {
			return tx.Error
		}
	}
	return nil

}
