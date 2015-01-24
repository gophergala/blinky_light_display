package main

import (
	"log"
	"net/http"
)

func main() {
	//Serve static configuration page
	//  GET config.html
	//	GET client whatnot
	//Serve configuration REST endpoint
	//	PUT /configuration/ {list of assets}
	//	Parse asset list
	//	Respond appropriately
	//Serve current display item uri
	//	GET /current/
	//	Respond w/ text/plain uri
	//	If none, 404? 500? I guess 404

	//static file server
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)

	log.Println("Listening...")
	http.ListenAndServe(":3030", nil)
}
