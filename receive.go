package main

import (
	"ADT/RabbitMq"
	"fmt"
	"log"
)

func maint() {

	//1 create all
	ch := RabbitMq.NewRabbitMq("amqp://guest:guest@127.0.0.1:5672/")
	//2 connect
	ch.Connect()
	//close
	//defer ch.Closeq()

	//3 crewate exchangetype
	exdata := ch.NewExchangeDecType()
	exdata.Name = "logs_direct"
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

	ch.QueueBind(q.Name, "black_meng", exdata.Name)

	//8 consume dat

	consuedat := ch.NewConsumeDatType(q)
	consuedat.Autoack = true

	consuedat.Queue = q.Name
	consuedat.Consumer = ""

	//9
	msg := ch.NewConsume(consuedat)

	forever := make(chan bool)

	go func() {
		for d := range msg {
			log.Printf(" [x] %s", d.Body)
			//d.Ack(true) shoudongquerenshoudao,feizen shibai fafougei geibie de jianceshi chuli
		}
	}()

	log.Printf(" [*] Waiting for logs. To exit press CTRL+C")
	<-forever

}
