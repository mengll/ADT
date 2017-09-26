// ADT project main.go
package main

import (
	"fmt"
	"net/http"
)

type AdtData struct {
	gameid   string
	imei     string
	adttype  string
	rid      string
	ordernum string
	ucid     string
	ip       string
	uuid     string
	channel  int
	actype   int
	addtime  string
}

//The data save

func dt(a string, b string) {
	fmt.Println(a, b)
}

func SaveRabbit(w http.ResponseWriter, r *http.Request) {
	//adt := &AdtData{}
	qu := r.URL.Query()
	gameid := qu.Get("gameid")

	//return back the data

	imei := qu.Get("imei")

	fmt.Println(interface{}(imei).(string)) // assert the data

	fmt.Println(len(imei))

	fmt.Println(imei)
	//chuangjian yige weifuwu de kuangjia de shixian
	w.Write([]byte(gameid))
}

func main1() {
	dt("2", "2")
	fmt.Println("Hello World!")
	http.HandleFunc("/adt", SaveRabbit)
	http.ListenAndServe(":8080", nil)
}
