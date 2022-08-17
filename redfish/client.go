package redfish

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/go-logr/logr"
	"log"
	"net/http"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
)

type Client struct {
	endpoint   string
	urlSystems string
	client     *http.Client
	l          logr.Logger
	ctx        context.Context
}

type ClientConfig struct {
	URL     string
	Logger  *logr.Logger
	Context *context.Context
}

func NewClient(c ClientConfig) *Client {
	var ctx context.Context
	var l logr.Logger
	if c.Context == nil {
		ctx = context.Background()
	} else {
		ctx = *c.Context
	}
	if c.Logger == nil {
		l = zap.New()
	}
	return &Client{
		endpoint:   c.URL,
		urlSystems: c.URL + "/redfish/v1/Systems",
		client:     &http.Client{},
		l:          l,
		ctx:        ctx,
	}
}

func (s *Client) Post(url string, body interface{}) error {
	data, err := json.Marshal(body)
	if err != nil {
		return err
	}
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(data))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		return err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return err
	}
	log.Println(resp)
	return nil
}

func (s *Client) Patch(url string, patch interface{}) error {
	data, err := json.Marshal(patch)
	if err != nil {
		return err
	}
	req, err := http.NewRequest(http.MethodPatch, url, bytes.NewBuffer(data))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		return err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return err
	}
	log.Println(resp)
	return nil
}

func (s *Client) Get(url string, target interface{}) error {
	s.l.Info("Calling GET", "url", url)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		s.l.Error(err, "GET", "resp", resp)
		return err
	}
	return json.NewDecoder(resp.Body).Decode(target)
}

func (s *Client) GetManagerById(id string) (manager Manager, err error) {
	jManager := jsonManager{}
	err = s.Get(s.endpoint+id, &jManager)
	if err != nil {
		return
	}
	manager.manager = jManager
	manager.client = s
	return
}
