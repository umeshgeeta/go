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

var chIncoming chan PayloadToPublish
var chToFb chan FbPayload
var chToTwitter chan TwitterPayload

func Server() {

	cfgDir, err := util.GetCfgHomeDir()
	if err != nil {
		log.Fatal("Unable to get Config Home Directory to locate certificates and keys")
	}
	crtKeyDir := cfgDir + "/smi-cert/"
	fmt.Printf("Certificates and key will be read from directory:  %s\n", crtKeyDir)

	caCert, err := ioutil.ReadFile(crtKeyDir + "client.crt")
	if err != nil {
		log.Fatal(err)
	}

	chIncoming = make(chan PayloadToPublish)
	chToFb = make(chan FbPayload)
	chToTwitter = make(chan TwitterPayload)

	go propagate()
	go fb()
	go twitter()

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

		chIncoming <- t

	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}

func propagate() {
	for {
		payload := <-chIncoming
		fp := FbPayload{payload.FbAuthCred, payload.Content}
		chToFb <- fp
		tp := TwitterPayload{payload.TwitterCred, payload.Content}
		chToTwitter <- tp
	}
}

func fb() {
	for {
		pl := <- chToFb
		fmt.Printf("Publishing %v to Fb\n", pl)
	}
}

func twitter() {
	for {
		pl := <- chToTwitter
		fmt.Printf("Publishing %v to Twitter\n", pl)
	}
}
