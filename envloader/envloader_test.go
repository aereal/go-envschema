package envloader

import (
	"reflect"
	"testing"
)

func Test_environAsMap(t *testing.T) {
	type args struct {
		pairs []string
	}
	tests := []struct {
		name string
		args args
		want map[string]string
	}{
		{
			name: "no pairs",
			args: args{
				pairs: []string{},
			},
			want: map[string]string{},
		},
		{
			name: "some pairs (no duplicates)",
			args: args{
				pairs: []string{"FOO=1", "BAR=2"},
			},
			want: map[string]string{"FOO": "1", "BAR": "2"},
		},
		{
			name: "some pairs w/duplicates",
			args: args{
				pairs: []string{"FOO=1", "BAR=2", "FOO=3"},
			},
			want: map[string]string{"FOO": "3", "BAR": "2"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := environAsMap(tt.args.pairs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("environAsMap() = %v, want %v", got, tt.want)
			}
		})
	}
}
