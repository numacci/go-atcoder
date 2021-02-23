package libs

import (
	"reflect"
	"testing"
)

func Test_intToStr(t *testing.T) {
	type args struct {
		i int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "Plus", args: args{i: 1}, want: "1"},
		{name: "Minus", args: args{i: -1}, want: "-1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := intToStr(tt.args.i); got != tt.want {
				t.Errorf("intToStr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_join(t *testing.T) {
	type args struct {
		x []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "Join", args: args{x: []string{"a", "b", "c"}}, want: "abc"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := join(tt.args.x); got != tt.want {
				t.Errorf("join() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_split(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{name: "Split", args: args{s: "abc"}, want: []string{"a", "b", "c"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := split(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("split() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_strToInt(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "Plus", args: args{s: "1"}, want: 1},
		{name: "Minus", args: args{s: "-1"}, want: -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := strToInt(tt.args.s); got != tt.want {
				t.Errorf("strToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
