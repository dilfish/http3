package main

import (
	"log"
	"net"
	"time"
)

type FakeConn struct {
	real net.Conn
}

func (f *FakeConn) Read(b []byte) (n int, err error) {
	n, err = f.real.Read(b)
	log.Println("fake conn read is:", string(b), n, err)
	return n, err
}

func (f *FakeConn) Write(b []byte) (n int, err error) {
	n, err = f.real.Write(b)
	log.Println("fake conn write is:", string(b), n, err)
	return n, err
}

func (f *FakeConn) Close() error {
	err := f.real.Close()
	log.Println("fake conn close:", err)
	return err
}

func (f *FakeConn) LocalAddr() net.Addr {
	a := f.real.LocalAddr()
	log.Println("fake conn localaddr is:", a)
	return a
}

func (f *FakeConn) RemoteAddr() net.Addr {
	a := f.real.RemoteAddr()
	log.Println("fake conn remoteaddr is:", a)
	return a
}

func (f *FakeConn) SetDeadline(t time.Time) error {
	log.Println("fake conn setdeadline:", t)
	return f.real.SetDeadline(t)
}

func (f *FakeConn) SetReadDeadline(t time.Time) error {
	log.Println("fake conn set read deadline:", t)
	return f.real.SetReadDeadline(t)
}

func (f *FakeConn) SetWriteDeadline(t time.Time) error {
	log.Println("fake conn set write deadline:", t)
	return f.real.SetWriteDeadline(t)
}
