package main

import (
  "log"

  "github.com/nsqio/go-nsq"
)

func main() {
    config := nsq.NewConfig()
    p, err := nsq.NewProducer("127.0.0.1:4150", config)
    if err != nil {
        log.Panic(err)
    }
    err = p.Publish("NSQ_Topic_1", []byte("Sample NSQ message 1 "))
    err = p.Publish("NSQ_Topic_2", []byte("Sample NSQ message 1 "))
    err = p.Publish("NSQ_Topic_3", []byte("Sample NSQ message 1 "))
    if err != nil {
        log.Panic(err)
    }
}
