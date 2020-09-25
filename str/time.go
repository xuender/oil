package str

import "fmt"

const minute = int64(60000000000)
const hour = minute * 60
const day = hour * 24

// Time 纳秒转换成时间
func Time(ns int64) string {
	if ns >= int64(100000000) {
		// 超过0.1秒以后才考虑分钟、小时、天
		if ns >= day {
			d := ns % day
			if d > 0 {
				return fmt.Sprintf("%d天%s", ns/day, Time(d))
			}
			return fmt.Sprintf("%d天", ns/day)
		}
		if ns >= hour {
			m := ns % hour
			if m > 0 {
				return fmt.Sprintf("%d小时%s", ns/hour, Time(m))
			}
			return fmt.Sprintf("%d小时", ns/hour)
		}
		if ns >= minute {
			s := ns % minute
			if s > 0 {
				return fmt.Sprintf("%d分钟%s", ns/minute, Time(s))
			}
			return fmt.Sprintf("%d分钟", ns/minute)
		}
		return fmt.Sprintf("%.3v秒", float64(ns)/1000000000)
	}
	if ns >= int64(100000) {
		return fmt.Sprintf("%.3v毫秒", float64(ns)/1000000)
	}
	if ns >= int64(100) {
		return fmt.Sprintf("%.3v微妙", float64(ns)/1000)
	}
	return fmt.Sprintf("%d纳秒", ns)
}
