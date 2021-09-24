package main

import "testing"

func Test_base7trans(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"test1", args{20}, "26"},
		{"test2", args{-7}, "-10"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := base7trans(tt.args.num); got != tt.want {
				t.Errorf("base7trans() = %v, want %v", got, tt.want)
			}
		})
	}
}
