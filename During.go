package timing

import (
	"time"
)

type T_time struct {
	start  string
	end    string
	during float64
}

func init() {
	var cstZone = time.FixedZone("CST", 8*3600) // 东八
	time.Local = cstZone
}
func MissionDuring(start, end time.Time) *T_time {
	t := &T_time{
		start:  start.Format("15:04:05"),
		end:    end.Format("15:04:05"),
		during: end.Sub(start).Minutes(),
	}
	return t
}
