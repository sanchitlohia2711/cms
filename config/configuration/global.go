package configuration

import "time"

//Global :global config
type Global struct {
	BackupSchedulerInterval time.Duration `json:"backup_scheduler_interval"`
	WorkerCount             int           `json:"worker_count"`
}

//Process :process
func (g *Global) Process() {

}
