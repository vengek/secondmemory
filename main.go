package main

import (
	"net/http"
	"github.com/codegangsta/negroni"
	"./common"
	"./routers"
)

func main() {
	common.Start()
	router := routers.InitRoutes()

	n := negroni.Classic()
	n.UseHandler(router)
	server := &http.Server{
		Addr: ":8989",
		Handler: n,
	}
	server.ListenAndServe()
}