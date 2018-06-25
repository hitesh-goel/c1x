package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/julienschmidt/httprouter"
)

var (
	maxWorkers, _       = GetIntEnv("MAX_WORKERS")
	bufferedMessages, _ = GetIntEnv("MAX_BUFFERED_MSG")
	ch                  = make(chan string, bufferedMessages)
)

//GetIntEnv returns Integer value of enverionment variables
func GetIntEnv(key string) (int, error) {
	s := os.Getenv(key)
	v, err := strconv.Atoi(s)
	if err != nil {
		os.Exit(1)
	}
	return v, nil
}

func main() {
	//As per https://github.com/julienschmidt/go-http-routing-benchmark using HttpRouter for routing purpose.
	server := httprouter.New()
	server.GET("/v1/normal", NormalGetHandler)
	server.GET("/v1/goroutine", GoroutineGetHandler)
	server.GET("/v1/gochannels", GoChannelGetHandler)

	//Start a channel worker
	Worker(ch)

	// Start http server Listen to port 8080
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

//GoChannelGetHandler Channel handler
func GoChannelGetHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	WriteToChannel(ch)
	io.WriteString(w, "OK\n")
}

//WriteToChannel which will act as a queue between worker & server
func WriteToChannel(ch chan<- string) {
	ch <- fmt.Sprintf("do 30 sec Job")
}

//Worker to consume messages
func Worker(ch <-chan string) {
	for k := 0; k < maxWorkers; k++ {
		go func(c <-chan string, k int) {
			for {
				select {
				case msg1 := <-c:
					fmt.Println(msg1+"in worker ", k)
					ScheduleJob()
				}
			}
		}(ch, k)
	}
}
