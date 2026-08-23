package handler

import (
	"api-backend/sample/app/application"
	"api-backend/sample/app/controller"
	"net/http"
)

func HandleRequest(apiHandler *application.ApiHandler) {
	http.ListenAndServe(":8080", controller.AddRoutes(apiHandler))
}
