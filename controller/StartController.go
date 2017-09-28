package controller

import (
	"ADT/RabbitMq"
	"fmt"
)

func InitRecivem() {
	kl := []string{"start", "register", "payment"}
	//controller.InitRecivem()
	ml := make(chan bool)
	for _, item := range kl {
		go StartRecive(item)
	}
	<-ml

}

func StartRecive(routerkey string) {
	ch := RabbitMq.NewRabbitMq("amqp://guest:guest@127.0.0.1:5672/")
	//2 connect
	ch.Connect()

	//close
	defer ch.Closeq()
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

	forever := make(chan bool)
	defer close(forever)

	go func() {
		for d := range msg {
			switch routerkey {
			case "start":
				fmt.Println("Start")
				Startm(string(d.Body))
			case "register":
				fmt.Println("register")
				Registerm(string(d.Body))
			case "payment":
				fmt.Println("payment")
				Paymentm(string(d.Body))

			}
		}
	}()

	<-forever
}

//startfunction

func Startm(dat string) {
	fmt.Println("hahah This is " + dat)
}

//register manager

func Registerm(dat string) {
	fmt.Println("register" + dat)
}

//payment manager

func Paymentm(dat string) {
	fmt.Println(dat + "hjhhh")
}
