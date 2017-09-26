package main

import (
	"ADT/RabbitMq"
	"fmt"
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
	fmt.Println(q)
	fmt.Println("-------------------->")
	pdat := rabbitmq.NewPublicDat(q)

	pdat.Body = "fjsa"
	pdat.Name = q.Name
	pdat.Exchange = "logs"

	rabbitmq.PublishTo(pdat)

}
