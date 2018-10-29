package main

import (
	"fmt"
	"net/http"
	//"time"
)

func routeHandler(w http.ResponseWriter, r *http.Request){
	http.Redirect(w, r, "/paragliding/api", 301)
}

func ApiResponse(w http.ResponseWriter, r *http.Request){

	fmt.Printf("hello world")
}

func main() {
	root := "/paragliding"
	http.HandleFunc(root, routeHandler)

	http.HandleFunc(root + "/api", ApiResponse)

}


