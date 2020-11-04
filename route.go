package main

import (
	"notes/logic"

	"github.com/julienschmidt/httprouter"
)

func initRouter() *httprouter.Router {

	router := httprouter.New()
	router.GET("/", logic.Hello)

	router.GET("/note", logic.GetAll)
	router.POST("/note/:id", logic.GetOne)
	router.POST("/note", logic.Add)
	router.DELETE("/note/:id", logic.Delete)

	return router
}
