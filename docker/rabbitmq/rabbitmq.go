package rabbitmq

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

//url 格式 amqp://帐号:密码@rabbitmq服务器地址：端口/vhost
const MQURL = "amqp://harrylee:harrylee@127.0.0.1:5672/chain33"

type RabbitMQ struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	//队列名称
	QueueName string
	//交换机
	Exchange string
	//key
	Key string
	//连接信息
	Mqurl string
}

//创建RabbitMQ结构体实例
func NewRabbitMQ(queueName string, exchange string, key string) *RabbitMQ {
	rabbitmq := &RabbitMQ{QueueName: queueName, Exchange: exchange, Key: key, Mqurl: MQURL}
	var err error
	//创建rabbitmq conn
	rabbitmq.conn, err = amqp.Dial(rabbitmq.Mqurl)
	rabbitmq.failOnErr(err, "failed to connect rabbitmq!")
	rabbitmq.channel, err = rabbitmq.conn.Channel()
	rabbitmq.failOnErr(err, "failed to open a channel!")
	return rabbitmq
}

//Close
func (r *RabbitMQ) Destrory() {
	r.channel.Close()
	r.conn.Close()
}

//错误处理
func (r *RabbitMQ) failOnErr(err error, message string) {
	if err != nil {
		log.Fatalf("%s:%s", message, err)
		panic(fmt.Sprintf("%s:%s", message, err))
	}
}

//简单模式
func NewRabbitMQSimple(queueName string) *RabbitMQ {
	rabbitmq := NewRabbitMQ(queueName, "", "")
	return rabbitmq
}

//简单模式下生产消息
func (r *RabbitMQ) PublishSimple(message string) {
	//1.申请队列，如果队列不存在会自动创建
	_, err := r.channel.QueueDeclare(
		r.QueueName,
		//是否持久化
		false,
		//是否自动删除
		false,
		//是否具有排他性
		false,
		//是否阻塞
		false,
		//额外属性
		nil,
	)
	if err != nil {
		fmt.Println(err)
	}

	//2.发送消息到队列中
	r.channel.Publish(
		r.Exchange,
		r.QueueName,
		//如果为true，根据exchange类型和routekey 规则，如果无法找到符合条件得队列，则会把发送得消息返回给发送者
		false,
		//如果为true，当exchange发送消息到队列后发现队列上没有绑定消费者，则会把发送得消息返还给发送者
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})
}

func (r *RabbitMQ) ConsumeSimple() {
	//1.申请队列，如果队列不存在会自动创建
	_, err := r.channel.QueueDeclare(
		r.QueueName,
		//是否持久化
		false,
		//是否自动删除
		false,
		//是否具有排他性
		false,
		//是否阻塞
		false,
		//额外属性
		nil,
	)

	if err != nil {
		fmt.Println(err)
	}
	//接收消息
	msgs, err := r.channel.Consume(
		r.QueueName,
		//用来区分多个消费者
		"",
		//是否设置自动应答
		true,
		//是否具有排他性
		false,
		//如果设置为true,表示不能将同一个connection中发送得消息传递给这个connection中得消费者
		false,
		false,
		nil,
	)
	if err != nil {
		fmt.Println(err)
	}

	forever := make(chan bool)
	go func() {
		for d := range msgs {
			//TODO 处理函数
			log.Printf("Received amessage: %s", d.Body)
		}
	}()
	log.Printf("[*] waiting for messages, to exit press ctrl + c")
	<-forever
}

//订阅模式
func NewRabbitMQPubSub(exchangeName string) *RabbitMQ {
	rabbitmq := NewRabbitMQ("", exchangeName, "")
	return rabbitmq
}

//订阅模式生产
func (r *RabbitMQ) PublishPub(message string) {
	//1.尝试创建交换机
	err := r.channel.ExchangeDeclare(
		r.Exchange,
		//交换机类型
		"fanout",
		//是否持久化
		true,
		false,
		//true表示这个exchange不可以被client用来推送消息，仅用来进行exchange和exchange之间得绑定
		false,
		false,
		nil,
	)
	if err != nil {
		fmt.Println(err)
	}

	//2.发送消息到队列中
	r.channel.Publish(
		r.Exchange,
		"",
		//如果为true，根据exchange类型和routekey 规则，如果无法找到符合条件得队列，则会把发送得消息返回给发送者
		false,
		//如果为true，当exchange发送消息到队列后发现队列上没有绑定消费者，则会把发送得消息返还给发送者
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})
}

