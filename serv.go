package main

import (
	"net/http"
)

func serv() {
	http.HandleFunc("/show", show)
	http.HandleFunc("/insert", handle)
	logger.Print("Start listening now...")

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		logger.Print("ListenAndServe error, Exit now.")
		logger.Fatal(err)
	}

}
