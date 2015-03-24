package main

import (
	"net/http"
)

func serv() {
	http.HandleFunc("/query", queryHandle)
	http.HandleFunc("/insert", insertHandle)
	logger.Print("Start listening now...")

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		logger.Print("ListenAndServe error, Exit now.")
		logger.Fatal(err)
	}

}
