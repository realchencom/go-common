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

func (this *Datasource) GetMylSqlDB() (*gorm.DB, error) {

	if err := Con.Sub("datasource").Unmarshal(&this); err != nil {
		panic(fmt.Errorf("unmarshal datasource conf failed, err:%s \n", err))
	}
	//DB, err := gorm.Open(mysql.Open(this.Mysql), &gorm.Config{Logger: this.LogMode(-1)})
	DB, err := gorm.Open(mysql.Open(this.Mysql), &gorm.Config{Logger: nil})
	return DB, err
}

//	func GetSqliteDB() (*gorm.DB, error) {
//		var this Datasource
//		if err := Con.Sub("datasource").Unmarshal(&this); err != nil {
//			panic(fmt.Errorf("unmarshal datasource conf failed, err:%s \n", err))
//		}
//		DB, err := gorm.Open(sqlite.Open(this.Sqlite), &gorm.Config{Logger: this.LogMode(-1)})
//		return DB, err
//	}
//func (this *Datasource) LogMode(level logger.LogLevel) logger.Interface {
//	if -1 == level {
//		switch logConf.Level {
//		case "debug":
//			//自定义日志级别：自定义Warn级别
//			this.LogLevel = logger.Silent
//		case "info":
//			this.LogLevel = logger.Info
//		case "warn":
//			this.LogLevel = logger.Warn
//		case "error":
//			this.LogLevel = logger.Error
//		}
//	}
//	return nil
//}
//func (this *Datasource) Info(ctx context.Context, msg string, data ...interface{}) {
//	//Log.Infof(msg, append([]interface{}{utils.FileWithLineNum()}, data...)...)
//}
//func (this *Datasource) Warn(ctx context.Context, msg string, data ...interface{}) {
//	//Log.Warnf(msg, append([]interface{}{utils.FileWithLineNum()}, data...)...)
//}
//func (this *Datasource) Error(ctx context.Context, msg string, data ...interface{}) {
//	//Log.Errorf(msg, append([]interface{}{utils.FileWithLineNum()}, data...)...)
//}
//func (this *Datasource) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
//	//elapsed := time.Since(begin)
//	//var rowsStr string
//	//sql, rows := fc()
//	//if -1 == rows {
//	//	rowsStr = "-"
//	//} else {
//	//	rowsStr = strconv.FormatInt(rows, 10)
//	//}
//	//if err != nil {
//	//	Log.Debugf("times %v rows=%v \n ===>>> error = %v \n ===>>> sql = %v", elapsed, rowsStr, err.Error(), sql)
//	//}
//
//}
