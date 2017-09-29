package controller

import (
	"ADT/RabbitMq"

	"fmt"
	"io/ioutil"
	"net/http"
)

func InitRecivem() {
	kl := []string{"start", "register", "payment"}

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
				Startm(d.Body)
			case "register":
				fmt.Println("register")
				Registerm(d.Body)
			case "payment":
				fmt.Println("payment")
				Paymentm(d.Body)
			}
		}
	}()

	<-forever
}

//startfunction

func Startm(dat []byte) {

	//json.Unmarshal(dat)
}

//register manager

func Registerm(dat []byte) {

}

//payment manager

func Paymentm(dat []byte) {

}

//send request data

func SendGetReq(urls string) []byte {

	res, err := http.Get(urls)
	if err != nil {
		return []byte("")
	}

	body, era := ioutil.ReadAll(res.Body)

	if era != nil {
		return []byte("")
	}
	//no content back the platform
	return body
}
