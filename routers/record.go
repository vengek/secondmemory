package routers

import (
	"github.com/gorilla/mux"
	"/controllers"
)

func SetRecordRoutes(router *mux.Router) *mux.Router  {
	recordRouter := mux.NewRouter()
	recordRouter.HandleFunc("/records", controllers.CreateRecord).Methods("POST")
	recordRouter.HandleFunc("/records/{id}", controllers.UpdateRecord).Methods("PUT")

	return router
}
