package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", SympleServer)
	http.ListenAndServe(":8080", nil)
}
func SympleServer(w http.Response, r *http.Request) {
	fmt.Fprint(w, "Hello %s", r.URL.Path[1:])
}
