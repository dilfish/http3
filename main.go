package main

import (
	"flag"
	"log"
)

var FlagRunServer = flag.Bool("s", false, "if run server")
var FlagUrl = flag.String("u", "https://dev.ug/ip", "url")
var FlagI = flag.Bool("I", false, "same as curl -I")
var FlagCertPath = flag.String("cert", "", "cert file")
var FlagKeyPath = flag.String("key", "", "private key file")

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	flag.Parse()
	if *FlagRunServer {
		go RunHTTP3()
		RunHTTP2()
	} else {
		RunHTTP3Client(*FlagUrl)
	}
}
