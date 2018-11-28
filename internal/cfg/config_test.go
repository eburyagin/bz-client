package cfg

import (
	"testing"
)

func TestLoadConfig(t *testing.T) {
	cf := "../../config.json"
	_, err := Load(cf)
	if err != nil {
		t.Errorf("cfg.Load(%s), error: %s", cf, err.Error())
	}
}
