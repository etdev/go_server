package main

import (
	"flag"
	"net/http"
)

func main() {
	var port = flag.String("p", "8000", "Please enter a valid port")
	var root_dir = flag.String("r", "public", "Please enter a valid root directory")
	flag.Parse()

	http.ListenAndServe(":"+*port, http.FileServer(http.Dir(*root_dir)))
}
