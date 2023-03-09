package main

import (
	consumer "fintech4go/mds_consumer"
	"fintech4go/service"
	"fmt"
)

func main() {
	fmt.Println("welcome to fintech4go.")
	producer := service.NewKafkaProducerWrapper("grpc_quote_realtime3", 0, []string{"10.80.88.25:9092"})
	producer.Work("fff-chao.zhou")
	producer.Work("fff-chao.zhou22")
	producer.Work("fff-chao.zhou333")

	mds_consumer := &consumer.MarketDataConsumer{}
	consumer := service.NewKafkaConsumerWrapper("grpc_quote_realtime2", 0, []string{"10.80.88.25:9092"})
	consumer.Consumer = mds_consumer
	consumer.Work()

}
