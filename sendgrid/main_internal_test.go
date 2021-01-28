package sendgrind

import (
	"encoding/json"
	"io/ioutil"
	"testing"
)

func Test(t *testing.T) {
	t.Log("Starting sendgrid test")
	// Setup MailData sample
	fileData, err := readFile("sample.json")
	if err != nil {
		t.Errorf("readFile %v", err)
	}
	t.Log("fetched sample data")
	md := MailData{}
	err = json.Unmarshal(fileData, &md)
	if err != nil {
		t.Errorf("json.Unmarshal %v", err)
	}

	// Set up custom config data
	configData, err := readFile("config.json")
	if err != nil {
		t.Errorf("readFile %v", err)
	}
	t.Log("fetched config data")
	config := map[string]string{}
	err = json.Unmarshal(configData, &config)
	if err != nil {
		t.Errorf("json.Unmarshal %v", err)
	}

	md.Personalizations[0].To[0].Email = config["email"]
	err = md.Send(config["apikey"])
	if err != nil {
		t.Errorf("md.Send %v", err)
	}
}

func readFile(path string) ([]byte, error) {
	jsonFile, err := ioutil.ReadFile(path)
	if err != nil {
		return []byte{}, err
	}
	return jsonFile, nil
}
