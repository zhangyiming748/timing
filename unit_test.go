package timing

import (
	"github.com/zhangyiming748/pretty"
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	before := time.Now()
	time.Sleep(3 * time.Second)
	after := time.Now()
	ret := MissionDuring(before, after)
	pretty.P(ret)
}
