package common

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"reflect"
)

const (
	Success           = "Success"
	Failed            = "Failed"
	JSONMarshalFailed = "JSONMarshalFailed"
)

type Controller struct {
	// handleFunc 的格式需要是func(ctx context.Context, req interface{}) (interface{}, error)
	handleFunc interface{}
}

func HTTPHandler(handleFunc interface{}) httprouter.Handle {
	controller := &Controller{handleFunc: handleFunc}

	checkHandleFunc(handleFunc)

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

func (c *Controller) HandleHTTP(wr http.ResponseWriter, r *http.Request, p httprouter.Params) {

	request, err := c.bindRequest(r)
	if err != nil {

	}

	if err := Validate.Struct(request); err != nil {

	}

	resp, err := c.callFunc(r, request)

	respStr := response2JSON(r.Context(), wr, resp, err)

	fmt.Println(respStr)

}

func (c *Controller) callFunc(r *http.Request, request interface{}) (interface{}, error) {
	f := reflect.ValueOf(c.handleFunc)
	returnVal := f.Call([]reflect.Value{reflect.ValueOf(r.Context()), reflect.ValueOf(request)})

	response := returnVal[0].Interface()
	err := returnVal[1].Interface().(error)

	return response, err
}

func (c *Controller) bindRequest(r *http.Request) (interface{}, error) {
	request := reflect.New(reflect.TypeOf(c.handleFunc).In(1).Elem()).Interface()

	return request, nil
}

type HTTPResponse struct {
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func response2JSON(ctx context.Context, wr http.ResponseWriter, resp interface{}, err error) string {

	respData := &HTTPResponse{
		Msg:  Success,
		Data: resp,
	}

	if err != nil {
		respData = &HTTPResponse{
			Msg:  Failed,
			Data: resp,
		}
	}

	res, err := json.Marshal(respData)
	if err != nil {
		respData = &HTTPResponse{
			Msg:  JSONMarshalFailed,
			Data: nil,
		}

		res = []byte(form2JSON(respData))
	}

	if err := writeResponse(wr, res); err != nil {
		return ""
	}

	return string(res)
}

func form2JSON(r *HTTPResponse) string {
	return fmt.Sprintf("{\"errmsg\":\"%v\",\"data\": \"\"}", r.Msg)
}

func writeResponse(wr http.ResponseWriter, res []byte) error {
	_, err := wr.Write(res)

	return err
}
