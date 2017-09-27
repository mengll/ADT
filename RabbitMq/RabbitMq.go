// RabbitMq project RabbitMq.go
/*
*@author menglingliang
*@email 1633094010@qq.com
*@describ go rabbitmq manager lib not include rpc manager
 */
package RabbitMq

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

type RabbitMq struct {
	urls    string
	Channel *amqp.Channel
	Conn    *amqp.Connection
}

//decaleTYpe

type DeclareType struct {
	Name      string
	Durable   bool
	UnusedDel bool
	Exclusive bool
	Nowait    bool
	Args      amqp.Table
}

//create publish data
type PublishDat struct {
	Exchange  string
	Routerkey string
	Mandatory bool
	Immediate bool
	Body      string
	Name      string //queue name
}

//create Consume datatype

type ConsumeType struct {
	Queue     string
	Consumer  string
	Autoack   bool
	Exclusive bool
	Nolocal   bool
	Nowait    bool
	Args      amqp.Table
}

//create QOS

type QosType struct {
	PrefetchCount int
	PrefetchSize  int
	Global        bool
}

//exchange data type
type ExchangeDeclareType struct {
	Name     string
	Type     string
	Durable  bool
	AutoDel  bool
	Internal bool
	Nowait   bool
	Args     amqp.Table
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

func NewRabbitMq(ct string) *RabbitMq {
	return &RabbitMq{urls: ct}
}

func (self *RabbitMq) Closeq() {
	self.Channel.Close()
}

//connect the rabbitserver
func (self *RabbitMq) Connect() {
	conn, err := amqp.Dial(self.urls)
	self.Conn = conn
	failOnError(err, "rabbit connect error")

	channel, erra := conn.Channel()

	failOnError(erra, "oso")
	self.Channel = channel //set the data
}

//create a new decear
func (self *RabbitMq) NewDeclareType() *DeclareType {
	return &DeclareType{Durable: false, UnusedDel: false, Exclusive: false, Nowait: false}
}

//create a

func (self *RabbitMq) NewQueueDeclare(dtype *DeclareType) amqp.Queue {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("cuowu", dtype)
			fmt.Println(self.Channel)
		}

	}()

	q, err := self.Channel.QueueDeclare(dtype.Name, dtype.Durable, dtype.UnusedDel, dtype.Exclusive, dtype.Nowait, dtype.Args)
	failOnError(err, "e")
	return q
}

//create of publish data

func (self *RabbitMq) NewPublicDatType(q amqp.Queue) *PublishDat {
	return &PublishDat{Mandatory: false, Immediate: false, Name: q.Name}
}

//send the data to rabbit
func (self *RabbitMq) PublishTo(dat *PublishDat) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("one tow one ")
		}
	}()

	err := self.Channel.Publish(dat.Exchange, dat.Name, dat.Mandatory, dat.Immediate, amqp.Publishing{DeliveryMode: amqp.Persistent, ContentType: "text/plain", Body: []byte(dat.Body)})
	failOnError(err, "publish data error")
}

//create new consuedata

func (self *RabbitMq) NewConsumeDatType(q amqp.Queue) *ConsumeType {
	return &ConsumeType{Queue: q.Name, Autoack: true, Exclusive: false, Nolocal: false, Nowait: false, Args: nil}
}

//create consumeDat

func (self *RabbitMq) NewConsume(dat *ConsumeType) <-chan amqp.Delivery {
	msgs, err := self.Channel.Consume(dat.Queue, dat.Consumer, dat.Autoack, dat.Exclusive, dat.Nolocal, dat.Nowait, nil)
	failOnError(err, "newConsue error")
	return msgs
}

//set the Qos size
func (self *RabbitMq) SetQos(dat *QosType) {
	err := self.Channel.Qos(dat.PrefetchCount, dat.PrefetchSize, dat.Global)
	failOnError(err, "Set Qos size error")
}

//create a dat type
func (self *RabbitMq) NewExchangeDecType() *ExchangeDeclareType {
	return &ExchangeDeclareType{Durable: false, AutoDel: false, Internal: false, Nowait: false}
}

//create queuebind
func (self *RabbitMq) QueueBind(queuename, routerkey, exchange string) {
	err := self.Channel.QueueBind(queuename, routerkey, exchange, false, nil)
}
