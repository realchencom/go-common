package realsmile

import "testing"

func TestSnowFlake_NextId(t *testing.T) {
	for i := 0; i < 10; i++ {
		Log.Infof("%v", Snow.NextId())
	}

}
