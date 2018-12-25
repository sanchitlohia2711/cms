package configuration

//ES  mysql
type ES struct {
	Hostname     string `toml:"hostname"`
	ProductIndex string `toml:"product_index"`
	SkuIndex     string `toml:"sku_index"`
	EntityIndex  string `toml:"entity_index"`
}

//Process mysql
func (p *ES) Process() {

}
