package ds

import (
	"bz-lib/cfg"

	"github.com/go-pg/pg"
)

func NewConnect(config *cfg.Config) (*pg.DB, error) {
	db := pg.Connect(&pg.Options{
		Addr:     config.Datastore.Addr,
		User:     config.Datastore.User,
		Password: config.Datastore.Password,
		Database: config.Datastore.Database,
	})
	var n int
	_, err := db.QueryOne(pg.Scan(&n), "SELECT 1")
	if err != nil || n != 1 {
		return db, err
	}
	return db, nil
}

func Disconnect(db *pg.DB) error {
	db.Close()
	return nil
}
