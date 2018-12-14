package configuration

//MySQL  mysql
type MySQL struct {
	Hostname     string `toml:"hostname"`
	Username     string `toml:"username"`
	Password     string `toml:"password"`
	TableName    string `toml:"table_name"`
	MaxOpenConns int    `toml:"max_open_conns"`
	Schema       string `toml:"schema"`
}

//Process mysql
func (p *MySQL) Process() {

}
