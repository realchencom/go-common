package realsmile

import (
	"testing"
)

func TestWritePidFile(t *testing.T) {
	WritePidFile("11111")

}

type RuleVO struct {
	ID string `json:"id"`
}
type Rule struct {
	ID int64 `json:"id"`
}

func TestClone(t *testing.T) {
	n1 := RuleVO{}
	//n1 := RuleVO{}
	n2 := Rule{
		ID: 160345903935254528,
	}

	Clone(&n1,
		n2)
}
