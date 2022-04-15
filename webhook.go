package gocord

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type Webhook struct {
	id             string
	type_webhook   int
	guild_id       string
	channel_id     string
	user           User
	name           string
	avatar         string
	token          string
	application_id string
	url            string
	source_guild   Guild
	source_channel Channel
}

func (c Client) createWebhook(w Webhook) *int {

	return nil
}

func (c Client) allWebhooksOfChannel(channelid string) []Webhook {
	var webhooks []Webhook
	// response, err := http.Get(c.api_url + c.api_version + "/channels/" + channel.id + "/webhooks")
	req, err := http.NewRequest("GET", c.api_url+c.api_version+"/channels/"+channelid+"/webhooks", nil) // Create a request
	req.Header.Set("Authorization", "Bot "+c.token)                                                     // Authorization for discord
	clientHttp := &http.Client{}                                                                        // Create http client for request
	response, err := clientHttp.Do(req)                                                                 // Get response from api
	if err != nil {                                                                                     // Error while sending response from api
		fmt.Println("BIG ERROR!!: Error while sending response to discord api")
		os.Exit(1)
	} else {
		defer response.Body.Close()

		contents, err := ioutil.ReadAll(response.Body)
		json.Unmarshal([]byte(contents), &webhooks)

		if err != nil {
			fmt.Println("BIG ERROR!!: Unknown error while getting webhooks of channel: " + channelid)
			os.Exit(1)
		}

		fmt.Println(webhooks)
		fmt.Println(string(contents))
	}
	return webhooks
}