//订阅模式消费
func (r *RabbitMQ) ReceivSub() {
	//1.尝试创建交换机
	err := r.channel.ExchangeDeclare(
		r.Exchange,
		//交换机类型
		"fanout",
		//是否持久化
		true,
		false,
		//true表示这个exchange不可以被client用来推送消息，仅用来进行exchange和exchange之间得绑定
		false,
		false,
		nil,
	)
	if err != nil {
		r.failOnErr(err, "failed to declare an exchange")
	}

	//2.试探性创建队列
	q, err := r.channel.QueueDeclare(
		//随机生成队列名称
		"",
		//如果为true，根据exchange类型和routekey 规则，如果无法找到符合条件得队列，则会把发送得消息返回给发送者
		false,
		//如果为true，当exchange发送消息到队列后发现队列上没有绑定消费者，则会把发送得消息返还给发送者
		false,
		true,
		false,
		nil,
	)
	if err != nil {
		r.failOnErr(err, "failed to declare an queue")
	}
	//绑定队列到exchange中
	err = r.channel.QueueBind(
		q.Name,
		//在pub/sub模式下这里的key必须为空
		"",
		r.Exchange,
		false,
		nil)
	if err != nil {
		r.failOnErr(err, "failed to queuebind")
	}
	//接收消息
	msgs, err := r.channel.Consume(
		q.Name,
		//用来区分多个消费者
		"",
		//是否设置自动应答
		true,
		//是否具有排他性
		false,
		//如果设置为true,表示不能将同一个connection中发送得消息传递给这个connection中得消费者
		false,
		false,
		nil,
	)
	if err != nil {
		fmt.Println(err)
	}

	forever := make(chan bool)
	go func() {
		for d := range msgs {
			//TODO 处理函数
			log.Printf("Received amessage: %s", d.Body)
		}
	}()
	log.Printf("[*] waiting for messages, to exit press ctrl + c")
	<-forever
}


//路由模式
func NewRabbitMQRouting(exchangeName,routingKey string) *RabbitMQ {
	rabbitmq := NewRabbitMQ("", exchangeName, routingKey)
	return rabbitmq
}

func (r *RabbitMQ) PublishRouting(message string) {
	//1.尝试创建交换机
	err := r.channel.ExchangeDeclare(
		r.Exchange,
		//交换机类型
		"direct",
		//是否持久化
		true,
		false,
		//true表示这个exchange不可以被client用来推送消息，仅用来进行exchange和exchange之间得绑定
		false,
		false,
		nil,
	)
	if err != nil {
		fmt.Println(err)
	}

	//2.发送消息到队列中
	r.channel.Publish(
		r.Exchange,
		r.Key,
		//如果为true，根据exchange类型和routekey 规则，如果无法找到符合条件得队列，则会把发送得消息返回给发送者
		false,
		//如果为true，当exchange发送消息到队列后发现队列上没有绑定消费者，则会把发送得消息返还给发送者
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})
}

