package rocketchat

import "github.com/yunerou/oauth2-client/providers/chat"

type RocketchatConfig struct {
	WebhookURL string
	Emoji      string
	Channel    string
}

type chatProvider struct {
	config *RocketchatConfig
}

func NewChatProvider(config *RocketchatConfig) chat.ChatProvider {
	chatInstance := &chatProvider{
		config: config,
	}
	return chatInstance
}
