package main

import (
	"ADT/RabbitMq"
)

func main() {
	//1 shuchushua
	rabbitmq := RabbitMq.NewRabbitMq("amqp://guest:guest@localhost:5672/")
	//2 connect
	rabbitmq.Connect()
	//3 create declaretype
	qudeclare := rabbitmq.NewDeclareType()
	qudeclare.Name = "fjk"
	//4 new declare dat
	q := rabbitmq.NewQueueDeclare(qudeclare)

	pdat := rabbitmq.NewPublicDat(q)
	pdat.Body = "fjsa" // the data message!
	pdat.Name = q.Name
	pdat.Exchange = ""

	rabbitmq.PublishTo(pdat)
}
