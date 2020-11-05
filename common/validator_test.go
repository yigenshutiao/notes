package common

import (
	"fmt"
	"testing"
)

func Test_Validator(t *testing.T) {
	fmt.Println(Validate.Var(11, "gte=0")) // no error
	if err := Validate.Var("", "required"); err != nil {
		fmt.Println(err) // error
	}
}
