package chart

import (
	"fmt"
	"strings"

	"github.com/xuender/oil/integer"
)

// 模板
const _svg = `<?xml version="1.0" standalone="no"?><!DOCTYPE svg PUBLIC "-//W3C//DTD SVG 1.1//EN" "http://www.w3.org/Graphics/SVG/1.1/DTD/svg11.dtd"><svg version="1.1" width="%d" height="%d" xmlns="http://www.w3.org/2000/svg"><rect height="100%%" width="100%%" fill="#f3f3f3"/>%s%s%s%s</svg>`

// SVG预估尺寸
const _size = 800

// 默认边距
const _margin = 40

// Polyline 折线图
func Polyline(data []int, style ...int) string {
	if len(data) < 2 {
		return fmt.Sprintf(_svg, _size, _size, "", "", "", "")
	}
	min := integer.MaxInt
	max := integer.MinInt
	for _, d := range data {
		if min > d {
			min = d
		}
		if max < d {
			max = d
		}
	}
	// 数据宽度
	dataWidth := len(data) - 1
	// 数据高度
	dataHeight := max - min
	bw := integer.Div(_size, dataWidth)
	if bw == 0 {
		bw = 1
	}
	bh := integer.Div(_size, dataHeight)
	width := bw*dataWidth + _margin*3/2
	height := bh*dataHeight + _margin

	// 数据
	points := make([]string, len(data))
	for i, d := range data {
		points[i] = fmt.Sprintf("%d,%d", i, (d-max)*-1)
	}
	dataArea := fmt.Sprintf(`<g transform="translate(40,20) scale(%d,%d)" fill="none" stroke="chocolate" stroke-width="1"><polyline points="%s"/></g>`, bw, bh, strings.Join(points, " "))
	// 零线
	zero := ""
	if min < 0 && max > 0 {
		zero = fmt.Sprintf(`<line x1="%d" y1="%d" x2="100%%" y2="%d" stroke="red"/>`, _margin, (dataHeight+min)*bh+_margin/2, (dataHeight+min)*bh+_margin/2)
	}
	// 坐标轴
	baseAxis := fmt.Sprintf(`<g fill="none" stroke="black" stroke-width="1"><line x1="0" y1="%d" x2="100%%" y2="%d"/><line x1="40" y1="0" x2="40" y2="100%%"/>%s</g>`, height-20, height-20, zero)

	// 垂线
	vertical := make([]line, 5)
	for i := 1; i < 5; i++ {
		vertical[i-1] = line{x: integer.Div(bw*dataWidth/5*i, 10)*10 + _margin}
	}
	// 最后一条
	vertical[4] = line{x: width - _margin/2}
	// 垂线
	vl := ""
	for _, v := range vertical {
		vl += fmt.Sprintf(`<line x1="%d" y1="0" x2="%d" y2="%d"/>`, v.x, v.x, height-_margin/2)
	}

	// 水平线
	horizontal := make([]line, 5)
	for i := 0; i < 5; i++ {
		horizontal[i] = line{y: integer.Div(dataHeight/5*i, 10)*10*bh + _margin/2}
	}
	horizontal[0] = line{y: _margin / 2}
	hl := ""
	for _, h := range horizontal {
		hl += fmt.Sprintf(`<line x1="%d" y1="%d" x2="100%%" y2="%d"/>`, _margin, h.y, h.y)
	}
	scaleAxis := fmt.Sprintf(`<g stroke="grey" stroke-dasharray="6" stroke-width="0.5">%s%s</g>`, hl, vl)

	// 文字
	texts := make([]text, 12)
	for i := 0; i < 5; i++ {
		texts[i] = text{
			v:      integer.Div((len(data)/5)*(i+1), 10) * 10,
			x:      vertical[i].x,
			y:      height - 6,
			anchor: "middle",
		}
	}
	texts[4].v = len(data)
	for i := 0; i < 5; i++ {
		texts[i+5] = text{
			v:      integer.Div((dataHeight/5)*(5-i)+min, 10) * 10,
			x:      _margin - 5,
			y:      horizontal[i].y + 5,
			anchor: "end",
		}
	}
	// 最大值
	texts[5] = text{v: max, anchor: "end", x: _margin - 5, y: _margin/2 + 5}
	// 最小值
	texts[10] = text{v: min, anchor: "end", x: _margin - 5, y: height - _margin/2 - 3}
	// 零线
	if min < 0 && max > 0 {
		texts[11] = text{v: 0, anchor: "end", x: _margin - 5, y: (dataHeight+min)*bh + _margin/2 + 5}
	}
	vt := ""
	for _, t := range texts {
		vt += fmt.Sprintf(`<text x="%d" y="%d" text-anchor="%s">%d</text>`, t.x, t.y, t.anchor, t.v)
	}
	textAxis := fmt.Sprintf(`<g fill="grey" stroke="transparent">%s</g>`, vt)
	return fmt.Sprintf(_svg, width, height, scaleAxis, baseAxis, dataArea, textAxis)
}
