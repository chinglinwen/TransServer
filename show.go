package main

import (
	"fmt"
	"net/http"
)

func show(w http.ResponseWriter, req *http.Request) {
	fmt.Printf("show test.")
	fmt.Printf("show test1.")
}
