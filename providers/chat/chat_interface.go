package chat

type ChatProvider interface {
	SendText(text string, mentionTo ...string) error
}
