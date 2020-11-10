package main

import (
	"net/http"
	"strings"
)

type server struct {
	fileHandler http.Handler
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch {
	// case strings.HasSuffix(r.URL.Path, ".reality"):
	// 	w.Header().Set("Content-Type", "x-reality")
	case strings.HasSuffix(r.URL.Path, ".usdz"):
		w.Header().Set("Content-Type", "model/usd")
		return
	}
	s.fileHandler.ServeHTTP(w, r)
}

func main() {
	http.ListenAndServe(":80", &server{fileHandler: http.FileServer(http.Dir("/htdocs"))})
}
