package main

import (
	"ADT/RabbitMq"
)

func main() {
	//1 create all
	ch := RabbitMq.NewRabbitMq("amqp://guest:guest@127.0.0.1:5672/")
	//2 connect
	ch.Connect()

	//close
	defer ch.Closeq()

	//3 crewate exchangetype
	exdata := ch.NewExchangeDecType()
	exdata.Name = "logs_direct"
	exdata.Type = "direct"
	exdata.Durable = true

	//4
	ch.NewExchangeDec(exdata)

	//5

	pusdat := ch.NewPublicDatType()

	pusdat.Body = "mengll"
	pusdat.Routerkey = "black_meng"
	pusdat.Mandatory = false
	pusdat.Immediate = false
	pusdat.Exchange = exdata.Name

	ch.PublishTo(pusdat)
	//	ch.Channel.Publish(
	//		"logs_direct", // exchange
	//		"black_meng",  // routing key
	//		false,         // mandatory
	//		false,         // immediate
	//		amqp.Publishing{
	//			ContentType: "text/plain",
	//			Body:        []byte("mengllllll"),
	//		})

}
