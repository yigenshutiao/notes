package common

import (
	"fmt"
	"testing"
)

func Test_form2JSON(t *testing.T) {
	type args struct {
		r *HTTPResponse
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"test 1",
			args{r: &HTTPResponse{
				Msg:  "hello",
				Data: nil,
			}},
			"{\"errmsg\":\"hello\",\"data\": \"\"}",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := form2JSON(tt.args.r); got != tt.want {
				fmt.Println(got)
				t.Errorf("form2JSON() = %v, want %v", got, tt.want)
			}
		})
	}
}
