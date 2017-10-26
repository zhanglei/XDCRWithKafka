package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"sync"

	"../aof"

	"github.com/Shopify/sarama"
)

var (
	broker = flag.String("broker", "localhost:9092", "broker address")
	topic  = flag.String("topic", "", "Required, the topic to consume from.  You can create a topic from console.")
	offset = flag.Int("offset", 0, "data offset to consume from")
)

func main() {
	if len(os.Args[1:]) == 0 {
		flag.Usage()
		os.Exit(1)
	}

	flag.Parse()

	conn, err := net.Dial("tcp", "localhost:6380")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer conn.Close()
	go func() {
		for {
			out := make([]byte, 1024)
			_, err := conn.Read(out)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			fmt.Println("resp: ", string(out))
		}
	}()

	if *topic == "" {
		panic("Argument topic is required.")
	}

	sarama.Logger = log.New(os.Stderr, "[sarama]", log.LstdFlags)

	config := sarama.NewConfig()
	config.Version = sarama.V0_10_1_0

	consumer, err := sarama.NewConsumer([]string{*broker}, config)
	if err != nil {
		log.Panic(err)
	}

	defer func() {
		if err := consumer.Close(); err != nil {
			log.Panic(err)
		}
	}()

	partitions, err := consumer.Partitions(*topic)
	if err != nil {
		log.Panic(err)
	}

	var wg sync.WaitGroup
	wg.Add(len(partitions))

	for _, partition := range partitions {
		log.Println("consume partition:", partition)

		shutdown := make(chan os.Signal, 1)
		signal.Notify(shutdown, os.Interrupt)

		go func(partition int32) {
			partitionConsumer, err := consumer.ConsumePartition(*topic, partition, int64(*offset))
			if err != nil {
				log.Fatalln(err)
				return
			}

			defer func() {
				wg.Done()
				if err := partitionConsumer.Close(); err != nil {
					log.Fatalln(err)
				}
			}()

		ConsumerLoop:
			for {
				select {
				case err := <-partitionConsumer.Errors():
					if err != nil {
						log.Fatalln("error:", err)
					}
				case msg := <-partitionConsumer.Messages():
					if msg != nil {
						op := aof.Operation{}
						op.UnmarshalMsg(msg.Value)
						rw := aof.RecordWriter{}
						op.ToAof(&rw)
						conn.Write(rw)
					}
				case <-shutdown:
					log.Println("stop consuming partition:", partition)
					break ConsumerLoop
				}
			}
		}(partition)
	}

	wg.Wait()
}

func checkFile(path string) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	stat, err := file.Stat()
	if err != nil {
		panic(err)
	}

	if stat.Size() == 0 {
		panic("Please replace " + path + " with your own. ")
	}

}
