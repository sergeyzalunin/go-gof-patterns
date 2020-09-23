package helper

import (
	"fmt"
	"testing"
)

func TestGetPrintOutput(t *testing.T) {
	type args struct {
		testingFunc func()
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "checking the fmt.Print",
			args: args{
				testingFunc: func() {
					fmt.Print("Hello World!!!")
				},
			},
			want: "Hello World!!!",
		},
		{
			name: "checking the fmt.Println",
			args: args{
				testingFunc: func() {
					fmt.Println("Hello World!!!")
				},
			},
			want: "Hello World!!!\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetPrintOutput(tt.args.testingFunc); got != tt.want {
				t.Errorf("GetPrintOutput() = %v, want %v", got, tt.want)
			}
		})
	}
}
