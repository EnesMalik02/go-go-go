package message

type Message struct {
	ID      int    `json:"id"`
	Content string `json:"content"`
	Sender  string `json:"sender"`
}
