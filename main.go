package main

import (
	"flag"
	"log"
)

const CertPath = "dilfish.ml/fullchain.cer"
const KeyPath = "dilfish.ml/dilfish.ml.key"

var FlagRunServer = flag.Bool("s", false, "if run server")
var FlagUrl = flag.String("u", "https://dev.ug/ip", "url")
var FlagI = flag.Bool("I", false, "same as curl -I")

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
