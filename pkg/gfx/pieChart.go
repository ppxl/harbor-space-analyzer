package gfx

import "math"

// CharFactor determines how often a character will be repeated and thus how broad the pie will be skewed.
// Usually, only values of 1 or 2 look okay.
var CharFactor = 2

// PrintChart prints a lovely pie chart.
func PrintChart(karacters []string, values []float64, r int) string {
	// this solution is heavily based on https://codegolf.stackexchange.com/a/23351/119676 so don't be mad
	resultPie := ""
	d := makeRange(-r, r)

	for yIdx, _ := range d {
		currentLine := ""
		y := d[yIdx]

		for xIdx, _ := range d {
			x := d[xIdx]
			if x*x+y*y < r*r {
				a := math.Atan2(float64(y), float64(x))/math.Pi/2 + .5
				nextChar := s(karacters, values, a)
				currentLine = currentLine + charTimes(nextChar, CharFactor)
			} else {
				currentLine = currentLine + charTimes(" ", CharFactor)
			}

		}
		resultPie += currentLine + "\n"
	}

	return resultPie
}

func s(karacters []string, values []float64, a float64) string {
	if len(values) == 0 {
		return "-"
	}
	if a < values[0] {
		return karacters[0]
	}
	return s(karacters[1:], values[1:], a-values[0])
}

// returns the same character n-times where factor is n.
func charTimes(s2 string, factor int) string {
	result := ""
	for i := 0; i < factor; i++ {
		result += s2
	}
	return result
}

func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}
