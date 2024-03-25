package main

import (
	"fmt"
	"github.com/Shopify/sarama"
)

func main() {
	//sarama第三方库，kafka的client
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForLocal        //发完数据仅需要leader确认
	config.Producer.Partitioner = sarama.NewRandomPartitioner //新选出一个partition
	config.Producer.Return.Successes = true                   //成功交付的message在success chanel返回

	//构造消息
	msg := &sarama.ProducerMessage{}
	msg.Topic = "web_log"
	//连接kafka
	client, err := sarama.NewSyncProducer([]string{"127.0.0.1:9092"}, config)
	if err != nil {
		fmt.Println("producer closed,err:", err)
		return
	}
	defer client.Close()
	//发送消息
	for i := 0; i <= 1000; i++ {
		msg.Value = sarama.StringEncoder(fmt.Sprintf("datastruct%d", i))
		pid, offset, err := client.SendMessage(msg)
		if err != nil {
			fmt.Println("send msg failed,err:", err)
			return
		}
		fmt.Printf("pid:%v offset:%v\n", pid, offset)
	}
}
