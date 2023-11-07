package gfx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_charTimes(t *testing.T) {
	type args struct {
		character string
		factor    int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"return original", args{character: "A", factor: 0}, ""},
		{"return original", args{character: "A", factor: 1}, "A"},
		{"return original", args{character: "A", factor: 2}, "AA"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, charTimes(tt.args.character, tt.args.factor), "charTimes(%v, %v)", tt.args.character, tt.args.factor)
		})
	}
}

func Test_makeRange(t *testing.T) {
	type args struct {
		min int
		max int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"zero all the way down", args{0, 0}, []int{0}},
		{"2 elements starting at 0", args{0, 1}, []int{0, 1}},
		{"5 elements starting at 0", args{0, 4}, []int{0, 1, 2, 3, 4}},
		{"2 elements starting at 1", args{1, 2}, []int{1, 2}},
		{"5 elements starting at 1", args{1, 5}, []int{1, 2, 3, 4, 5}},
		{"2 elements starting at -1", args{-1, 0}, []int{-1, 0}},
		{"5 elements starting at -1", args{-1, 3}, []int{-1, 0, 1, 2, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, makeRange(tt.args.min, tt.args.max), "makeRange(%v, %v)", tt.args.min, tt.args.max)
		})
	}

	t.Run("should panic at max < min", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("The code did not panic")
			}
		}()
		makeRange(90, 0)
	})
}
