package pagerduty

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/davecgh/go-spew/spew"
)

type Client struct {
	http  *http.Client
	token string
}

func New(token string) *Client {
	tr := &http.Transport{
		TLSHandshakeTimeout: 5 * time.Second,
		IdleConnTimeout:     30 * time.Second,
	}
	httpClient := &http.Client{
		Transport: tr,
		Timeout:   10 * time.Second,
	}

	return &Client{http: httpClient, token: token}
}

func (c *Client) IsL1(email, scheduleID string) (bool, error) {
	req, err := http.NewRequest(http.MethodGet, "https://api.pagerduty.com/oncalls", nil)
	if err != nil {
		return false, err
	}

	req.Header.Add("Accept", "application/vnd.pagerduty+json;version=2")

	req.Header.Add("Authorization", fmt.Sprintf("Token token=%s", c.token))

	v := url.Values{}
	v.Set("include[]", "users")
	v.Set("time_zone", "UTC")
	v.Set("schedule_id", scheduleID)
	req.URL.RawQuery = v.Encode()
	resp, err := c.http.Do(req)

	defer func() {
		err = resp.Body.Close()
		if err != nil {
			log.Printf("%v", err)
		}
	}()

	if err != nil {
		return false, err
	}

	var data OnCalls

	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&data)

	spew.Dump(data)

	if err != nil {
		return false, err
	}
	for _, onCall := range data.OnCalls {
		if onCall.EscalationLevel > 1 {
			break
		}
		if onCall.User.Email == email {
			return true, nil
		}
	}
	return false, nil
}
