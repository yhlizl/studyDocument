package assignment

import (
	"reflect"
	"testing"
)

func Test_q1FindDuplicated(t *testing.T) {
	type args struct {
		n []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"test1", args{[]int{1, 3, 3, 3, 2}}, 3},
		{"test2", args{[]int{1, 2, 3, 4, 4}}, 4},
		{"test3", args{[]int{1, 2, 4, 3, 2}}, 2},
		{"test4", args{[]int{1, 1}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := q1FindDuplicated(tt.args.n); got != tt.want {
				t.Errorf("q1FindDuplicated() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_floodFill(t *testing.T) {
	type args struct {
		image    [][]int
		sr       int
		sc       int
		newColor int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		// TODO: Add test cases.
		{"test1", args{[][]int{{1, 1, 1}, {1, 1, 0}, {1, 0, 1}}, 1, 1, 2}, [][]int{{2, 2, 2}, {2, 2, 0}, {2, 0, 1}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := floodFill(tt.args.image, tt.args.sr, tt.args.sc, tt.args.newColor); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("floodFill() = %v, want %v", got, tt.want)
			}
		})
	}
}
