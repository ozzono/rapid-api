package sendgrind

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// MailData contains all the needed info to send an email
type MailData struct {
	Personalizations []Personalization `json:"personalizations"`
	From             struct {
		Email string `json:"email"`
	} `json:"from"`
	Content []Content `json:"content"`
}

// Content uses text/plain as default Type
type Content struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

// Personalization ...
type Personalization struct {
	To []struct {
		Email string `json:"email"`
	} `json:"to"`
	Subject string `json:"subject"`
}

// Send 's the mail as customized by MailData
func (md *MailData) Send(APIKey string) error {
	url := "https://rapidprod-sendgrid-v1.p.rapidapi.com/mail/send"

	payload, err := json.Marshal(md)
	if err != nil {
		return fmt.Errorf("json.Marshal: %v", err)
	}

	req, err := http.NewRequest("POST", url, strings.NewReader(string(payload)))
	if err != nil {
		return fmt.Errorf("http.NewRequest %v", err)
	}

	req.Header.Add("content-type", "application/json")
	req.Header.Add("x-rapidapi-key", APIKey)
	req.Header.Add("x-rapidapi-host", "rapidprod-sendgrid-v1.p.rapidapi.com")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("http.DefaultClient.Do %v", err)
	}

	defer res.Body.Close()
	if res.StatusCode != 200 && res.StatusCode != 202 {
		return fmt.Errorf("failed to send email: %v", res.Status)
	}
	return nil
}
