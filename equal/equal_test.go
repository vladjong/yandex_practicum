package equal

import "testing"

func TestEqual(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "ok_false",
			args: args{a: 10, b: 5},
			want: false,
		},
		{
			name: "pk_true",
			args: args{a: 10, b: 10},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Equal(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Equal() = %v, want %v", got, tt.want)
			}
		})
	}
}
