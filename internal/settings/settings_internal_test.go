package settings

import (
	"os"
	"testing"

	"github.com/spf13/viper"
)

func Test_envPrefix(t *testing.T) {
	os.Setenv("FOO", "notsomewhere")
	os.Setenv("DEVID_FOO", "somewhere")

	Init()

	if viper.GetString("foo") != "somewhere" {
		t.Errorf("foo got %v, want %v", viper.GetString("foo"), os.Getenv("DEVID_FOO"))
	}
}

func Test_setXDGBaseDirs(t *testing.T) {
	setXDGBaseDirs()

	tests := []struct {
		name string
		want string
	}{
		{name: "XDG_DATA_HOME", want: os.ExpandEnv("$HOME/.local/share")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if os.Getenv(tt.name) != tt.want {
				t.Errorf("%s got %v, want %v", tt.name, os.Getenv(tt.name), tt.want)
			}
		})
	}
}

func Test_setDefaults(t *testing.T) {
	// make sure that test order does not affect this test
	viper.Reset()
	setDefaults()

	tests := []struct {
		name string
		want string
	}{
		{name: "personas_location", want: "$XDG_DATA_HOME/devid/personas"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if viper.GetString(tt.name) != tt.want {
				t.Errorf("%s got %v, want %v", tt.name, viper.GetString(tt.name), tt.want)
			}
		})
	}
}

func Test_setConstants(t *testing.T) {
	setConstants()

	tests := []struct {
		name string
		want string
	}{
		{name: "active_persona_env", want: "DEVID_ACTIVE_PERSONA"},
		{name: "active_persona_path_env", want: "DEVID_ACTIVE_PERSONA_PATH"},
		{name: "shell_loader_filename", want: "load.sh"},
		{name: "shell_runner_filename", want: "run.sh"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if viper.GetString(tt.name) != tt.want {
				t.Errorf("%s got %v, want %v", tt.name, viper.GetString(tt.name), tt.want)
			}
		})
	}
}

func Test_expandEnvs(t *testing.T) {
	// make sure that test order does not affect this test
	viper.Reset()
	expandEnvs()

	tests := []struct {
		name string
		want string
	}{
		{name: "personas_location", want: os.ExpandEnv(viper.GetString("personas_location"))},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if viper.GetString(tt.name) != tt.want {
				t.Errorf("%s got %v, want %v", tt.name, viper.GetString(tt.name), tt.want)
			}
		})
	}
}
