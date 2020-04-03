package uploader

import (
	"encoding/json"
	"fmt"
	"github.com/ThreeDotsLabs/watermill/message"
	"watermill_training/workers"
	"watermill_training/workers/notifier"
)

const ProjectTopic = "project.upload"
const InvoiceTopic = "invoice.upload"
const ClientTopic = "client.upload"

type Payload struct {
	Email string
}

func ProjectHandler(msg *message.Message) {

	payload := Payload{}
	_ = json.Unmarshal(msg.Payload, &payload)

	projects := getProjectsByEmail(payload.Email)
	link := uploadProjectsToS3(projects)

	// project.uploaded event ---> email.notify handler
	workers.AMQPClient.Publisher(notifier.Topic, notifier.Payload{
		Email: payload.Email,
		Link:  link,
	})
}

func ClientHandler(msg *message.Message) {

	payload := Payload{}
	_ = json.Unmarshal(msg.Payload, &payload)

	clients := getClientsByEmail(payload.Email)
	link := uploadProjectsToS3(clients)

	// client.uploaded event ---> email.notify handler
	workers.AMQPClient.Publisher(notifier.Topic, notifier.Payload{
		Email: payload.Email,
		Link:  link,
	})
}

func InvoiceHandler(msg *message.Message) {

	payload := Payload{}
	_ = json.Unmarshal(msg.Payload, &payload)

	invoices := getInvoicesByEmail(payload.Email)
	link := uploadProjectsToS3(invoices)

	// invoice.uploaded event ---> email.notify handler
	workers.AMQPClient.Publisher(notifier.Topic, notifier.Payload{
		Email: payload.Email,
		Link:  link,
	})
}

func uploadProjectsToS3(things string) string {
	fmt.Println("file was uploaded")
	return "I am link"
}

func getProjectsByEmail(email string) string {
	return " I am Project List from MySQL database"
}

func getClientsByEmail(email string) string {
	return " I am Client List from MySQL database"
}

func getInvoicesByEmail(email string) string {
	return " I am Invoice List from MySQL database"
}
