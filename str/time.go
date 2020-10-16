package str

import "fmt"

const (
	microsecond = int64(1000)
	millisecond = microsecond * 1000
	second      = millisecond * 1000
	minute      = second * 60
	hour        = minute * 60
	day         = hour * 24
)

// Time 纳秒转换成时间
func Time(ns int64) string {
	if ns >= second {
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

	if ns >= millisecond {
		return fmt.Sprintf("%.3v毫秒", float64(ns)/1000000)
	}

	if ns >= microsecond {
		return fmt.Sprintf("%.3v微妙", float64(ns)/1000)
	}

	return fmt.Sprintf("%d纳秒", ns)
}
