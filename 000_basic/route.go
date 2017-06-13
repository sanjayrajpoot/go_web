package main

import (
		"fmt"
		"net/http"
)

func main() {

	userAges := map[string]int{
		"Alice": 26,
		"Bob": 27,
		"Claire":29,
	}

	http.HandleFunc("/users/", func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Path[len("/users/"):]
		age := userAges[name]
		fmt.Fprintf(w, "%s is %d year old", name, age)
		})
	http.ListenAndServe(":8080",nil)
    	
}