package go_shuffle

import (
	"reflect"
	"testing"
)

func TestList_Start(t *testing.T) {
	type args struct {
		t Times
	}
	tests := []struct {
		name string
		l    List
		args args
		want int
	}{
		// test cases.
		{
			name: "start",
			l:    []interface{}{1, 2, 3, 4, 5},
			args: args{t: 5},
			want: 5,
		},
		{
			name: "start",
			l:    []interface{}{"a", "b", "c", "d", "e", "f"},
			args: args{t: 6},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := len(tt.l.Start(tt.args.t)); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Start() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestList_shuffle(t *testing.T) {
	tests := []struct {
		name string
		l    List
		want List
	}{
		// test cases.
		{
			name: "shuffle",
			l:    []interface{}{1, 2, 3, 4, 5},
			want: []interface{}{1, 2, 3, 4, 5},
		},
		{
			name: "shuffle",
			l:    []interface{}{"a", "b", "c", "d", "e"},
			want: []interface{}{"a", "b", "c", "d", "e"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.shuffle(); reflect.DeepEqual(got, tt.want) {
				t.Errorf("shuffle() = %v, want %v", got, tt.want)
			}
		})
	}
}
