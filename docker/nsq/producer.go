package nsq

import (
	"fmt"
	"github.com/nsqio/go-nsq"
	"io/ioutil"
	"log"
	"sync"
	"time"
)

var err error

// 推送消息
func testproducer() {

	url := "127.0.0.1:4160"
	topicName := TOPIC
	config := nsq.NewConfig()

	// new
	producer, err := nsq.NewProducer(url, config)
	if err != nil {
		fmt.Println("nsq.NewProducer", err)
		return
	}
	defer producer.Stop()

	producer.SetLogger(log.New(ioutil.Discard, "", log.LstdFlags), nsq.LogLevelInfo)
	//  ping
	err = producer.Ping()
	if err != nil {
		fmt.Println("producer.Ping", err)
		return
	}
	fmt.Println("producer.Ping", "√")

	msgCt := 1000
	wg := &sync.WaitGroup{}
	wg.Add(msgCt)
	// 测试10 次
	for i := 0; i < msgCt; i++ {

		// 消息内容
		msg := time.Now().Format("0102150405")
		sendMessage(producer, topicName, msg)
		wg.Done()

		time.Sleep(10 * time.Millisecond)
		// time.Sleep(1 * time.Second)
	}

	wg.Wait()
	fmt.Println("producer.Push.Status", "ok")
}

// 发送消息
func sendMessage(producer *nsq.Producer, topicName string, msg string) {

	err = producer.Publish(topicName, []byte(msg))
	if err != nil {
		fmt.Println("producer.Publish", err)
		return
	}
	fmt.Println("producer.Publish", msg, "√")

}

