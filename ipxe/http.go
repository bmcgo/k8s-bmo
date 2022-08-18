package ipxe

import (
	_ "embed"
	"log"
	"net/http"
	"text/template"
)

//go:embed ipxe_script_template.txt
var ipxeScriptTemplate string

type ipxeScriptParameters struct {
	ServerAddr string
	Kernel     string
	Initrd     string
}

type HttpServer struct {
	listenAddr           string
	ipxeScript           []byte
	ipxeScriptParameters ipxeScriptParameters
	ipxeScriptTemplate   template.Template
}

func NewHttpServer(listenAddr string, ipxeScriptParams ipxeScriptParameters) (*HttpServer, error) {
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

func (s *HttpServer) Start() error {
	mux := http.NewServeMux()
	mux.HandleFunc("/ipxe", s.handleIPXE)
	err := http.ListenAndServe(s.listenAddr, mux)
	return err
}
