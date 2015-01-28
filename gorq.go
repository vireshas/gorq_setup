package main

import (
	"fmt"
	"github.com/goibibo/gorq"
	"github.com/goibibo/gorq-rcvr"
	"github.com/goibibo/t-settings"
)

func main() {

	settings.Configure()
	gorq.InitRedisPool("r1")
	rqrcvr.InitClient("r1")

	out := make(chan string)
	kwargs := map[string]string{"pubsub": "true"}

	job := gorq.NewRQJob("add.add", []string{"14", "14"}, kwargs)
	job.Enqueue()
	job.Start()

	rqrcvr.Subscribe(job.Id, out)
	go rqrcvr.StartPublisher()
	fmt.Println(<-out)
}
