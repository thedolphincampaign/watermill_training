package workers

import (
	"context"
	"encoding/json"
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-amqp/pkg/amqp"
	"github.com/ThreeDotsLabs/watermill/message"
)

var AMQPClient *AmqpClient

type AmqpClient struct {
	Config amqp.Config
}

func New() *AmqpClient {
	AMQPClient = &AmqpClient{Config: amqp.NewDurableQueueConfig("amqp://guest:guest@localhost:5672/")}
	return AMQPClient
}

func (c *AmqpClient) Subscriber(topic string, messageHandler func(*message.Message)) {
	subscriber, err := amqp.NewSubscriber(
		c.Config,
		watermill.NewStdLogger(false, false),
	)
	if err != nil {
		panic(err)
	}

	messages, err := subscriber.Subscribe(context.Background(), topic)
	if err != nil {
		panic(err)
	}

	go process(messages, messageHandler)

}

func process(messages <-chan *message.Message, messageHandler func(*message.Message)) {
	for msg := range messages {
		messageHandler(msg)
		msg.Ack()
	}
}

func (c *AmqpClient) Publisher(topic string, payload interface{}) {

	publisher, err := amqp.NewPublisher(c.Config, watermill.NewStdLogger(false, false))
	if err != nil {
		panic(err)
	}

	publishMessages(publisher, topic, payload)
}

func publishMessages(publisher message.Publisher, topic string, payload interface{}) {
	payloadJSON, _ := json.Marshal(payload)
	msg := message.NewMessage(watermill.NewUUID(), payloadJSON)

	if err := publisher.Publish(topic, msg); err != nil {
		panic(err)
	}
}
