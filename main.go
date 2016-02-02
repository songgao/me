package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/songgao/nush"
	"golang.org/x/crypto/ssh"
)

func main() {
	s, err := buildSSHAcceptor()
	if err != nil {
		log.Fatal(err)
	}
	mux := buildMux()
	h, err := buildHTTPAcceptor(mux)
	if err != nil {
		log.Fatal(err)
	}
	nush.SetLoggerOutput(os.Stdout)
	go nush.TerminalServer(nil).ListenAndServe([]nush.Acceptor{s, h})
	log.Println("Ready")
	select {}
}

func buildMux() *http.ServeMux {
	assets := "./assets"
	if len(os.Args) > 1 {
		assets = os.Args[1]
	}
	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir(assets)))
	return mux
}

func buildHTTPAcceptor(mux *http.ServeMux) (nush.Acceptor, error) {
	acceptor, err := nush.NewHTTPListener(mux)
	if err != nil {
		return nil, err
	}
	go func() {
		server := &http.Server{
			Addr:         ":80",
			Handler:      mux,
			ReadTimeout:  32 * time.Second,
			WriteTimeout: 32 * time.Second,
		}
		if err = server.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()
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
	l, err := nush.NewSSHListener(config, ":22")
	if err != nil {
		return nil, err
	}
	return l, nil
}
