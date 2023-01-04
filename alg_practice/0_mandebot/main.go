package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

// 是否绘制
func inMandelbrotZone(x0 float64, y0 float64, n int) bool {
	var (
		x float64 = 0.0
		y float64 = 0.0
	)
	for n > 0 {
		xtemp := x*x - y*y + x0
		y = 2.0*x*y + y0
		x = xtemp
		n = n - 1
		if x*x+y*y > 4.0 {
			return false
		}
	}
	return true
}

// 计算代码
func DoPlot() ([]string, []string) {

	var (
		xs         float64 = -2.0
		xe         float64 = 2.0
		ys         float64 = -1.5
		ye         float64 = 1.5
		width      float64 = 80.0
		height     float64 = 40.0
		threshhold float64 = 10000.0
	)
	endStr := []string{}
	endStrTwo := []string{}
	for i := 0; i < 2; i++ {

		dx := (xe - xs) / width
		dy := (ye - ys) / height
		y := ye

		n := 0
		for y >= ys {
			x := xs
			lineStr := []string{}
			for x < xe {
				n += 1
				if inMandelbrotZone(x, y, int(threshhold)) {
					lineStr = append(lineStr, "*")
				} else {
					lineStr = append(lineStr, ".")

				}
				x += dx
			}
			if i == 0 {
				lines := fmt.Sprintf("%v", strings.Join(lineStr, ""))
				endStr = append(endStr, lines)

			} else {
				stwo := ChangePlace(lineStr)
				lines := fmt.Sprintf("%v", strings.Join(stwo, ""))
				endStrTwo = append(endStrTwo, lines)

			}
			y -= dy
		}
	}

	return endStr, endStrTwo
}

// 存图
func MergeP() {

	a, a2 := DoPlot()
	var a3 = []string{}
	start := time.Now()
	aa2 := a2[:]

	for i, v := range a {
		a3 = append(a3, fmt.Sprintf("%v%v", v, aa2[i]))
	}

	fmt.Println(time.Since(start))

	var tstr string
	for _, k := range a3 {
		fmt.Println(k)
		tstr += k + "\n"
	}
	os.WriteFile("./md.txt", []byte(tstr), 0772)

}

func ChangePlace(st []string) []string {

	reverSlice := []string{}
	for i, _ := range st {
		reverSlice = append(reverSlice, st[len(st)-(1+i)])
	}
	return reverSlice
}

func main() {
	MergeP()
}
