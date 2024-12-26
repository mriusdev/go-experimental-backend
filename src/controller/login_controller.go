package controller

import (
	"fmt"
	"net/http"
)

func Login(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "logging in...")
}
