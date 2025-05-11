package handler

import (
	"fmt"
	"net/http"
)

func GreetingHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Haloo, selamat datang di dibimbing GO3!")
}
