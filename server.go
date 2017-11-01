// Fileserver is a minimal service for to serving contents from the file system over HTTP.
package main

import (
	"fmt"
	"crypto/tls"
	"log"
	"net/http"
	"os"

	"golang.org/x/crypto/acme/autocert"
)

var subdomains = []string{
	"anton", "dana", "hkjn",
}

type multiHostHandler struct {
	fileHandlers map[string]http.Handler
}

func (mhh multiHostHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("Got request for https://%s%s", r.Host, r.URL.Path)
	inner, ok := mhh.fileHandlers[r.Host]
	if !ok {
		log.Printf("No domain %s, serving 404.\n", r.Host)
		http.NotFound(w, r)
		return
	}
	inner.ServeHTTP(w, r)
}

func main() {
	hostnames := []string{
		"blockpress.me",
	}
	fh := map[string]http.Handler{
		"blockpress.me": http.FileServer(http.Dir("/var/www/root")),
	}
	for _, name := range subdomains {
		sd := fmt.Sprintf("%s.blockpress.me", name)
		fp := fmt.Sprintf("/var/www/%s", name)
		log.Printf("Serving %q from file path %q..\n", sd, fp)
		fh[sd] = http.FileServer(http.Dir(fp))
		hostnames = append(hostnames, sd)
	}
	http.Handle("/", multiHostHandler{fh})

	addr := os.Getenv("BLOCKPRESS_ADDR")
        if addr == "" {
	    addr = ":8080"
	}
	s := &http.Server{
		Addr: addr,
	}
	if addr == ":443" {
		fmt.Printf("Serving TLS as %q..\n", hostnames)
		m := autocert.Manager{
			Prompt:     autocert.AcceptTOS,
			Cache:      autocert.DirCache("/etc/secrets/acme/"),
			HostPolicy: autocert.HostWhitelist(hostnames...),
		}
		s.TLSConfig = &tls.Config{GetCertificate: m.GetCertificate}
		log.Fatal(s.ListenAndServeTLS("", ""))
	} else {
		fmt.Printf("Serving plaintext HTTP on %s..\n", addr)
		log.Fatal(s.ListenAndServe())
	}
}
