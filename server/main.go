package main

import "net/http"

func main() {
	homeRoute()

	http.ListenAndServe(":8080", nil)
}

func homeRoute() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})
}
