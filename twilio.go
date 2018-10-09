package main

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func TwilioSendText(msgData url.Values, sid string, token string, twilioUrl string) error  {
	msgDataReader := *strings.NewReader(msgData.Encode())
	client := &http.Client{}
	req, _ := http.NewRequest("POST", twilioUrl, &msgDataReader)
	req.SetBasicAuth(sid, token)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		var data map[string]interface{}
		decoder := json.NewDecoder(resp.Body)
		err := decoder.Decode(&data)
		if err == nil {
			log.Printf("Text message sent (sid: %s)\n", data["sid"])
		}
		return err
	} else {
		log.Println(resp.Status)
	}

	return err
}