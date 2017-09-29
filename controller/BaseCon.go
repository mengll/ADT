package controller

import (
	"ADT/RabbitMq"
	"fmt"

	"github.com/streadway/amqp"
)

func NewRecive(routerkey string) (rbq *RabbitMq.RabbitMq, mch <-chan amqp.Delivery) {
	ch := RabbitMq.NewRabbitMq("amqp://guest:guest@127.0.0.1:5672/")
	//2 connect
	ch.Connect()

	//3 crewate exchangetype
	exdata := ch.NewExchangeDecType()
	exdata.Name = "adt"
	exdata.Type = "direct"
	exdata.Durable = true

	//4
	ch.NewExchangeDec(exdata)

	//5 queue data

	queuedat := ch.NewDeclareType()
	queuedat.Name = ""
	queuedat.Exclusive = true

	//6
	q := ch.NewQueueDeclare(queuedat)

	//7 binddata
	fmt.Println(q.Name)

	ch.QueueBind(q.Name, routerkey, exdata.Name)

	//8 consume dat

	consuedat := ch.NewConsumeDatType(q)
	consuedat.Autoack = true

	consuedat.Queue = q.Name
	consuedat.Consumer = ""

	//9
	msg := ch.NewConsume(consuedat)
	return ch, msg
}
