package configuration

//Cron : back up cron
type Cron struct {
	BackupCronEnabled bool `json:"backup_cron_enabled"`
}

//Process : process the cron
func (l *Cron) Process() {

}
