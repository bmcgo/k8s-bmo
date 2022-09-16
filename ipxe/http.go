package ipxe

import (
	_ "embed"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"text/template"
)

//go:embed ipxe_script_template.txt
var ipxeScriptTemplate string

type IPXEScriptParameters struct {
	ServerAddr string
	Kernel     string
	Initrd     string
}

type HttpServer struct {
	listenAddr           string
	ipxeScript           []byte
	ipxeScriptParameters IPXEScriptParameters
	ipxeScriptTemplate   template.Template
}

func NewHttpServer(listenAddr string, ipxeScriptParams IPXEScriptParameters) (*HttpServer, error) {
	tmpl, err := template.New("ipxe-script-template").Parse(ipxeScriptTemplate)
	if err != nil {
		return nil, err
	}
	server := &HttpServer{
		listenAddr:           listenAddr,
		ipxeScriptParameters: ipxeScriptParams,
		ipxeScriptTemplate:   *tmpl,
	}
	return server, nil
}

func (s *HttpServer) handleIPXE(w http.ResponseWriter, r *http.Request) {
	err := s.ipxeScriptTemplate.Execute(w, s.ipxeScriptParameters)
	if err != nil {
		log.Println(r, err)
		return
	}
	log.Println(r, "ok")
}

func (s *HttpServer) handleKernel(w http.ResponseWriter, r *http.Request) {
	fd, err := os.Open("/store/iso/boot/bzImage")
	if err != nil {
		log.Println(err)
		return
	}
	_, err = io.Copy(w, fd)
	if err != nil {
		log.Println(err)
		return
	}
}

func (s *HttpServer) handleInitrd(w http.ResponseWriter, r *http.Request) {
	sendFile(w, "/store/iso/boot/initrd")
}

func sendFile(w http.ResponseWriter, filename string) error {
	fi, err := os.Stat(filename)
	if err != nil {
		log.Println(err)
	}
	w.Header().Set("Content-Length", fmt.Sprintf("%d", fi.Size()))
	w.WriteHeader(http.StatusOK)
	fd, err := os.Open(filename)
	bs, err := io.Copy(w, fd)
	if err != nil {
		log.Println(err)
	}
	log.Printf("%d bytes sent", bs)
	return err
}

func (s *HttpServer) Start() {
	go func() {
		mux := http.NewServeMux()
		mux.HandleFunc("/ipxe", s.handleIPXE)
		mux.HandleFunc("/kernel", s.handleKernel)
		mux.HandleFunc("/initrd", s.handleInitrd)
		err := http.ListenAndServe(s.listenAddr, mux)
		if err != nil {
			log.Println(err)
		}
	}()
}
