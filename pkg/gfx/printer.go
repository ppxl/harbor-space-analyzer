package gfx

import (
	"fmt"
	c "github.com/fatih/color"
	"github.com/ppxl/harbor-space-analyzer/pkg/core"
	"os"
)

var colorList = []c.Attribute{
	c.FgRed,
	c.FgGreen,
	c.FgYellow,
	c.FgBlue,
	c.FgMagenta,
	c.FgCyan,
	c.FgHiRed,
	c.FgHiGreen,
	c.FgHiYellow,
	c.FgHiBlue,
	c.FgHiMagenta,
	c.FgHiCyan,
}

var osPrinter = os.Stdout

// PrintColorChart prints the given chart in colors based on the given karacters.
func PrintColorChart(karacters []string, output string) {
	colors := buildColorMap(karacters)

	for _, kar := range []byte(output) {
		foundColor, ok := colors[kar]
		karToPrint := string(kar)
		if ok {
			karToPrint = foundColor.Sprint(karToPrint)
		}

		_, _ = fmt.Fprint(osPrinter, karToPrint)
	}
}

func buildColorMap(karacters []string) map[byte]*c.Color {
	result := map[byte]*c.Color{}

	for idx, kar := range karacters {
		colorIdx := idx % len(colorList)
		chosenColor := colorList[colorIdx]
		result[toByte(kar)] = c.New(chosenColor)
	}

	return result
}

// PrintLegend prints a colorized legend.
func PrintLegend(summaries []core.ProjectSummary, karacters []string, percent []float64) {
	colors := buildColorMap(karacters)

	for idx, kar := range karacters {
		foundColor, ok := colors[toByte(kar)]
		karToPrint := kar
		if ok {
			karToPrint = foundColor.Sprint(karToPrint)
		}

		fmt.Printf("%s: %5.2f - %s\n", karToPrint, percent[idx]*100.0, summaries[idx].ProjectName)
	}
}

func toByte(s string) byte {
	if len(s) != 1 {
		panic(fmt.Sprintf("could not handle string (expected 1, got %d)", len(s)))
	}

	return []byte(s)[0]
}
