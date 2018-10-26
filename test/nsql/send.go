//Nsq发送测试
package main

import (

	"fmt"
	"github.com/nsqio/go-nsq"
)

var producer *nsq.Producer

// 主函数12
func main() {
	strIP1 := "127.0.0.1:4150"
	//strIP2 := "127.0.0.1:4152"
	InitProducer(strIP1)

	for i:=0;i<10;i++ {


		for err := Publish("test","value:"+string(i) ); err != nil; err = Publish("test", "value:"+string(i)) {
			//切换IP重连
			//strIP1, strIP2 = strIP2, strIP1
			//InitProducer(strIP1)
		}
	}
}

// 初始化生产者
func InitProducer(str string) {
	var err error
	fmt.Println("address: ", str)
	producer, err = nsq.NewProducer(str, nsq.NewConfig())
	if err != nil {
		panic(err)
	}
}

//发布消息
func Publish(topic string, message string) error {
	var err error
	if producer != nil {
		if message == "" { //不能发布空串，否则会导致error
			return nil
		}
		err = producer.Publish(topic, []byte(message)) // 发布消息
		return err
	}
	return fmt.Errorf("producer is nil", err)
}