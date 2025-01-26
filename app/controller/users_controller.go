package controller

import (
	"fmt"
	"net/http"
)

func GetUsers(rw http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(rw, "getting users")
}

func GetUser(rw http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(rw, "getting user")
}

func DeleteUser(rw http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(rw, "deleting user")
}

func UpdateUser(rw http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(rw, "updating user")
}

func CreateUser(rw http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(rw, "creating user")
}
