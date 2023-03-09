package consumer

import (
	fintech4go_pb_basic "fintech4go/pb/public/go/fintech4go.pb.basic"
	fintech4go_pb_rpc "fintech4go/pb/public/go/fintech4go.pb.rpc"
	"fmt"

	"google.golang.org/protobuf/proto"
)

type MarketDataConsumer struct {
}

func (consumer *MarketDataConsumer) Consume(message string) {
	quoteDataArray := &fintech4go_pb_rpc.QuoteDataPackageArray{}
	err := proto.Unmarshal([]byte(message), quoteDataArray)
	if err != nil {
		fmt.Printf("proto.Unmarshal failed, err:%v", err)
		return
	}
	consumer.ShowQuoteData(quoteDataArray)
}

func (consumer *MarketDataConsumer) ShowQuoteData(quote *fintech4go_pb_rpc.QuoteDataPackageArray) {
	for _, record := range quote.Record {
		consumer.ShowQuoteRecord(record)
	}
}

func (consumer *MarketDataConsumer) ShowQuoteRecord(record *fintech4go_pb_rpc.QuoteDataPackage) {
	switch record.MessageType {
	case 2:
		basicQuote := &fintech4go_pb_basic.BasicQot{}
		err := proto.Unmarshal(record.Payload, basicQuote)
		if err != nil {
			fmt.Printf("proto.Unmarshal failed, err:%v", err)
			return
		}
		fmt.Printf("basicQuote:%s", basicQuote.String())

	case 3:
		orderBook := &fintech4go_pb_basic.OrderBook{}
		err := proto.Unmarshal(record.Payload, orderBook)
		if err != nil {
			fmt.Printf("proto.Unmarshal failed, err:%v", err)
			return
		}
		fmt.Printf("orderBook:%s", orderBook.String())

	default:
		fmt.Printf("unknown message type:%d", record.MessageType)
	}
}
