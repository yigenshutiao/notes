package common

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"reflect"
)

func HTTPHandler(handlefunc interface{}) httprouter.Handle {
	controller := &Controller{handleFunc: handlefunc}

	checkHandleFunc(handlefunc)

	return controller.HandleHTTP
}

func checkHandleFunc(handleFunc interface{}) {
	value := reflect.ValueOf(handleFunc)
	typeOf := reflect.TypeOf(handleFunc)

	if value.Kind() != reflect.Func {
		panic("[checkHandleFunc] handleFunc is not a func")
	}

	if typeOf.NumIn() != 2 {
		panic("[checkHandleFunc] handleFunc need 2 input")
	}

	if typeOf.NumOut() != 2 {
		panic("[checkHandleFunc] handleFunc need 2 output")
	}

	firstParamType := typeOf.In(0)
	if firstParamType.String() != "context.Context" {
		panic("[checkHandleFunc] handleFunc first int param must be [context.Context]")
	}

	secondParamType := typeOf.In(1)
	if secondParamType.Kind() != reflect.Ptr {
		panic("[checkHandleFunc] handleFunc second int param must be [struct's pointer]")
	}

	secondOutParamType := typeOf.Out(1)
	if secondOutParamType.String() != "error" {
		panic("[checkHandleFunc] handleFunc second out param must be [error]")
	}

}

type Controller struct {
	// handleFunc 的格式需要是func(ctx context.Context, req interface{}) (interface{}, error)
	handleFunc interface{}
}

func (*Controller) HandleHTTP(http.ResponseWriter, *http.Request, httprouter.Params) {

}

func (c *Controller) bindRequest(r *http.Request) (interface{}, error) {
	request := reflect.New(reflect.TypeOf(c.handleFunc).In(1).Elem()).Interface()

	return request, nil
}
