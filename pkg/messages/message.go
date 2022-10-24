package messages

type Message struct {
	Nickname string
	Text     string
}

func NewMessage(nick, text string) Message {
	return Message{
		Nickname: nick,
		Text:     text,
	}
}
