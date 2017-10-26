package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"../aof"
	"github.com/Shopify/sarama"
)

var (
	broker   = flag.String("broker", "localhost:9092", "broker address")
	topic    = flag.String("topic", "", "Required, the topic to consume from.  You can create a topic from console.")
	feedfile = flag.String("feed", "", "Required, the operation feed from")
)

func main() {
	if len(os.Args[1:]) == 0 {
		flag.Usage()
		os.Exit(1)
	}

	flag.Parse()

	if *topic == "" {
		panic("Argument topic is required.")
	}
	if *feedfile == "" {
		panic("Argument feedfile is required.")
	}

	sarama.Logger = log.New(os.Stderr, "[sarama]", log.LstdFlags)

	config := sarama.NewConfig()
	config.Version = sarama.V0_10_1_0

	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = sarama.WaitForLocal
	config.Producer.Retry.Max = 5

	producer, err := sarama.NewSyncProducer([]string{*broker}, config)
	if err != nil {
		log.Panic(err)
	}

	defer func() {
		if err := producer.Close(); err != nil {
			log.Panic(err)
		}
	}()

	file, err := os.Open("/home/dev/RunEnv/Redis3.0/redis/6379/appendonly.aof")
	defer file.Close()

	if err != nil {
		fmt.Errorf("Open file failed. Error:%s", err.Error())
		return
	}
	reader := aof.NewBufioTailReader(file, 10)

	for {
		op, err := reader.ReadOperation()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Errorf("ReadOperation failed. %s", err.Error())
			break
		}
		//create a message
		data, _ := op.MarshalMsg(nil)
		msg := &sarama.ProducerMessage{
			Topic: *topic,
			Key:   sarama.StringEncoder(op.Key),
			Value: sarama.StringEncoder(data),
		}
		partition, offset, err := producer.SendMessage(msg)
		if err != nil {
			log.Fatal(err)
		}
		log.Println("message sent, topic:", *topic, "partition:", partition, "offset:", offset)
	}

}
