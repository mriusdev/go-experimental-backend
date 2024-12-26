package handler

import (
	"api-backend/sample/src/controller"
	"net/http"
)

func HandleRequest() {
	// mux := http.NewServeMux()
	// http.HandleFunc("/", controller.RouteRequestWithMethod)

	http.ListenAndServe(":8080", controller.GetRouter())
}
