package logic

import (
	"fmt"
	"testing"
)

func Test_genID(t *testing.T) {
	type args struct {
		content string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"test 1",
			args{content: "hello"},
			"should different",
		},
		{
			"test 1",
			args{content: "hello"},
			"should different",
		},
	}
	for _, tt := range tests {
		fmt.Println(genID(tt.args.content))
	}
}
