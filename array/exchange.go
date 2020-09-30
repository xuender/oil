package array

// Exchange 交换组合
// 假设有手机颜色、内存、运营商不同属性要求组合出所有结果
// 红8G移动，红8G联通,红16G移动,红16G联通,蓝8G移动...
type Exchange struct {
	curr []int
	data []int
	init bool
}

// Next 是否有下一个
func (e *Exchange) Next() bool {
	if e.init {
		e.init = false
		return true
	}
	return e.add(len(e.data) - 1)
}

// 加法
func (e *Exchange) add(num int) bool {
	if num < 0 || num >= len(e.data) {
		return false
	}
	if e.curr[num] < e.data[num]-1 {
		e.curr[num]++
		return true
	}
	if e.add(num - 1) {
		e.curr[num] = 0
		return true
	}
	return false
}

// Value 取值
func (e *Exchange) Value() []int {
	return e.curr
}

// NewExchange 创建交换组合
func NewExchange(data []int) *Exchange {
	curr := make([]int, len(data))
	return &Exchange{
		data: data,
		curr: curr,
		init: true,
	}
}
