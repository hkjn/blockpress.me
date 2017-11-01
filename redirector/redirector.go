// A simple redirect server that forwards any http:// request to https://.
package main

import (
	"fmt"
	"net/http"
	"log"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		next := fmt.Sprintf("https://%s%s", r.Host, r.URL.Path)
		log.Printf("Redirecting to %q..\n", next)
		http.Redirect(w, r, next, http.StatusFound)
	})
	addr := ":80"
	fmt.Printf("Serving redirect server on %s..\n", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
