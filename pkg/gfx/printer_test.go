package gfx

import (
	c "github.com/fatih/color"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"os"
	"testing"
)

var testByteCharA = []byte("A")[0]

func Test_buildColorMap(t *testing.T) {
	t.Run("should return nothing", func(t *testing.T) {
		actual := buildColorMap([]string{})

		assert.Empty(t, actual)
	})
	t.Run("should return first color", func(t *testing.T) {
		actual := buildColorMap([]string{"A"})

		assert.Len(t, actual, 1)
		assert.Equal(t, actual[testByteCharA], c.New(c.FgRed))
	})
	t.Run("should return first color after running out of colors in the first round", func(t *testing.T) {
		actual := buildColorMap([]string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M"})

		require.Len(t, actual, 13)
		assert.Equal(t, actual[testByteCharA], c.New(c.FgRed))
		assert.Equal(t, actual[[]byte("B")[0]], c.New(c.FgGreen))
		assert.Equal(t, actual[[]byte("C")[0]], c.New(c.FgYellow))
		assert.Equal(t, actual[[]byte("D")[0]], c.New(c.FgBlue))
		assert.Equal(t, actual[[]byte("E")[0]], c.New(c.FgMagenta))
		assert.Equal(t, actual[[]byte("F")[0]], c.New(c.FgCyan))
		assert.Equal(t, actual[[]byte("G")[0]], c.New(c.FgHiRed))
		assert.Equal(t, actual[[]byte("H")[0]], c.New(c.FgHiGreen))
		assert.Equal(t, actual[[]byte("I")[0]], c.New(c.FgHiYellow))
		assert.Equal(t, actual[[]byte("J")[0]], c.New(c.FgHiBlue))
		assert.Equal(t, actual[[]byte("K")[0]], c.New(c.FgHiMagenta))
		assert.Equal(t, actual[[]byte("L")[0]], c.New(c.FgHiCyan))
		assert.Equal(t, actual[[]byte("M")[0]], c.New(c.FgRed)) // here we start the color list anew
	})
}

func TestPrintColorChart(t *testing.T) {
	t.Run("should print the character A in red", func(t *testing.T) {
		c.NoColor = false // tests are noColor by default but color is part of the whole thing

		// given
		rescueStdout := osPrinter
		defer func() { osPrinter = rescueStdout }()
		r, w, _ := os.Pipe()
		osPrinter = w

		// when
		PrintColorChart([]string{"A"}, "   AAA   ")

		// then
		_ = w.Close()
		out, _ := io.ReadAll(r)
		assert.Equal(t, "   \u001B[31mA\u001B[0m\u001B[31mA\u001B[0m\u001B[31mA\u001B[0m   ", string(out))
	})
}

func Test_toByte(t *testing.T) {
	t.Run("should panic on empty", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("The code did not panic")
			}
		}()
		toByte("")
	})
	t.Run("should panic on multibyte string", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("The code did not panic")
			}
		}()
		toByte("12")
	})
	t.Run("should return byte value", func(t *testing.T) {
		actual := toByte(" ")
		assert.Equal(t, " ", string(actual))
	})
}
