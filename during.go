package timing

import (
	"fmt"
	"time"
)

func init() {
	var cstZone = time.FixedZone("CST", 8*3600) // 东八
	time.Local = cstZone
}

func During(start, end time.Time) string {
	// 获取开始时间
	//start := time.Now()

	// 模拟一些操作，这里可以替换为实际的程序运行或其他操作
	//time.Sleep(15 * time.Second)

	// 获取结束时间
	//end := time.Now()

	// 计算持续时间
	duration := end.Sub(start)

	// 分别计算小时数和分钟数
	hours := int(duration.Hours())
	minutes := int(duration.Minutes()) % 60
	seconds := int(duration.Seconds()) % 60

	// 将分钟数和秒数转换为小数形式的分钟数
	decimalMinutes := float64(minutes) + float64(seconds)/60

	// 按照 xx时x.xxx分 格式输出持续时间
	if hours > 0 {
		return fmt.Sprintf("%d时%.3f分\n", hours, decimalMinutes)
	} else {
		return fmt.Sprintf("%.3f分\n", decimalMinutes)
	}
}
