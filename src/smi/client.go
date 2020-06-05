// MIT License
// Author: Umesh Patil, Neosemantix, Inc.

package smi

import (
	"../util"
	"bytes"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var fbAuth FbAuthCred
var payload2Pub PayloadToPublish

// HTTPS client which makes one sample GET and POST calls. It expects
// GO_CFG_HOME environment is set up and it points to a directory which
// contains smi-client.json file. Under cfg folder you can find a template
// which you should change as per your local set up.
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

	// This is kind of redundant when client side cert validation is turned off.
	//cert, err := tls.LoadX509KeyPair(crtKeyDir+"client.crt", crtKeyDir+"client.key")
	//if err != nil {
	//	log.Fatal(err)
	//}

	client := &http.Client{
		Transport: &http.Transport{
			//TLSClientConfig: &tls.Config{
			//	RootCAs:      caCertPool,
			//	Certificates: []tls.Certificate{cert},
			//},
		},
	}

	url := "https://dev.smi.com:8443/"

	getAboutHtml(client, url)
	postSampleContent(client, url)
}

func readConfig() {
	// Read the configuration - credentials to FB and Twitter are shown
	// in case server wants to cash those. Though it is HTTPS, so safe to
	// send credentials; one would not use this method. The GET call
	// which gets FB login, user directly logs in to FB directly and then
	// a short term token is obtained which is used for subsequent FB API
	// calls. In that sense, this client and the approach of sharing credentials
	// through config file would not work. It is retained only for
	// demonstration purposes.
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
