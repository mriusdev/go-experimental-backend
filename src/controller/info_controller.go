package controller

import (
	"fmt"
	"net/http"
)

func GetInfo(rw http.ResponseWriter, request *http.Request) {
	fmt.Fprint(rw, "getting info")
}

func PostInfo(rw http.ResponseWriter, request *http.Request) {
	fmt.Fprint(rw, "posting info")
}
