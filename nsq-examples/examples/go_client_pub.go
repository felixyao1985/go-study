package main

import (
	"fmt"
	"github.com/youzan/go-nsq"
)

var producer *nsq.Producer

func main() {
	nsqd := "172.0.0.1:4150"
	producer, err := nsq.NewProducer(nsqd, nsq.NewConfig())

	var data = []byte(`{"Name":"felix","Age":33}`)

	producer.Publish("test2", []byte(string(data)))
	if err != nil {
		fmt.Println("error")
		panic(err)
	}
}
