package main

import (
	"notes/common"
	"notes/logic"

	"github.com/julienschmidt/httprouter"
)

func initRouter() *httprouter.Router {
	router := httprouter.New()

	router.GET("/", logic.Hello)
	router.GET("/note", common.HTTPHandler(logic.GetAll))
	router.POST("/note/:id", common.HTTPHandler(logic.GetOne))
	router.POST("/note", common.HTTPHandler(logic.Add))
	router.DELETE("/note/:id", common.HTTPHandler(logic.Delete))

	return router
}
