package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/lucas-clemente/quic-go/http3"
)

func RunHTTP3Client(path string) {
	log.Println("using:", path)
	h3Client := http.Client{
		Transport: &http3.RoundTripper{},
	}
	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		log.Println("new request error:", err)
		return
	}
	req.Header["User-Agent"] = []string{"http3.client.dilfish.dev"}
	log.Println("request is:", req)
	resp, err := h3Client.Do(req)
	if err != nil {
		log.Println("do error:", err)
		return
	}
	defer resp.Body.Close()
	bt, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("read all error:", err)
		return
	}
	log.Println("Header are:")
	for k, v := range resp.Header {
		log.Println(k, v)
	}
	if *FlagI {
		return
	}
	ct := http.DetectContentType(bt)
	if len(resp.Header["Content-Type"]) > 0 && len(resp.Header["Content-Type"][0]) > 0 {
		ct = resp.Header["Content-Type"][0]
	}
	log.Println("RESPONSE TYPE IS:", ct)
	if strings.Index(ct, "text") >= 0 {
		log.Println(string(bt))
	} else {
		log.Println("Binary 100 bytes:", string(bt[:100]))
		log.Println("Binary 100 bytes:", bt[:100])
	}
}
