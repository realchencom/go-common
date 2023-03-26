package realsmile

import (
	"testing"
)

func TestWritePidFile(t *testing.T) {
	WritePidFile("11111")

}

type RuleVO struct {
	ID  string `json:"id"`
	Age int64  `json:"age"`
}
type Rule struct {
	Base
	Age int64 `json:"age"`
}

func TestClone(t *testing.T) {
	news := &RuleVO{}
	if news == nil {
		panic("new is nil")
		return
	}

	//n1 := RuleVO{}
	//n1 := RuleVO{
	//	ID:  "160345903935254528",
	//	Age: 30,
	//}
	//n2 := Rule{}
	//n2 := Rule{
	//	Base: Base{
	//		ID: 160345903935254528,
	//	},
	//
	//	Age: 32,
	//}
	//marshal, err := json.Marshal(n2)
	//if err != nil {
	//	return
	//}
	//n3 := Rule{}
	//func(marshal []byte, v interface{}) {
	//	if err := json.Unmarshal(marshal, v); err != nil {
	//		Log.Debugf("%v", err)
	//	}
	//}(marshal, &n3)
	//
	//Log.Debugf("%v", marshal)
	//Clone(&n1, n2)
	//val := reflect.ValueOf(&n2)
	//v1 := reflect.ValueOf(n1)
	//Clone(&n2, n1)
	//Log.Debug("d")
	size := Combination.CombinationSize(60, 2)
	Log.Debug(size)
	combination, err := Combination.GetCombination(60, 2)
	if err != nil {
		return
	}
	Log.Debug(combination)
}
