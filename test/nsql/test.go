package main

//import (
//	"github.com/nsqio/go-nsq"
//	"fmt"
//	"time"
//)
//
//var producer *nsq.Producer
//var consumer *nsq.Consumer
//
//type t struct {
//
//}
//
//func (this *t)HandleMessage(message *nsq.Message) error{
//	fmt.Println(string(message.Body))
//	return nil
//}
//func init()  {
//
//	producer,_ = nsq.NewProducer("127.0.0.1:4150",nsq.NewConfig()) // 初始化生产者
//	//consumer,_ = nsq.NewConsumer("fwd","ch1",nsq.NewConfig())
//	//consumer.AddHandler(new(t))
//	//consumer.ConnectToNSQD("127.0.0.1:4150")
//}
//func main() {
//
//
//	go func() {
//
//	}()
//	for{
//		time.Sleep(time.Second)
//		producer.Publish("test",[]byte("哈哈哈哈哈"))
//	}
//
//}