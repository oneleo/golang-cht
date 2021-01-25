package c01

import (
	"fmt"
	"testing"
)

func TestMultiMath(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name      string
		args      args
		wantAdd   int
		wantMinus int
	}{
		// TODO: Add test cases.
		{
			name: "a = 1; b = 1",
			args: args{
				a: 1,
				b: 1,
			},
			wantAdd:   2,
			wantMinus: 0,
		},
		{
			name: "a = 2; b = 1",
			args: args{
				a: 2,
				b: 1,
			},
			wantAdd:   3,
			wantMinus: 1,
		},
		{
			name: "a = 1; b = 2",
			args: args{
				a: 1,
				b: 2,
			},
			wantAdd:   3,
			wantMinus: -1,
		},
		{
			name: "a = 2; b = 2",
			args: args{
				a: 2,
				b: 2,
			},
			wantAdd:   4,
			wantMinus: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotAdd, gotMinus := MultiMath(tt.args.a, tt.args.b)
			if gotAdd != tt.wantAdd {
				t.Errorf("MultiMath() gotAdd = %v, want %v", gotAdd, tt.wantAdd)
			}
			if gotMinus != tt.wantMinus {
				t.Errorf("MultiMath() gotMinus = %v, want %v", gotMinus, tt.wantMinus)
			}
		})
	}
}

func ExampleMultiMath() {
	a := 1
	b := 1
	out1, out2 := MultiMath(a, b)
	fmt.Println(out1, out2)
	// Output:
	// 2 0
}
