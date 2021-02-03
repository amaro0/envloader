package envloader

import (
	"os"
	"testing"
)

type testConfig struct {
	Port    string `env:"PORT" envDefault:"3002" validate:"numeric"`
	GinMode string `env:"GIN_MODE" envDefault:"debug" validate:"oneof=debug release"`
}

func TestLoad(t *testing.T) {
	err, conf := Load(testConfig{})

	if err != nil {
		t.Error("Should parse config without error")
	}

	testConf := conf.(*testConfig)

	if testConf.Port != "3002" {
		t.Errorf("Parsed port is incorrect, got: %s, want: %s", testConf.Port, "3002")
	}

	if testConf.GinMode != "debug" {
		t.Errorf("Parsed GinMode is incorrect, got: %s, want: %s", testConf.GinMode, "debug")
	}

	_ = os.Setenv("PORT", "1")
	_ = os.Setenv("GIN_MODE", "release")

	err, conf = Load(testConfig{})

	if err != nil {
		t.Error("Should parse config without error")
	}

	testConf = conf.(*testConfig)

	if testConf.Port != "1" {
		t.Errorf("Parsed port is incorrect, got: %s, want: %s", testConf.Port, "1")
	}

	if testConf.GinMode != "release" {
		t.Errorf("Parsed GinMode is incorrect, got: %s, want: %s", testConf.GinMode, "release")
	}

	_ = os.Setenv("PORT", "wrong")

	err, conf = Load(testConfig{})

	if err == nil {
		t.Error("Should error")
	}
}
