package plugin

const (
	apiVersion = "v1"
)

func NewConfig() Config {
	return Config{
		APIVersion: apiVersion,
	}
}

// Config contains configuration available to configure this application
// behaviour.
// Using a separate configuration object to load external configuration allow
// stricter control of what is unmarshalled from the configuration YAML file.
// This should reduce YAML attach surface (DOS kind of attacks against this
// CLI are not part of the threat model).
type Config struct {
	APIVersion string `yaml:"apiVersion"`

	Identity struct {
		Email string
		Name  string
	}
	Envs map[string]string

	Ssh struct {
		Enabled bool

		Keys      []string
		CachePath string
	}

	Tmux struct {
		Enabled bool
	}
}
