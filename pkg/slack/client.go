package slack

import (
	"log"

	api "github.com/nlopes/slack"
)

type Client struct {
	api *api.Client
}

func New(token string) *Client {
	api := api.New(token)
	return &Client{api: api}
}

func (c *Client) Send(msg *MessageInput) error {
	params := api.PostMessageParameters{}
	params.LinkNames = 1
	params.IconURL = "https://a.slack-edge.com/7f1a0/plugins/hubot/assets/service_512.png"
	chanID, ts, err := c.api.PostMessage(msg.Channel, msg.Message, params)
	if err != nil {
		log.Printf("Got error %v", err)
		return err
	}
	log.Printf("Sent message to %s at %s", chanID, ts)
	return nil
}
