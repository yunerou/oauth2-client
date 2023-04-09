package rocketchat

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/yunerou/oauth2-client/singleton"
)

func (p *chatProvider) send(msg rocketChatMessage) error {
	buffer, err := json.Marshal(msg)

	if err != nil {
		return err
	}

	response, err := http.Post(
		p.config.WebhookURL,
		"application/json",
		bytes.NewReader(buffer),
	)

	if err != nil {
		return err
	}

	defer response.Body.Close()
	if response.StatusCode != 200 {
		return fmt.Errorf("send error: %d", response.StatusCode)
	}

	return nil
}

func (p *chatProvider) SendText(text string, mentionTo ...string) error {
	mentionStr := ""
	for _, who := range mentionTo {
		mentionStr += " 🧊 " + "@" + who
	}
	sendMsg := fmt.Sprintf(
		"%s \n- 🥥Enviroment **%s** \t 🥭Version **%s**\n---\n%s",
		mentionStr,
		singleton.GetViper().GetString("ENV"),
		singleton.GetViper().GetString("VERSION"),
		text)
	return p.send(rocketChatMessage{
		Text:    sendMsg,
		Emoji:   p.config.Emoji,
		Channel: p.config.Channel,
	})
}

// Message to send
type rocketChatMessage struct {
	Text    string `json:"text"`
	Emoji   string `json:"emoji"`
	Channel string `json:"channel,omitempty"`
}
