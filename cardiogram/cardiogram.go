package cardiogram

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"time"
)

// Heartbeat contains the configuration for Opsgenie Heartbeats.
type Heartbeat struct {
	Client  *http.Client
	Timeout time.Duration
	URL     string
	APIKey  string
}

// Check scrapes the targets and send the heartbeats to Opsgenie.
func (h *Heartbeat) Check(url string, expected int, name string) {
	if h.call(url, expected) == nil {
		h.send(name)
	}
}

func (h *Heartbeat) call(url string, expected int) error {
	res, err := h.Client.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != expected {
		return errors.New("Target returns an unexpected status code")
	}
	return nil
}

func (h *Heartbeat) send(name string) {
	req := struct {
		APIKey string `json:"apiKey"`
		Name   string `json:"name"`
	}{h.APIKey, name}

	buf, err := json.Marshal(req)
	if err != nil {
		log.Println("Cannot marshal request json")
		return
	}

	res, err := h.Client.Post(h.URL, "application/json", bytes.NewReader(buf))
	if err != nil {
		log.Printf("Error while sending Heartbeat for '%s': %s", name, err)
	}
	defer res.Body.Close()

	resp := struct {
		Status string
		Code   int
	}{}

	if err := json.NewDecoder(res.Body).Decode(&resp); err != nil {
		log.Println("Cannot read response from Opsgenie")
		return
	}

	if resp.Code != 200 || resp.Status != "successful" {
		log.Println("Sending Heartbeat was not successful")
	}
}
