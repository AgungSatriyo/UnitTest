package services

import "testing"

func Test_grade(t *testing.T) {
	type args struct {
		skor int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case return C",
			args: args{
				skor: 35,
			},
			want: "C",
		},
		{
			name: "case return C",
			args: args{
				skor: 50,
			},
			want: "B",
		},
		{
			name: "case return C",
			args: args{
				skor: 80,
			},
			want: "A",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := grade(tt.args.skor); got != tt.want {
				t.Errorf("grade() = %v, want %v", got, tt.want)
			}
		})
	}
}
