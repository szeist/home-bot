package events

type Message struct {
	Text   string
	Sender interface{}
}

type Response struct {
	Text   string
	Sender interface{}
	Error  error
}