func (r *RabbitMQ) ReceiveRouting() {
	//1.尝试创建交换机
	err := r.channel.ExchangeDeclare(
		r.Exchange,
		//交换机类型
		"direct",
		//是否持久化
		true,
		false,
		//true表示这个exchange不可以被client用来推送消息，仅用来进行exchange和exchange之间得绑定
		false,
		false,
		nil,
	)
	if err != nil {
		r.failOnErr(err, "failed to declare an exchange")
	}

	//2.试探性创建队列
	q, err := r.channel.QueueDeclare(
		//随机生成队列名称
		"",
		//如果为true，根据exchange类型和routekey 规则，如果无法找到符合条件得队列，则会把发送得消息返回给发送者
		false,
		//如果为true，当exchange发送消息到队列后发现队列上没有绑定消费者，则会把发送得消息返还给发送者
		false,
		true,
		false,
		nil,
	)
	if err != nil {
		r.failOnErr(err, "failed to declare an queue")
	}
	//绑定队列到exchange中
	err = r.channel.QueueBind(
		q.Name,
		//在Routing模式下指定key
		r.Key,
		r.Exchange,
		false,
		nil)
	if err != nil {
		r.failOnErr(err, "failed to queuebind")
	}
	//接收消息
	msgs, err := r.channel.Consume(
		q.Name,
		//用来区分多个消费者
		"",
		//是否设置自动应答
		true,
		//是否具有排他性
		false,
		//如果设置为true,表示不能将同一个connection中发送得消息传递给这个connection中得消费者
		false,
		false,
		nil,
	)
	if err != nil {
		fmt.Println(err)
	}

	forever := make(chan bool)
	go func() {
		for d := range msgs {
			//TODO 处理函数
			log.Printf("Received amessage: %s", d.Body)
		}
	}()
	log.Printf("[*] waiting for messages, to exit press ctrl + c")
	<-forever
}

//话题模式
func NewRabbitMQTopic(exchangeName,routingKey string) *RabbitMQ {
	rabbitmq := NewRabbitMQ("", exchangeName, routingKey)
	return rabbitmq
}

func (r *RabbitMQ) PublishTopic(message string) {
	//1.尝试创建交换机
	err := r.channel.ExchangeDeclare(
		r.Exchange,
		//交换机类型
		"topic",
		//是否持久化
		true,
		false,
		//true表示这个exchange不可以被client用来推送消息，仅用来进行exchange和exchange之间得绑定
		false,
		false,
		nil,
	)
	if err != nil {
		fmt.Println(err)
	}

	//2.发送消息到队列中
	r.channel.Publish(
		r.Exchange,
		r.Key,
		//如果为true，根据exchange类型和routekey 规则，如果无法找到符合条件得队列，则会把发送得消息返回给发送者
		false,
		//如果为true，当exchange发送消息到队列后发现队列上没有绑定消费者，则会把发送得消息返还给发送者
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})
}

//话题模式接收消息
//要注意key,规则
//其中 * 用于匹配一个单词， # 用于匹配多个单词（可以是零个）
//chain33.hello  => chain33.*
//chain33.hello.world => chain33.#
func (r *RabbitMQ) ReceiveTopic() {
	//1.尝试创建交换机
	err := r.channel.ExchangeDeclare(
		r.Exchange,
		//交换机类型
		"topic",
		//是否持久化
		true,
		false,
		//true表示这个exchange不可以被client用来推送消息，仅用来进行exchange和exchange之间得绑定
		false,
		false,
		nil,
	)
	if err != nil {
		r.failOnErr(err, "failed to declare an exchange")
	}

	//2.试探性创建队列
	q, err := r.channel.QueueDeclare(
		//随机生成队列名称
		"",
		//如果为true，根据exchange类型和routekey 规则，如果无法找到符合条件得队列，则会把发送得消息返回给发送者
		false,
		//如果为true，当exchange发送消息到队列后发现队列上没有绑定消费者，则会把发送得消息返还给发送者
		false,
		true,
		false,
		nil,
	)
	if err != nil {
		r.failOnErr(err, "failed to declare an queue")
	}
	//绑定队列到exchange中
	err = r.channel.QueueBind(
		q.Name,
		//在Routing模式下指定key
		r.Key,
		r.Exchange,
		false,
		nil)
	if err != nil {
		r.failOnErr(err, "failed to queuebind")
	}
	//接收消息
	msgs, err := r.channel.Consume(
		q.Name,
		//用来区分多个消费者
		"",
		//是否设置自动应答
		true,
		//是否具有排他性
		false,
		//如果设置为true,表示不能将同一个connection中发送得消息传递给这个connection中得消费者
		false,
		false,
		nil,
	)
	if err != nil {
		fmt.Println(err)
	}

	forever := make(chan bool)
	go func() {
		for d := range msgs {
			//TODO 处理函数
			log.Printf("Received amessage: %s", d.Body)
		}
	}()
	log.Printf("[*] waiting for messages, to exit press ctrl + c")
	<-forever
}