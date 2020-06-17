package nsq
import (
	"fmt"
	"github.com/nsqio/go-nsq"
	"io/ioutil"
	"log"
	"sync"
)


type NSQHandler struct {
}

func (this *NSQHandler) HandleMessage(msg *nsq.Message) error {
	fmt.Println("receive", msg.NSQDAddress, "message:", string(msg.Body))
	return nil
}

const (
	TOPIC   = "test"
	CHANNEL_1 = "consumer_channel_1"
	CHANNEL_2 = "consumer_channel_2"
	URL     = "127.0.0.1:4160"
)

func testNSQ() {

	waiter := sync.WaitGroup{}
	waiter.Add(1)

	go func() {
		defer waiter.Done()

		config := nsq.NewConfig()
		config.MaxInFlight = 10

		for i := 0; i < 10; i++ {
			consumer, err := nsq.NewConsumer(TOPIC, CHANNEL_1, config)
			if nil != err {
				fmt.Println("err", err)
				return
			}
			consumer.SetLogger(log.New(ioutil.Discard, "", log.LstdFlags), nsq.LogLevelInfo)
			consumer.AddHandler(&NSQHandler{})
			err = consumer.ConnectToNSQD(URL)
			if nil != err {
				fmt.Println("err", err)
				return
			}

			fmt.Println(CHANNEL_1,i)
		}
		select {}
	}()


	waiter.Wait()
}