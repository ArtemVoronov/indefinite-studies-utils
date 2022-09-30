package kafka

const (
	EVENT_TYPE_SEND_EMAIL string = "SEND_EMAIL"
)

type SendEmailEvent struct {
	Sender    string `json:"Sender"`
	Recepient string `json:"Recepient"`
	Subject   string `json:"Subject"`
	Body      string `json:"Body"`
}

func GetPossibleEventTypes() []string {
	return []string{EVENT_TYPE_SEND_EMAIL}
}
