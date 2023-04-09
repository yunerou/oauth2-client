package googlechat

import (
	"github.com/yunerou/oauth2-client/providers/chat"
)

type GooglechatConfig struct {
	WebhookURL string
}

type googlechatProvider struct {
	config *GooglechatConfig
}

func NewChatProvider(config *GooglechatConfig) chat.ChatProvider {
	chatInstance := &googlechatProvider{
		config: config,
	}
	return chatInstance
}
