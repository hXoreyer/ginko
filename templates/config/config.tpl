package config

type Configuration struct {
	App      App      `mapstructure:"app" json:"app" yaml:"app"`
	Database Database `mapstructure:"database" json:"database" yaml:"database"`
	Log      Log      `mapstructure:"log" json:"log" yaml:"log"`
}
