// The redirect command redirects all HTTP requests to a target URL.
package main

import (
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	target, err := url.Parse(os.Getenv("TARGET"))
	if err != nil {
		log.Fatalf("error parsing TARGET: %v", err)
	}
	if !target.IsAbs() {
		log.Fatal("TARGET must be absolute URL")
	}

	status, _ := strconv.Atoi(os.Getenv("STATUS"))
	if status == 0 {
		status = http.StatusFound
	}

	handle := func(w http.ResponseWriter, r *http.Request) {
		u := *(r.URL)
		u.Scheme = ""
		u.Host = ""
		u.Path = strings.TrimPrefix(u.Path, "/")
		dst := target.ResolveReference(&u).String()
		http.Redirect(w, r, dst, status)
	}

	log.Print("Listening on port " + port)
	log.Fatal(http.ListenAndServe(":"+port, http.HandlerFunc(handle)))
}
