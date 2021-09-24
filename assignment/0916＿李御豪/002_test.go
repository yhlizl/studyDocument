package main

import (
	"reflect"
	"testing"
)

func Test_twosum(t *testing.T) {
	type args struct {
		arr    []int
		target int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		// TODO: Add test cases.
		{"test1", args{[]int{1, 2, 3, 4}, 4}, [][]int{{1, 2}}},
		{"test2", args{[]int{1, 2, 2, 3, 4}, 5}, [][]int{{1, 3}, {2, 2}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twosum(tt.args.arr, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twosum() = %v, want %v", got, tt.want)
			}
		})
	}
}
