package main

import (
	"encoding/json"
	"fmt"
	"github.com/youzan/go-nsq"
	"sync"
)

type NSQHandler struct {
}

type singer struct {
	Name string
	Age  int
}

func (this *NSQHandler) HandleMessage(msg *nsq.Message) error {

	var s singer

	json.Unmarshal([]byte(string(msg.Body)), &s)

	fmt.Println("receive", msg.NSQDAddress, msg.Partition, "message:", string(msg.Body), s.Name)

	return nil
}

/*

	WaitGroup总共有三个方法：Add(delta int),Done(),Wait()。

	Add:添加或者减少等待goroutine的数量

	Done:相当于Add(-1)

	Wait:执行阻塞，直到所有的WaitGroup数量变成0
*/
func testNSQ() {
	waiter := sync.WaitGroup{}
	waiter.Add(1)

	go func() {

		//defer waiter.Done()
		config := nsq.NewConfig()
		config.MaxInFlight = 9

		//建立多个连接
		for i := 0; i < 10; i++ {
			consumer, err := nsq.NewConsumer("test2", "nsq_felix", config)
			if nil != err {
				fmt.Println("err", err)
				return
			}

			consumer.AddHandler(&NSQHandler{})
			err = consumer.ConnectToNSQD("172.0.0.1:4150", i)
			if nil != err {
				fmt.Println("err", err)
				return
			}
		}
		select {}

	}()

	waiter.Wait()
}
func main() {
	testNSQ()
}
