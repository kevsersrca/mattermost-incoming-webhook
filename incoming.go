package miw

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type Payload struct {
	Text string `json:"text"`
	URL  string `json:"url"`
}

func (m *Payload) Do() error {
	data := m
	payloadBytes, err := json.Marshal(data)
	if err != nil {
		return err
	}
	body := bytes.NewReader(payloadBytes)

	req, err := http.NewRequest("POST", m.URL, body)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}
