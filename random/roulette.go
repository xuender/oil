package random

import (
	"math/rand"
	"sort"

	"github.com/xuender/oil/integer"
)

// Roulette 轮盘
type Roulette struct {
	Scores []Scorer
	list   [][2]int
	max    int
}

// Add 增加
func (r *Roulette) Add(score ...Scorer) {
	if len(score) == 0 {
		return
	}
	r.Scores = append(r.Scores, score...)
}

// init 初始化轮盘
func (r *Roulette) init() {
	if len(r.Scores) == 0 {
		return
	}
	if r.list == nil || len(r.list) != len(r.Scores) {
		// 轮盘初始化
		r.list = make([][2]int, len(r.Scores))
		// 初始化坐标, 并寻找最小积分
		// list[i][0]:积分, list[i][1]:序号
		min := integer.MaxInt
		for i, s := range r.Scores {
			r.list[i][1] = i
			a := s.Score()
			r.list[i][0] = a
			if min > a {
				min = a
			}
		}
		// 将所有积分设置成大于1的正数
		if min < 1 {
			abs := -1*min + 1
			for i := range r.list {
				r.list[i][0] += abs
			}
		}
		// 按照积分排序
		sort.Slice(r.list, func(i, j int) bool {
			return r.list[i][0] < r.list[j][0]
		})
		// 后面的值进行偏移,扩大面积
		add := 0
		for i := range r.list {
			r.list[i][0] += add
			add = r.list[i][0]
		}
		// 最大值
		r.max = r.list[len(r.list)-1][0]
	}
}

// Take 旋转轮盘按照积分概率获取
func (r *Roulette) Take() Scorer {
	r.init()
	s := rand.Intn(r.max)
	for _, l := range r.list {
		if s < l[0] {
			return r.Scores[l[1]]
		}
	}
	return r.Scores[0]
}

// NewRoulette 新建轮盘
func NewRoulette(scores []Scorer) *Roulette {
	ret := &Roulette{Scores: scores}
	ret.init()
	return ret
}
