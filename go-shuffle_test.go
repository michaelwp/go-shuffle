package go_shuffle

import (
	"reflect"
	"testing"
)

func TestListTimes_Shuffle(t *testing.T) {
	type fields struct {
		List  List
		Times int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// test cases.
		{
			name: "ListTimes Shuffle",
			fields: fields{
				List:  List{1, 2, 3, 4, 5},
				Times: 5,
			},
			want: 5,
		},
		{
			name: "ListTimes Shuffle",
			fields: fields{
				List:  List{"a", "b", "c", "d", "e"},
				Times: 3,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := ListTimes{
				List:  tt.fields.List,
				Times: tt.fields.Times,
			}
			if got := len(l.Shuffle()); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Shuffle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestList_Shuffle(t *testing.T) {
	tests := []struct {
		name string
		l    List
		want List
	}{
		// test cases.
		{
			name: "List shuffle",
			l:    List{1, 2, 3, 4, 5},
			want: List{1, 2, 3, 4, 5},
		},
		{
			name: "List shuffle",
			l:    List{"a", "b", "c", "d", "e"},
			want: List{"a", "b", "c", "d", "e"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.Shuffle(); reflect.DeepEqual(got, tt.want) {
				t.Errorf("Shuffle() = %v, want %v", got, tt.want)
			}
		})
	}
}
