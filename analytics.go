package analytics

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

var (
	secret      string
	measurement string
	debug       bool
)

type Payload struct {
	ClientID           string            `json:"client_id"`
	UserID             string            `json:"user_id,omitempty"`
	TimestampMicros    int               `json:"timestamp_micros,omitempty"`
	UserProperties     map[string]string `json:"user_properties,omitempty"`
	NonPersonalizedAds bool              `json:"non_personalized_ads,omitempty"`
	Events             []Event           `json:"events"`
}

type Event struct {
	Name   string                 `json:"name"`
	Params map[string]interface{} `json:"params"`
}

func SetKeys(apiSecret, measurementID string) {
	secret = apiSecret
	measurement = measurementID
}

func Debug(b bool) {
	debug = b
}

func Send(payload Payload) {
	query := url.Values{}
	query.Add("api_secret", secret)
	query.Add("measurement_id", measurement)
	uri := fmt.Sprintf("https://www.google-analytics.com/mp/collect?%s", query.Encode())
	if debug {
		uri = fmt.Sprintf("https://www.google-analytics.com/debug/mp/collect?%s", query.Encode())
	}
	for i, event := range payload.Events {
		if event.Params == nil {
			event.Params = map[string]interface{}{}
			payload.Events[i] = event
		}
	}
	bs, err := json.Marshal(payload)
	if debug && err != nil {
		log.Println(err)
		return
	}
	if debug {
		fmt.Println("url: ", uri)
		fmt.Println("payload: ", string(bs))
	}
	body := bytes.NewReader(bs)
	res, err := http.Post(uri, "application/json", body)
	if debug && err != nil {
		log.Println(err)
		return
	}
	if !debug {
		return
	}
	bs, err = ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(string(bs))
}
