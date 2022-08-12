package redfish

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

type Client struct {
	endpoint   string
	urlSystems string
	client     *http.Client
}

type ClientConfig struct {
	URL string
}

func NewClient(c ClientConfig) *Client {
	return &Client{
		endpoint:   c.URL,
		urlSystems: c.URL + "/redfish/v1/Systems",
		client:     &http.Client{},
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
	log.Println(url)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Println(err)
		return err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		log.Println(resp)
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
