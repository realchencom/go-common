package realsmile

import "testing"

func TestLog(t *testing.T) {
	Log.Sync()
	Log.Debug("debug test")
	//var ds Datasource
	//db, err := ds.GetMylSqlDB()
	//if err != nil {
	//	return
	//}
	//Log.Debugf("%v", db)
}
