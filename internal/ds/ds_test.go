package ds

import (
	"testing"

	cfg "bz-client/internal/cfg"
)

func TestConnection(t *testing.T) {
	cf := "../../config_test.json"
	config, err := cfg.Load(cf)
	if err != nil {
		t.Errorf("Config(%q) = %v, error: %s", cf, config, err)
	}
	db, err := NewConnect(config)
	if err != nil {
		t.Errorf("Connect(%q) = %v, error: %s", config, db, err)
	}
	defer Disconnect(db)
}

func TestDisconnection(t *testing.T) {
	cf := "../../config_test.json"
	config, err := cfg.Load(cf)
	if err != nil {
		t.Errorf("Config(%q) = %v, error: %s", cf, config, err)
	}
	db, _ := NewConnect(config)
	if err != nil {
		t.Errorf("Connect(%q) = %v, error: %s", config, db, err)
	}
	err = Disconnect(db)
	if err != nil {
		t.Errorf("Disconnect(%q), error: %s", db, err)
	}
}
