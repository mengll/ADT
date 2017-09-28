package controller

import (
	"ADT/RabbitMq"
)

var Ch *RabbitMq.RabbitMq

func init() {
	//1 create all
	Ch = RabbitMq.NewRabbitMq("amqp://guest:guest@127.0.0.1:5672/")
	//2 connect
	Ch.Connect()

	//close
	//defer Ch.Closeq()
}

//send the data

func SendBrocast(router, content string) {

	//3 crewate exchangetype
	exdata := Ch.NewExchangeDecType()
	exdata.Name = "adt"
	exdata.Type = "direct"
	exdata.Durable = true

	//4
	Ch.NewExchangeDec(exdata)

	//5
	pusdat := Ch.NewPublicDatType()
	pusdat.Body = content
	pusdat.Routerkey = router
	pusdat.Exchange = exdata.Name

	Ch.PublishTo(pusdat)
}
