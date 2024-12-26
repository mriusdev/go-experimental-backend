package controller

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Route struct {
	method      string
	path        string
	httpHandler func(rw http.ResponseWriter, request *http.Request)
}

var routes = []Route{
	{
		http.MethodGet,
		"/info",
		GetInfo,
	},
	{
		http.MethodPost,
		"/info",
		PostInfo,
	},
}

// func RouteRequestWithMethod(rw http.ResponseWriter, request *http.Request) {
// 	// find handler by
// 	// info
// 	// path
// 	// request.Method
// 	fmt.Println("handling request")
// 	for _, route := range routes {
// 		if request.Method != route.method || request.URL.Path != route.path {
// 			fmt.Println("bad request", request.URL.Path, request.Method)
// 			continue
// 		}

// 		fmt.Println("good request", [][]string{
// 			{"request", request.URL.Path, request.Method},
// 			{"routes list", route.path, route.method},
// 		})

// 		route.httpHandler(rw, request)
// 	}
// 	// fmt.Fprint(rw, request.URL.Path)

// 	router := chi.NewRouter()
// 	router.Route("/login", func(r chi.Router) {
// 		r.Post("/", Login)
// 	})
// }

func GetRouter() chi.Router {
	router := chi.NewRouter()
	router.Route("/login", func(r chi.Router) {
		r.Post("/", Login)
	})
	router.Route("/info", func(r chi.Router) {
		r.Get("/", GetInfo)
	})
	router.Route("/users", func(r chi.Router) {
		r.Get("/", GetUsers)
		r.Post("/", CreateUser)
		r.Patch("/{id}", UpdateUser)
		r.Get("/{id}", GetUser)
		r.Delete("/{id}", DeleteUser)
	})

	return router
}
