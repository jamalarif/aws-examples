package main

import (
    "log"
    "sync"

    "github.com/nsqio/go-nsq"
)

func main() {
    wg := &sync.WaitGroup{}
    wg.Add(1)

    decodeConfig := nsq.NewConfig()
    c, err := nsq.NewConsumer("NSQ_Topic_1", "NSQ_Channel_1", decodeConfig)
    if err != nil {
        log.Panic("Could not create consumer")
    }
    //c.MaxInFlight defaults to 1

    c.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
        log.Println("NSQ message received:")
        log.Println(string(message.Body))
        return nil
    }))

    err = c.ConnectToNSQLookupd("10.0.20.173:4161")
    if err != nil {
        log.Panic("Could not connect")
    }
    log.Println("Awaiting messages from NSQ topic \"My NSQ Topic\"...")
    wg.Wait()
}
