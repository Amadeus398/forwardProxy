package config

var (
	Cache = EnvCache{}
)

type EnvCache struct {
	Port string `envconfig:"PORT"`
}
