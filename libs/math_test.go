package libs

import (
	"reflect"
	"testing"
)

func Test_abs(t *testing.T) {
	type args struct {
		a int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "Plus", args: args{3}, want: 3},
		{name: "Minus", args: args{-3}, want: 3},
		{name: "Zero", args: args{0}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := abs(tt.args.a); got != tt.want {
				t.Errorf("abs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_max(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "Plus", args: args{a: 1, b: 3}, want: 3},
		{name: "Minus", args: args{a: -1, b: -3}, want: -1},
		{name: "Different", args: args{a: -1, b: 3}, want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := max(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("max() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_min(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "Plus", args: args{a: 1, b: 3}, want: 1},
		{name: "Minus", args: args{a: -1, b: -3}, want: -3},
		{name: "Different", args: args{a: -1, b: 3}, want: -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := min(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("min() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_gcd(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "GCD", args: args{a: 4, b: 6}, want: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := gcd(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("gcd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_nextPermutation(t *testing.T) {
	x := []int{1, 2, 3, 4}
	wants := [][]int{
		{1, 2, 3, 4},
		{1, 2, 4, 3},
		{1, 3, 2, 4},
		{1, 3, 4, 2},
		{1, 4, 2, 3},
		{1, 4, 3, 2},
		{2, 1, 3, 4},
		{2, 1, 4, 3},
		{2, 3, 1, 4},
		{2, 3, 4, 1},
		{2, 4, 1, 3},
		{2, 4, 3, 1},
		{3, 1, 2, 4},
		{3, 1, 4, 2},
		{3, 2, 1, 4},
		{3, 2, 4, 1},
		{3, 4, 1, 2},
		{3, 4, 2, 1},
		{4, 1, 2, 3},
		{4, 1, 3, 2},
		{4, 2, 1, 3},
		{4, 2, 3, 1},
		{4, 3, 1, 2},
		{4, 3, 2, 1},
	}
	for i := 0; ; i++ {
		if !reflect.DeepEqual(wants[i], x) {
			t.Errorf("%v expected, but %v\n", wants[i], x)
		}
		if !nextPermutation(x) {
			break
		}
	}
}

func Test_sum(t *testing.T) {
	type args struct {
		a []int
	}
	tests := []struct {
		name   string
		args   args
		wantSu int
	}{
		{name: "Sum", args: args{a: []int{1, 2, 3}}, wantSu: 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotSu := sum(tt.args.a); gotSu != tt.wantSu {
				t.Errorf("sum() = %v, want %v", gotSu, tt.wantSu)
			}
		})
	}
}
