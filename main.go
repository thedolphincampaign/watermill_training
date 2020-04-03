package main

import (
	"watermill_training/workers"
	"watermill_training/workers/notifier"
	"watermill_training/workers/uploader"
)

func main() {

	amqpClient := workers.New()
	amqpClient.Subscriber(uploader.ProjectTopic, uploader.ProjectHandler)
	amqpClient.Subscriber(uploader.ClientTopic, uploader.ClientHandler)
	amqpClient.Subscriber(uploader.InvoiceTopic, uploader.InvoiceHandler)

	amqpClient.Subscriber(notifier.Topic, notifier.MessageHandler)
	amqpClient.Subscriber(notifier.Topic, notifier.MessageHandler)
	amqpClient.Subscriber(notifier.Topic, notifier.MessageHandler)

	//User just click a "Export Projects" button for exporting projects
	// exportButtonClicked event ---> project.upload handler (upload projects to S3 amazon)
	amqpClient.Publisher(uploader.ProjectTopic, uploader.Payload{
		Email: "user1@gmail.com",
	})

	//User just click a "Export Clients" button for exporting clients
	// exportButtonClicked event ---> client.upload handler (upload projects to S3 amazon)
	amqpClient.Publisher(uploader.ClientTopic, uploader.Payload{
		Email: "user1@gmail.com",
	})

	//User just click a "Export Invoices" button for exporting clients
	// exportButtonClicked event ---> invoice.upload handler (upload projects to S3 amazon)
	amqpClient.Publisher(uploader.InvoiceTopic, uploader.Payload{
		Email: "user1@gmail.com",
	})
}
