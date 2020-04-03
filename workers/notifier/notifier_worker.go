package notifier

import (
	"encoding/json"
	"fmt"
	"github.com/ThreeDotsLabs/watermill/message"
)

const Topic = "email.notify"

type Payload struct {
	Email string
	Link  string
}

func MessageHandler(msg *message.Message) {
	payload := Payload{}
	_ = json.Unmarshal(msg.Payload, &payload)

	fmt.Println("email sent")

}
