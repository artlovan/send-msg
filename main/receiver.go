package main

import (
	"net/http"
	"io/ioutil"
	"fmt"
)

func handler(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	body, _ := ioutil.ReadAll(r.Body)

	fmt.Println(string(body))
}

func main() {
	service := http.NewServeMux()
	service.HandleFunc("/api/msg", handler)
	http.ListenAndServe(":6061", service)
}
