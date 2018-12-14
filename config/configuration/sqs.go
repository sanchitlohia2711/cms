package configuration

//SQS : sqs
type SQS struct {
	QueueURL               string `json:"queue_url"`
	Region                 string `json:"region"`
	MaxSQSReceiveConsumer  int    `json:"max_sqs_receive_consumer"`
	MaxSQSWorkers          int    `json:"max_sqs_workers"`
	MaxMessagesRead        int64  `json:"max_messages_read"`
	MessageReadTimeoutInMS int64  `json:"message_read_timeout_in_ms"`
	MaxMessageTransit      int64  `json:"max_message_transit"`
	MaxReceiveCount        int    `json:"max_receive_count"`
	Enabled                bool   `json:"enabled"`
}

//Process : process
func (k *SQS) Process() {

}
