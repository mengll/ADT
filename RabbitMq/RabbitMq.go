// RabbitMq project RabbitMq.go
package RabbitMq

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

type RabbitMq struct {
	urls    string
	channel *amqp.Channel
}

//decaleTYpe

type DeclareType struct {
	name      string
	durable   bool
	unusedDel bool
	exclusive bool
	nowait    bool
	args      interface{}
}

//create publish data
type PublishDat struct {
	exchange  string
	routerkey string
	mandatory bool
	immediate bool
	body      string
	name      string //queue name
}

//create Consume datatype

type ConsumeType struct {
	queue     string
	consumer  string
	autoack   bool
	exclusive bool
	nolocal   bool
	nowait    bool
	args      interface{}
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

/*
*@author menglingliang
*@email 1633094010@qq.com
 */

func NewRabbitMq(ct string) *RabbitMq {
	return &RabbitMq{urls: ct}
}

//connect the rabbitserver
func (self *RabbitMq) connect() {
	conn, err := amqp.Dial(self.urls)
	defer conn.Close()
	failOnError(err, "rabbit connect error")

	channel, erra := conn.Channel()
	defer channel.Close()
	failOnError(erra)
	self.channel = channel //set the data
}

//create a new decear
func (self *RabbitMq) NewDeclare() *DeclareType {
	return &DeclareType{durable: false, unusedDel: false, exclusive: false, nowait: false, args: nil}
}

//create a

func (self *RabbitMq) NewQueueDeclare(dtype DeclareType) amqp.Queue {
	q, err := self.channel.QueueDeclare(dtype.qname, dtype.durable, dtype.unusedDel, dtype.exclusive, dtype.nowait, dtype.args)
	failOnError(err)
	return q
}

//create of publish data

func (self *RabbitMq) NewPublicDat(q amqp.Queue) *PublishDat {
	return &PublishDat{mandatory: false, immediate: false, name: q.Name}
}

//send the data to rabbit
func (self *RabbitMq) PublishTo(dat *PublishDat) {
	err := self.channel.Publish(dat.exchange, dat.name, dat.mandatory, dat.immediate, amqp.Publishing{ContentType: "text/plain", Body: []byte(dat.body)})
	failOnError(err, "publish data error")
}

//create new consuedata

func (self *RabbitMq) NewConsumeDat(q amqp.Queue) *ConsumeType {
	return &ConsumeType{queue: q.Name, consumer: false, autoack: true, exclusive: false, nolocal: false, nowait: false, args: nil}
}

//create consumeDat

func (self *RabbitMq) NewConsume(dat *ConsumeType) <-chan amqp.Delivery {
	msgs, err := self.channel.Consume(dat.queue, dat.consumer, dat.autoack, dat.exclusive, dat.nolocal, dat.nowait, dat.args)
	failOnError(err, "newConsue error")
	return msgs
}
