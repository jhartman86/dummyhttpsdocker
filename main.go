package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.Header().Add("Strict-Transport-Security", "max-age=63072000; includeSubDomains")
		w.Write([]byte(fmt.Sprintf(
			"This is an example HTTPS server. Time: %s\n",
			time.Now().Format(time.RFC1123),
		)))
		w.WriteHeader(http.StatusOK)
	})
	cfg := &tls.Config{
		InsecureSkipVerify:       true,
		MinVersion:               tls.VersionTLS12,
		CurvePreferences:         []tls.CurveID{tls.CurveP521, tls.CurveP384, tls.CurveP256},
		PreferServerCipherSuites: true,
		CipherSuites: []uint16{
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
			tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_RSA_WITH_AES_256_CBC_SHA,
		},
	}
	srv := &http.Server{
		Addr:         ":443",
		Handler:      mux,
		TLSConfig:    cfg,
		TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler), 0),
	}
	go startHTTP()
	fmt.Println("----------------------------")
	fmt.Println("STARTING SSL SERVICE ON :443")
	fmt.Println("----------------------------")
	log.Fatal(srv.ListenAndServeTLS("/root/.ssl-certs/server.crt", "/root/.ssl-certs/server.key"))
}

func startHTTP() error {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte(fmt.Sprintf(
			"This is an example HTTP server. Time: %s\n",
			time.Now().Format(time.RFC1123),
		)))
		w.WriteHeader(http.StatusOK)
		// w.Header().Add("Strict-Transport-Security", "max-age=63072000; includeSubDomains")
	})
	srv := &http.Server{
		Addr:    ":80",
		Handler: mux,
	}
	fmt.Println("----------------------------")
	fmt.Println("STARTING HTTP SERVICE ON :80")
	fmt.Println("----------------------------")
	log.Fatal(srv.ListenAndServe())
	return nil
}
