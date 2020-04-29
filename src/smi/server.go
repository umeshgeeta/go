/*
 * MIT License
 * Copyright (c) 2020. Neosemantix, Inc.
 * Author: Umesh Patil
 */

package smi

import (
	"../util"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
)

func Server() {

	cfgDir, err := util.GetCfgHomeDir()
	if err != nil {
		log.Fatal("Unable to get Config Home Directory to locate certificates and keys")
	}
	crtKeyDir := cfgDir + "/smi-cert/"

	caCert, err := ioutil.ReadFile(crtKeyDir + "client.crt")
	if err != nil {
		log.Fatal(err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)
	cfg := &tls.Config{
		ClientAuth: tls.RequireAndVerifyClientCert,
		ClientCAs:  caCertPool,
	}
	srv := &http.Server{
		Addr:      ":8443",
		Handler:   &handler{},
		TLSConfig: cfg,
	}
	log.Fatal(srv.ListenAndServeTLS(crtKeyDir+"server.crt", crtKeyDir+"server.key"))
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
		absPath, _ := filepath.Abs("../src/smi/about.html")
		http.ServeFile(w, r, absPath)
		//http.Error(w, "406 not found.", http.StatusNotAcceptable)

	case "POST":
		decoder := json.NewDecoder(r.Body)
		var t SmContent
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
