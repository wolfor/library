// KafkaHelper project KafkaHelper.go
package KafkaHelper

import (
	"log"
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/Shopify/sarama"
	"github.com/bsm/sarama-cluster" //support automatic consumer-group rebalancing and offset tracking
)

type KafkaHelper struct {
	address, user, password string
}

func NewKafkaHelper(kafkaAddr, kafkaUser, kafkaPWD string) *KafkaHelper {
	helper := new(KafkaHelper)
	helper.address = kafkaAddr
	helper.user = kafkaUser
	helper.password = kafkaPWD

	return helper
}

type ConsumerCallback func(message []byte) bool

//consumer 消费者
// This example shows how to use the consumer to read messages
// from a multiple topics through a multiplexed channel.
func (this *KafkaHelper) Consumer(groupID, topics string, callback ConsumerCallback) {
	config := cluster.NewConfig()
	config.Group.Return.Notifications = true
	config.Consumer.Return.Errors = true
	config.Consumer.Offsets.CommitInterval = 1 * time.Second
	config.Consumer.Offsets.Initial = sarama.OffsetNewest //初始从最新的offset开始
	//kafka 用户认证
	config.Net.SASL.Enable = true

	config.Net.SASL.User = this.user
	config.Net.SASL.Password = this.password

	consumer, err := cluster.NewConsumer(strings.Split(this.address, ","), groupID, strings.Split(topics, ","), config)

	if err != nil {
		log.Println("Failed open consumer: %v", err)
		return
	}
	defer consumer.Close()

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	// 接收错误
	go func() {
		for err := range consumer.Errors() {
			log.Printf("Error: %s\n", err.Error())
		}
	}()

	// 打印一些rebalance的信息
	go func() {
		for ntf := range consumer.Notifications() {
			log.Printf("Rebalanced: %+v\n", ntf)
		}
	}()

	// 消费消息
	for {
		select {
		case msg, ok := <-consumer.Messages():
			if ok {
				//				fmt.Fprintf(os.Stdout, "%s/%d/%d\t%s\n", msg.Topic, msg.Partition, msg.Offset, msg.Value)
				//				log.Println("%s/%d/%d\t%s\n", msg.Topic, msg.Partition, msg.Offset, msg.Value)

				if callback(msg.Value) {
					log.Println("kafka consumer finish")

					consumer.MarkOffset(msg, "") // 提交offset
				} else {
					log.Println("kafka consumer no finish")
				}

				return
			}
		case <-signals:
			return
		}
	}
}

//同步生产者
//并发量小时，可以用这种方式
func (this *KafkaHelper) SyncProducer(topics, pushData string) error {
	config := sarama.NewConfig()
	//  config.Producer.RequiredAcks = sarama.WaitForAll
	//  config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true
	config.Producer.Timeout = 5 * time.Second
	//kafka 用户认证
	config.Net.SASL.Enable = true
	config.Net.SASL.User = this.user
	config.Net.SASL.Password = this.password

	producer, err := sarama.NewSyncProducer(strings.Split(this.address, ","), config)
	defer producer.Close()
	if err != nil {
		//		log.Println(err)
		return err
	}

	msg := &sarama.ProducerMessage{
		Topic: topics,
		Value: sarama.ByteEncoder(pushData),
	}

	if _, _, err := producer.SendMessage(msg); err != nil {
		//		log.Println(err)
		return err
	}

	return nil
}

//异步生产者
//并发量大时，必须采用这种方式
func (this *KafkaHelper) AsyncProducer(topics, pushData string) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true //必须有这个选项
	config.Producer.Timeout = 5 * time.Second
	//kafka 用户认证
	config.Net.SASL.Enable = true
	config.Net.SASL.User = this.user
	config.Net.SASL.Password = this.password

	producer, err := sarama.NewAsyncProducer(strings.Split(this.address, ","), config)
	defer producer.Close()
	if err != nil {
		return
	}

	//必须有这个匿名函数内容
	go func(producer sarama.AsyncProducer) {
		errors := producer.Errors()
		success := producer.Successes()
		for {
			select {
			case err := <-errors:
				if err != nil {
					log.Println(err)
				}
			case <-success:
			}
		}
	}(producer)

	msg := &sarama.ProducerMessage{
		Topic: topics,
		Value: sarama.ByteEncoder(pushData),
	}

	producer.Input() <- msg
}

/*
func (this *KafkaHelper) client() {
	config := cluster.NewConfig()
	config.Group.Return.Notifications = true
	config.Consumer.Return.Errors = true
	config.Consumer.Offsets.CommitInterval = 1 * time.Second
	config.Consumer.Offsets.Initial = sarama.OffsetNewest //初始从最新的offset开始
	//kafka 用户认证
	config.Net.SASL.Enable = true

	config.Net.SASL.User = this.user
	config.Net.SASL.Password = this.password

	client, err := cluster.NewClient(strings.Split(this.address, ","), config)

	if err != nil {
		panic(err)
	}
	defer client.Close()

	client.WritablePartitions()
}
*/
