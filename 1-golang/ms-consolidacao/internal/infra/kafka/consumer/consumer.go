package consumer

import (
	"context"
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/sialka/cartola_fc/internal/infra/kafka/factory"
	"github.com/sialka/cartola_fc/pkg/uow"
)

func Consume(topics []string, servers string, msgChan chan *kafka.Message) {
	kafkaConsumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": servers,
		"group.id":          "gostats",
		"auto.offset.reset": "earliest", // pegar msg mais novas
	})
	if err != nil {
		panic(err)
	}
	// Se inscreve no topico para ler as msg
	kafkaConsumer.SubscribeTopics(topics, nil)

	// Loop para ler as mensagens
	for {
		msg, err := kafkaConsumer.ReadMessage(-1)
		if err == nil {
			msgChan <- msg // Pega a msg e add no canal msgChan
		}
	}
}

// Processo que le as msg do kafka
func ProcessEvents(ctx context.Context, msgChan chan *kafka.Message, uwo uow.UowInterface) {
	for msg := range msgChan {
		fmt.Println("Received message", string(msg.Value), "on topic", *msg.TopicPartition.Topic)
		strategy := factory.CreateProcessMessageStrategy(*msg.TopicPartition.Topic)
		err := strategy.Process(ctx, msg, uwo)
		if err != nil {
			fmt.Println(err)
		}
	}
}
