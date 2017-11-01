// Fileserver is a minimal service for to serving contents from the file system over HTTP.
package main

import (
	"fmt"
	"crypto/tls"
	"log"
	"net/http"
	"strings"
	"os"

	"golang.org/x/crypto/acme/autocert"
)

func main() {
	filesDir := "/var/www"
	fs := http.FileServer(http.Dir(filesDir))
	http.Handle("/", fs)

	addr := os.Getenv("BLOCKPRESS_ADDR")
        if addr == "" {
	    addr = ":8080"
	}
	s := &http.Server{
		Addr: addr,
	}
	if addr == ":443" {
		hostsenv := os.Getenv("BLOCKPRESS_HOSTS")
		if hostsenv == "" {
			log.Fatalf("BLOCKPRESS_HOSTS must be set to serve TLS.")
		}
		hosts := strings.Split(hostsenv, ";")
		fmt.Printf("Serving TLS as %q..\n", hosts)
		m := autocert.Manager{
			Prompt:     autocert.AcceptTOS,
			Cache:      autocert.DirCache("/etc/secrets/acme/"),
			HostPolicy: autocert.HostWhitelist(hosts...),
		}
		s.TLSConfig = &tls.Config{GetCertificate: m.GetCertificate}
		log.Fatal(s.ListenAndServeTLS("", ""))
	} else {
		fmt.Printf("Serving plaintext HTTP on %s..\n", addr)
		log.Fatal(s.ListenAndServe())
	}
}
