package config

import "github.com/IBM/sarama"

type KafkaConfig struct {
	Brokers   []string `yaml:"brokers"`
	Group     string   `yaml:"group"`
	TopicName string   `yaml:"topicName"`
}

func (k KafkaConfig) Config() *sarama.Config {
	config := sarama.NewConfig()
	config.Version = sarama.V2_6_0_0
	config.Consumer.Group.Rebalance.GroupStrategies = []sarama.BalanceStrategy{sarama.NewBalanceStrategyRoundRobin()}
	config.Consumer.Offsets.Initial = sarama.OffsetOldest

	return config
}
