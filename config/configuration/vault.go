package configuration

//Vault : vault struct
type Vault struct {
	Hostname  string `json:"hostname"`
	CredPath  string `json:"cred_path"`
	SkipVault bool   `json:"skip_vault"`
}

//Process : process config
func (p *Vault) Process() {

}
