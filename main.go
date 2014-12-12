package main

import (
	"io/ioutil"
	"net/http"
	"os"

	"github.com/songgao/nush"
	"golang.org/x/crypto/ssh"
)

func main() {
	s, err := buildSSHAcceptor()
	if err != nil {
		panic(err)
	}
	h, err := buildHTTPAcceptor()
	if err != nil {
		panic(err)
	}
	_ = h
	go nush.TerminalServer(nil).ListenAndServe([]nush.Acceptor{s, h})
	go http.ListenAndServe(":80", http.RedirectHandler("https://song.gao.io/", http.StatusMovedPermanently))
	select {}
}

func buildHTTPAcceptor() (nush.Acceptor, error) {
	acceptor, mux, err := nush.NewHTTPListener()
	if err != nil {
		return nil, err
	}
	assets := "./assets"
	if len(os.Args) > 1 {
		assets = os.Args[1]
	}
	mux.Handle("/", http.FileServer(http.Dir(assets)))
	go http.ListenAndServeTLS(":443", "./conf/bundle.pem", "./conf/ca.key", mux)
	return acceptor, nil
}

func buildSSHAcceptor() (nush.Acceptor, error) {
	config := &ssh.ServerConfig{
		NoClientAuth: true,
	}
	privateBytes, err := ioutil.ReadFile("./conf/id_rsa")
	if err != nil {
		return nil, err
	}
	private, err := ssh.ParsePrivateKey(privateBytes)
	if err != nil {
		return nil, err
	}
	config.AddHostKey(private)
	l, err := nush.NewSSHListener(config, "localhost:22")
	if err != nil {
		return nil, err
	}
	return l, nil
}
