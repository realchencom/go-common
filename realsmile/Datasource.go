package realsmile

import (
	"fmt"
	"github.com/pkg/errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"reflect"
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
		Log.Errorf("Failed to get MySQL database: %v", err)
		panic(fmt.Errorf("failed to get MySQL database: %v", err))
	} else {
		return DB, nil
	}
}
func Saves(DB *gorm.DB, slices interface{}, size int) error {
	reflectValue := reflect.Indirect(reflect.ValueOf(slices))
	for reflectValue.Kind() == reflect.Ptr || reflectValue.Kind() == reflect.Interface {
		reflectValue = reflect.Indirect(reflectValue)
	}
	value := slices.([]interface{})
	switch reflectValue.Kind() {
	case reflect.Slice:
		{
			//每组大小
			amount := len(value)     //总数
			group := amount / size   //分组
			surplus := amount % size //余数
			if surplus != 0 {
				group++
			}
			for i := 0; i < group; i++ {
				var element interface{}
				if (i + 1) == group {
					element = value[i*size : i*size+surplus]
				} else {
					element = value[i*size : (i+1)*size]
				}
				tx := DB.Save(&element)
				if tx.Error != nil {
					return tx.Error
				}
			}
			return nil
		}
	default:
		{
			return errors.New("目前只支持切片批量入库。")
		}
	}

}
