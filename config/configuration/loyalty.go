package configuration

import (
	"time"
)

const (
	EnvLoyaltyAppID = "cashbackworker_LOYALTY_APP_ID"
	EnvLoyaltyKey   = "cashbackworker_LOYALTY_KEY"
)

type Loyalty struct {
	Enable  bool          `json:"enable"`
	URL     string        `json:"url"`
	Timeout time.Duration `json:"timeout"`
	AppID   string        `json:"app_id"`
	Key     string        `json:"key"`
	Method  string        `json:"method"`
}

func (l *Loyalty) Process() {

}
