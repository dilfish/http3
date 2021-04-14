package main

import (
	"encoding/hex"
	"log"

	"github.com/marten-seemann/qpack"
)

func DecodeQPack() error {
	frame := "0000d95f44b49d983f9b8d34cff3f6a5238197000fa52765664507f371a699fe7ed4a47032e001f4a4ecac27e0fe6e34d33fcfda948e065c003f2f04f2b4e94d6cac5a871d0593834d96972bf2b1ab802df2b4e7427f89f739a7b578cd52f6cd2ef2b20b6772d98736b9cf5154711f2f02f2b585ecb4e51c85b3ffa801fffdfcc24216b4ff373ff7fca4be52c4e9a68fa1d75d0620d263d4c799d34d1fefcd347d0ebae83106931ea63cd347df6800bbff46a473158f058ebfaf963e7efb4005defe7ff9f5fcc24216b4ad82a21e435537f373ff7fccdecd5fd292165a0692fd291d9fcfff3ebf984842d695b38ea9ad1cc5fe6e7feff92d4b70ddf45abefb4005dbfe7ff9f5fcc34256e082c9fcdcffdff3ee734f6af19aa5ed9bf9ffe7d7f378649cab5e3d49b0f47f373ff7fcd23f2b0e62c00fe7ff9f5fcf0ae6b072156c9520a4b6c2adb4bdad2a128fe6e7feff90ff3ffcfafe7820b62d0cc5a93fcdcffdff3a0fecd450361b5c0a7f5a06435493a27fb5325492d0a681914d5b94fc50205c2e7da9677b8f36b83fb531149d4ec0801000200a984d61653f961e6d707f3ffcfff7"
	data := make([]byte, hex.DecodedLen(len(frame)))
	n, err := hex.Decode(data, []byte(frame))
	if err != nil {
		log.Println("decode error:", err)
		return err
	}
	data = data[:n]
	decoder := qpack.NewDecoder(nil)
	hfs, err := decoder.DecodeFull(data)
	if err != nil {
		log.Println("decode full error:", err)
		return err
	}
	values := make(map[string]string)
	for _, hf := range hfs {
		values[hf.Name] = hf.Value
	}
	req := `{"Accept":["text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8"],"Accept-Encoding":["gzip, deflate, br"],"Accept-Language":["en-US,en;q=0.5"],"Alt-Used":["dev.ug"],"Cache-Control":["max-age=0"],"Upgrade-Insecure-Requests":["1"],"User-Agent":["Mozilla/5.0 (Macintosh; Intel Mac OS X 10.16; rv:85.0) Gecko/20100101 Firefox/85.0"]}`
	values["x-host"] = "dev.ug"
	values["x-req-header"] = req
	log.Println("QPack Decoded Headers as:")
	for k, v := range values {
		log.Println(k, v)
	}
	log.Println("Thanks to: github.com/marten-seemann/qpack")
	return nil
}
