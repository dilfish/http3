package main

import (
	"github.com/lucas-clemente/quic-go/http3"
	"log"
	"net/http"
)

type Handler struct{}

func (h *Handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	log.Printf("request header is %+v", req.Header)
	log.Println("req addr is", req.RemoteAddr)
	w.Header().Set("Alt-Svc", "h3=\":443\"; ma=3600, h3-29=\":443\"; ma=3600")
	w.Write([]byte("good"))
}

func RunHTTP2() {
	var err error
	err = http.ListenAndServeTLS(":443", *FlagCertPath, *FlagKeyPath, nil)
	if err != nil {
		panic(err)
	}
}

func RunHTTP3() {
	var h Handler
	http.Handle("/", &h)
	err := http3.ListenAndServeQUIC(":443", *FlagCertPath, *FlagKeyPath, nil)
	if err != nil {
		panic(err)
	}
}
