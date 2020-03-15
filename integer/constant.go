package integer

// MaxInt 最大整数
// 64 位等于 math.MaxUint64
// 32 位等于 math.MaxUint32
const MaxInt = int(^uint(0) >> 1)

// MinInt 最小整数
// 64 位等于 math.MinUint64
// 32 位等于 math.MinUint32
const MinInt = ^MaxInt
