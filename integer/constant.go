package integer

// MaxInt 最大整数
// 64 位等于 math.MaxInt64
// 32 位等于 math.MaxInt32
const MaxInt = int(^uint(0) >> 1)

// MinInt 最小整数
// 64 位等于 math.MinInt64
// 32 位等于 math.MinInt32
const MinInt = ^MaxInt
