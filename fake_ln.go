package main

import (
	"log"
	"net"
)

type FakeListener struct {
	real net.Listener
}

func (f *FakeListener) Accept() (net.Conn, error) {
	c, err := f.real.Accept()
	if err != nil {
		log.Println("fakels.accept error:", err)
		return c, err
	}
	var fakeConn FakeConn
	fakeConn.real = c
	return &fakeConn, nil
}

func (f *FakeListener) Close() error {
	err := f.real.Close()
	log.Println("fake ls close")
	return err
}

func (f *FakeListener) Addr() net.Addr {
	a := f.real.Addr()
	log.Println("fake ls addr:", a)
	return a
}
