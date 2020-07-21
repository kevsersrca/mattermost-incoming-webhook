package miw

import (
	"encoding/json"
	"net/http"
	"strings"
)

type Webhook struct {
	URL      string
	IconURL  string `json:"icon_url"`
	Username string `json:"username"`
	Channel  string `json:"channel"`
}

type Message struct {
	Webhook
	Text string `json:"text"`
}

func (m *Message) Do() error {
	messageJSON, err := json.Marshal(m)
	if err != nil {
		return err
	}

	body := strings.NewReader(string(messageJSON))
	req, err := http.NewRequest("POST", m.URL, body)
	if err != nil {
		return err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	return nil
}
