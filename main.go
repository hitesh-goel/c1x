package main

import (
	"io"
	"log"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

func main() {
	//As per https://github.com/julienschmidt/go-http-routing-benchmark using HttpRouter for routing purpose.
	server := httprouter.New()
	server.GET("/v1/normal", NormalGetHandler)
	server.GET("/v1/goroutine", GoroutineGetHandler)
	log.Fatal(http.ListenAndServe(":8080", server))
}

//GoroutineGetHandler to handle goroutine jobs scheduled
func GoroutineGetHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	go ScheduleJob()
	io.WriteString(w, "OK\n")
}

//NormalGetHandler to handle jobs scheduled
func NormalGetHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	ScheduleJob()
	io.WriteString(w, "OK\n")
}

//ScheduleJob of 30ms
func ScheduleJob() {
	time.Sleep(30 * time.Millisecond)
}
