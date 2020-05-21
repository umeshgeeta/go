/*
 * MIT License
 * Copyright (c) 2020. Neosemantix, Inc.
 * Author: Umesh Patil
 */

package smi

import (
	"../util"
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var fbAuth FbAuthCred
var payload2Pub PayloadToPublish

func Client() {

	readConfig()

	cfgDir, err := util.GetCfgHomeDir()
	if err != nil {
		log.Fatal("Unable to get Config Home Directory to locate certificates and keys")
	}
	crtKeyDir := cfgDir + "/smi-cert/"

	// it should not be the server certificate, but rather the CA certificate
	//caCert, err := ioutil.ReadFile(crtKeyDir + "server.crt")
	caCert, err := ioutil.ReadFile(crtKeyDir + "ca.crt")
	if err != nil {
		log.Fatal(err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	cert, err := tls.LoadX509KeyPair(crtKeyDir+"client.crt", crtKeyDir+"client.key")
	if err != nil {
		log.Fatal(err)
	}

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				RootCAs:      caCertPool,
				Certificates: []tls.Certificate{cert},
			},
		},
	}

	//url := "https://127.0.0.1:8443/about.html"
	//url := "https://127.0.0.1:8443/fb-login.html"
	url := "https://127.0.0.1:8443/"
	//url := "https://localhost:8443/"
	getAboutHtml(client, url)
	postSampleContent(client, url)
}

func readConfig() {
	// Read the configuration - meaning AWS bucket names and local destination
	err := util.ReadCfg(&payload2Pub, "smi-client.json")
	if err != nil {
		fmt.Printf("%v\n", err)
		log.Fatal("Could not read smi-client configurations")
	} else {
		fmt.Printf("%v\n", payload2Pub)
	}
}

func postSampleContent(client *http.Client, url string) {
	caption := "Corona Days: Hate as Escape from Fear"
	imageUrl := "https://1.bp.blogspot.com/-BOGCAOd89hQ/XpQXjTOLl2I/AAAAAAAAFxs/eeavwgd270A_RFdxPHhHAsBb8scm9VRHACLcBGAsYHQ/s1600/download.jpg"
	contentLink := "https://21centurypolitics.com/blog_post.html?id=405"
	cc := NewCommonContent(caption, imageUrl, contentLink)
	payload2Pub.Content = *cc
	jsonStr, err := json.Marshal(payload2Pub)
	if err != nil {
		fmt.Println("error:", err)
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
}

func getAboutHtml(client *http.Client, url string) {
	resp, err := client.Get(url)
	if err != nil {
		log.Println(err)
		return
	}
	htmlData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	fmt.Printf("%v\n", resp.Status)
	fmt.Printf(string(htmlData))
}