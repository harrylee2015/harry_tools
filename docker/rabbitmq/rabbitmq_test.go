package rabbitmq

import (
	"strconv"
	"testing"
	"time"
)

func TestNewRabbitMQSimple(t *testing.T) {

	rabbitmq := NewRabbitMQSimple("simple")
	for i := 0; i <= 100; i++ {
		rabbitmq.PublishSimple("Hello World!" + strconv.Itoa(i))
	}
	for i := 0; i < 10; i++ {
		go func() {
			rabbitmqs := NewRabbitMQSimple("simple")
			rabbitmqs.ConsumeSimple()
		}()
	}
	time.Sleep(5 * time.Second)
}

func TestNewRabbitMQPubSub(t *testing.T) {
	for i := 0; i < 2; i++ {
		go func() {
			rabbitmq := NewRabbitMQPubSub("sub")
			rabbitmq.ReceivSub()
		}()
	}
	time.Sleep(time.Second)
	rabbitmq := NewRabbitMQPubSub("sub")
	for i := 0; i <= 10; i++ {
		rabbitmq.PublishPub("Hello World!" + strconv.Itoa(i))
	}
	time.Sleep(5 * time.Second)
}

func TestNewRabbitMQRouting(t *testing.T) {
	for i := 0; i < 2; i++ {
		go func(a int) {
			rabbitmq := NewRabbitMQRouting("route", "key"+strconv.Itoa(a))
			rabbitmq.ReceiveRouting()
		}(i)
	}
	time.Sleep(time.Second)
	rabbitmq := NewRabbitMQRouting("route", "key1")
	for i := 0; i <= 10; i++ {
		rabbitmq.PublishRouting("Hello World!" + strconv.Itoa(i))
	}
	time.Sleep(5 * time.Second)
}

func TestNewRabbitMQTopic(t *testing.T) {

	go func() {
		rabbitmq := NewRabbitMQTopic("topic", "key.*.1")
		rabbitmq.ReceiveTopic()
	}()
	go func() {
		rabbitmq := NewRabbitMQTopic("topic", "key.#")
		rabbitmq.ReceiveTopic()
	}()
	time.Sleep(time.Second)
	rabbitmq := NewRabbitMQTopic("topic", "key.chain33.1")
	for i := 0; i <= 10; i++ {
		rabbitmq.PublishTopic("Hello World!" + strconv.Itoa(i))
	}
	time.Sleep(5 * time.Second)
}

func TestNewRabbitMQPRC(t *testing.T) {

	go func() {
		rabbitmq := NewRabbitMQPRC("rpc")
		rabbitmq.ServerPRC()
	}()
	time.Sleep(time.Second)
	for i := 0; i <= 10; i++ {
		rabbitmq := NewRabbitMQPRC("rpc")
		rabbitmq.ClientRPC(strconv.Itoa(i))
	}
	time.Sleep(5 * time.Second)
}
