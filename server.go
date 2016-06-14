package main

import (
	"flag"
	"net/http"
)

func main() {
	var port = flag.String("p", "8000", "Please enter a valid port")
	http.ListenAndServe(":"+*port, http.FileServer(http.Dir("public")))
}
