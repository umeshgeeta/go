// MIT License
// Author: Umesh Patil, Neosemantix, Inc.

// Package SMI contains HTTPS client and server sample implementation.
// It also contains some login examples for common social media platforms
// like Facebook.
package smi

import (
	"../util"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
)

// If you run this code from IDE (like GoLand) ensure that you set the
// environment variable GO_CFG_HOME to point to a directory which contains a
// directory 'smi-cert' which contains a certificate and a key for domain
// where you want to host this server.
//
// Certificates are generated using OpenSSL. If you use client side certificate
// validation, you will need to uncomment corresponding code.
//
// In this example hostname is "dev.smi.com" and the URL to be used is
// "https://dev.smi.com:8443". You can check this through browser or the client
// program too.
func Server() {

	cfgDir, err := util.GetCfgHomeDir()
	if err != nil {
		log.Fatal("Unable to get Config Home Directory to locate certificates and keys")
	}
	crtKeyDir := cfgDir + "/smi-cert/"
	fmt.Printf("Certificates and key will be read from directory:  %s\n", crtKeyDir)

	caCertPool := x509.NewCertPool()
	//caCertPool.AppendCertsFromPEM(caCert)
	cfg := &tls.Config{
		ClientAuth: tls.NoClientCert, //.RequireAndVerifyClientCert,
		ClientCAs:  caCertPool,
	}
	srv := &http.Server{
		Addr:      ":8443",
		Handler:   &handler{},
		TLSConfig: cfg,
	}
	//log.Fatal(srv.ListenAndServeTLS(crtKeyDir+"server.crt", crtKeyDir+"server.key"))
	log.Fatal(srv.ListenAndServeTLS(crtKeyDir+"dev.smi.com.crt", crtKeyDir+"dev.smi.com.key"))
}

type handler struct{}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//w.Write([]byte("PONG\n"))
	if r.URL.Path != "/" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	switch r.Method {
	case "GET":
		//absPath, _ := filepath.Abs("../src/smi/about.html")
		absPath, _ := filepath.Abs("../src/smi/fb-login.html")
		http.ServeFile(w, r, absPath)
		//http.Error(w, "406 not found.", http.StatusNotAcceptable)

	case "POST":
		decoder := json.NewDecoder(r.Body)
		var t PayloadToPublish
		err := decoder.Decode(&t)
		if err != nil {
			panic(err)
		}
		log.Printf("%v\n", t)
		w.Write([]byte("Received Post request \n"))

	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}
