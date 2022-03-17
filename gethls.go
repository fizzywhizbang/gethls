package gethls

import (
	"crypto/tls"
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"net/http"
)

const StatusUrl = "https://chaturbate.com/api/chatvideocontext/"

var (
	insecure = flag.Bool("insecure-ssl", false, "Accept/Ignore all server SSL certificates")
)

func GetBCStatus(name string) Room {
	url := StatusUrl + name + "/"
	body := GetBody(url)
	var room Room
	json.Unmarshal(body, &room)

	return room
}

func GetBody(url string) []byte {

	flag.Parse()

	config := &tls.Config{
		InsecureSkipVerify: *insecure,
	}

	tr := &http.Transport{TLSClientConfig: config}
	client := &http.Client{Transport: tr}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", "Mozilla/5.0")
	res, getErr := client.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	return body
}
