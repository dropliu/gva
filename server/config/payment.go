package config

type Payment struct {
	ID          string `mapstructure:"id" json:"id" yaml:"id"`
	Env         string `mapstructure:"env" json:"env" yaml:"env"`
	OrgID       string `mapstructure:"org_id" json:"org_id" yaml:"org_id"`
	WorkspaceID string `mapstructure:"workspace_id" json:"workspace_id" yaml:"workspace_id"`
	PrivateKey  string `mapstructure:"private_key" json:"private_key" yaml:"private_key"`
}
