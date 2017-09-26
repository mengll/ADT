package main

import (
	"ADT/RabbitMq"
	"log"
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

	da := rabbitmq.NewConsumeDat(q)
	forever := make(chan bool)
	msg := rabbitmq.NewConsume(da)

	go func() {
		for d := range msg {
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	<-forever

}
