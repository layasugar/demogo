// 测试kafka生产和消费

package main

import (
	"github.com/layasugar/kafka-go"
	"log"
)

var cfg = &gkafka.Config{
	Brokers:  "xxx",
	Topic:    "test",
	Group:    "test",
	User:     "xxx",
	Pwd:      "xxx",
	Ca:       "",
	Version:  "0.10.2.0",
	Protocol: "sasl_ssl",
}

func MainKafka() {
	Producer := gkafka.InitProducer(cfg)
	err := Producer.SendMsg("test", "测试kafka消息的推送", 6)
	if err != nil {
		log.Print(err)
	}

	log.Print("推送成功")

	consumerData := make(chan *gkafka.ConsumerData)
	go gkafka.InitConsumer(cfg, consumerData, gkafka.SetClientId("demo-go"))

	for data := range consumerData {
		log.Printf("pool submit topic:%q partition:%d offset:%d", data.Topic, data.Partition, data.Offset)
	}
}
