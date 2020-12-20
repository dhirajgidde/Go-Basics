package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/segmentio/kafka-go"
)

const (
	topic         = "message-log"
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
	l := log.New(os.Stdout, "Kafka reader: ", 0)
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{brokerAddress},
		Topic:   topic,

		Logger: l,
	})

	// WriteMessages are always in form of key-and-value pair
	for j := 0; j < 7; j++ {
		err := w.WriteMessages(ctx, kafka.Message{
			Key:   []byte(strconv.Itoa(i)),
			Value: []byte("The Message from a producer" + strconv.Itoa(i)),
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
	l := log.New(os.Stdout, "Kafka Writer: ", 0)
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{brokerAddress},
		Topic:   topic,
		GroupID: "my-group",
		Logger:  l,
	})

	for j := 0; j < 7; j++ {
		msg, err := r.ReadMessage(ctx)
		if err != nil {
			panic("Could not read the messages" + err.Error())
		}
		fmt.Println("Recieved in consumer :", string(msg.Value))
	}
}

func main() {
	ctx := context.Background()

	go produce(ctx)
	consume(ctx)
}
