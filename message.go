package hipchat2

type Message struct {
	Color string `json:"color"`
	Message string `json:"message"`
	Notify bool `json:"notify"`
	MessageFormat string `json:"message_format"`
}
