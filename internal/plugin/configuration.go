package plugin

const (
	apiVersion = "v1"
)

func NewConfig() Config {
	return Config{
		APIVersion: apiVersion,
	}
}

type Config struct {
	APIVersion string `yaml:"apiVersion"`

	Identity struct {
		Email string
		Name  string
	}
	Envs map[string]string
}
