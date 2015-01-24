package main

import (
	"log"
	"net/http"
	"strings"
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

	//configuration REST endpoint
	http.HandleFunc("/configure/", configure)

	//Listen for connections and serve
	log.Println("Listening...")
	err := http.ListenAndServe(":3030", nil)
	if err != nil {
		log.Println(err.Error())
	}
}

func configure(w http.ResponseWriter, req *http.Request) {
	resourceList := req.FormValue("list")
	parseResourceList(resourceList)
	w.Write([]byte("Success!  http://localhost:3030/config.html"))
}

func parseResourceList(resourceList string) []string {
	return strings.Split(resourceList, "\n")
}
