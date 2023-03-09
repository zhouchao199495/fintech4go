package service

type KafkaConfiguration struct {
	Topic            string
	Partition        int
	BootstrapServers []string
}
