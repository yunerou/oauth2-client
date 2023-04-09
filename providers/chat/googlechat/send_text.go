package googlechat

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/yunerou/oauth2-client/singleton"
)

func (p *googlechatProvider) send(msg ggChatMessage) error {

	buffer, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	response, err := http.Post(
		p.config.WebhookURL,
		"application/json; charset=UTF-8",
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

func (p *googlechatProvider) SendText(text string, mentionTo ...string) error {

	mentionStr := ""
	for _, who := range mentionTo {
		mentionStr += " 🧊 " + who
	}
	sendMsg := fmt.Sprintf(
		"%s \n- 🥥Enviroment *%s* \t 🥭Version *%s*\n---\n%s",
		mentionStr,
		singleton.GetViper().GetString("ENV"),
		singleton.GetViper().GetString("VERSION"),
		text)
	return p.send(ggChatMessage{
		Text: sendMsg,
	})
}

// Message to send.
type ggChatMessage struct {
	Text string `json:"text"`
}
