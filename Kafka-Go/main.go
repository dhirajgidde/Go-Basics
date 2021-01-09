package main

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/segmentio/kafka-go"
)

const (
	topic         = "MyTopic"
	brokerAddress = "localhost:9092"
)

type Logger interface {
	printf(string, ...interface{})
}

func produce(ctx context.Context) {
	//initailise a counter
	i := 0
	//Brocker used to connect with the server
	//topic : name of the topic
	//l := log.New(os.Stdout, "Kafka reader: ", 0)
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers:      []string{brokerAddress},
		Topic:        topic,
		RequiredAcks: 1,
	})

	// WriteMessages are always in form of key-and-value pair
	for {
		err := w.WriteMessages(ctx, kafka.Message{
			Key:   []byte(strconv.Itoa(i)),
			Value: []byte(strconv.Itoa(i)),
		})

		if err != nil {
			panic("Could not write a message" + err.Error())
		}
		fmt.Println("Writes: ", i)
		i++
		time.Sleep(time.Second)
	}

}

func consume(ctx context.Context) {
	//l := log.New(os.Stdout, "Kafka Writer: ", 0)
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{brokerAddress},
		Topic:   topic,
		GroupID: "my-group",
	})

	for {
		msg, err := r.ReadMessage(ctx)
		if err != nil {
			panic("Could not read the messages" + err.Error())
		}
		fmt.Println(string(msg.Value))
		time.Sleep(time.Second)
	}
}

func main() {
	ctx := context.Background()

	go produce(ctx)
	consume(ctx)

}
