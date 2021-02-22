package libs

import (
	"reflect"
	"testing"
)

func Test_reverse(t *testing.T) {
	type args struct {
		x interface{}
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{name: "Reverse string", args: args{x: "wantToReverse1234"}, want: "4321esreveRoTtnaw"},
		{name: "Reverse []string", args: args{x: []string{
			"elementOne", "elementTwo", "elementThree",
		}}, want: []string{"elementThree", "elementTwo", "elementOne"}},
		{name: "Reverse []int", args: args{x: []int{6, 5, 4, 3, 2, 1}},
			want: []int{1, 2, 3, 4, 5, 6}},
		{name: "Reverse []float64", args: args{x: []float64{6.0, 5.0, 4.0, 3.0, 2.0, 1.0}},
			want: []float64{1.0, 2.0, 3.0, 4.0, 5.0, 6.0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverse(tt.args.x); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_lowerBound(t *testing.T) {
	type args struct {
		x   interface{}
		key interface{}
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "[]int", args: args{
			x:   []int{1, 3, 4, 4, 6, 9},
			key: 4,
		}, want: 2},
		{name: "[]float64", args: args{
			x:   []float64{3.0, 4.0, 4.0, 6.0, 6.0, 9.0},
			key: 6.0,
		}, want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lowerBound(tt.args.x, tt.args.key); got != tt.want {
				t.Errorf("lowerBound() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_upperBound(t *testing.T) {
	type args struct {
		x   interface{}
		key interface{}
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "[]int", args: args{
			x:   []int{1, 3, 4, 4, 6, 9},
			key: 4,
		}, want: 4},
		{name: "[]float64", args: args{
			x:   []float64{3.0, 4.0, 4.0, 6.0, 6.0, 9.0},
			key: 6.0,
		}, want: 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := upperBound(tt.args.x, tt.args.key); got != tt.want {
				t.Errorf("upperBound() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_count(t *testing.T) {
	type args struct {
		x   interface{}
		key interface{}
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "String", args: args{x: "ababaac", key: "a"}, want: 4},
		{name: "[]string", args: args{
			x:   []string{"apple", "banana", "orange", "apple"},
			key: "apple",
		}, want: 2},
		{name: "[]int", args: args{
			x:   []int{1, -1, 2, 0, 1, -1, 3},
			key: -1,
		}, want: 2},
		{name: "[]float64", args: args{
			x:   []float64{1.0, -1.0, 2.0, 0.0, 1.0, -1.0, 3.0},
			key: 1.0,
		}, want: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := count(tt.args.x, tt.args.key); got != tt.want {
				t.Errorf("count() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_find(t *testing.T) {
	type args struct {
		x   interface{}
		key interface{}
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "String", args: args{x: "ababaac", key: "b"}, want: 1},
		{name: "[]string", args: args{
			x:   []string{"apple", "banana", "orange", "apple"},
			key: "orange",
		}, want: 2},
		{name: "[]int", args: args{
			x:   []int{1, -1, 2, 0, 1, -1, 3},
			key: -1,
		}, want: 1},
		{name: "[]float64", args: args{
			x:   []float64{1.0, -1.0, 2.0, 0.0, 1.0, -1.0, 3.0},
			key: 0.0,
		}, want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := find(tt.args.x, tt.args.key); got != tt.want {
				t.Errorf("find() = %v, want %v", got, tt.want)
			}
		})
	}
}
