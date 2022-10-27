package realsmile

import (
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	ti := Time.SnowFlakeStartTime(2022, time.January, 1, 0, 0, 0, 0)
	Log.Infof("%v", ti)

}
