package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/net/webdav"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	addr := flag.String("addr", "127.0.0.1:8080", "listen address (host:port)")
	dir := flag.String("dir", ".", "directory to serve")
	verbose := flag.Bool("verbose", false, "log file access")
	flag.Parse()

	h := &webdav.Handler{
		FileSystem: webdav.Dir(*dir),
		LockSystem: webdav.NewMemLS(),
	}
	if *verbose {
		h.Logger = func(r *http.Request, err error) {
			if err != nil {
				log.Printf("%s %s error=%v", r.Method, r.URL.Path, err)
				return
			}
			log.Printf("%s %s", r.Method, r.URL.Path)
		}
	}

	log.Printf("serving %s on http://%s", *dir, *addr)
	return http.ListenAndServe(*addr, h)
}
