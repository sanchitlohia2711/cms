package configuration

//Kafka : kafka configuration struct
type Kafka struct {
	Broker          string `json:"broker"`
	GroupID         string `json:"group_id"`
	Topic           string `json:"topic"`
	IsActive        bool   `json:"is_active"`
	MessageBulkRead int    `json:"message_bulk_read"`
	MaxConsumer     int    `json:"max_consumer"`
}

//Process : process
func (k *Kafka) Process() {

}
