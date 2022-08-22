package main

import (
	"fmt"
	"github.com/Shopify/sarama"
	"sync"
)

var wg sync.WaitGroup

func main() {
	consumer, err := sarama.NewConsumer([]string{"127.0.0.1:9092"}, nil)
	if err != nil {
		fmt.Printf("fail to start consumer,err:%v\n", err)
		return
	}
	partitionList, err := consumer.Partitions("web_log") //获取web_log该分区下的所有partition
	if err != nil {
		fmt.Printf("fail to get list of partition:err%v\n", err)
		return
	}
	for _, partition := range partitionList {
		//针对每一个分区创建一个对应的消费者
		pc, err := consumer.ConsumePartition("web_log", int32(partition), sarama.OffsetNewest)
		if err != nil {
			fmt.Printf("ailed to start consumer for partition %d,err:%v\n", partition, err)
			return
		}
		defer pc.AsyncClose()
		wg.Add(1)
		go func(partitionConsumer sarama.PartitionConsumer) {
			for msg := range pc.Messages() {
				fmt.Printf("Partition:%d Offset:%d Key:%v Value:%s\n", msg.Partition, msg.Offset, msg.Key, string(msg.Value))
			}
			wg.Done()
		}(pc)
	}
	wg.Wait()
}
