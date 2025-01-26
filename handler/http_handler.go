package handler

import (
	"api-backend/sample/app/controller"
	"net/http"
)

func HandleRequest() {
	http.ListenAndServe(":8080", controller.GetRouter())
}
