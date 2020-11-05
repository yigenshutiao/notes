package common

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"reflect"
)

func HTTPHandler(handlefunc interface{}) httprouter.Handle {
	controller := &Controller{handleFunc: handlefunc}
	return controller.HandleHTTP
}

type Controller struct {
	handleFunc interface{}
}

func (*Controller) HandleHTTP(http.ResponseWriter, *http.Request, httprouter.Params) {

}

func (c *Controller) bindRequest(r *http.Request) (interface{}, error) {
	request := reflect.New(reflect.TypeOf(c.handleFunc).In(1).Elem()).Interface()

	return request, nil
}
