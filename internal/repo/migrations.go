package repo

import (
	"log"
)

func (d *Database) MigrateDB() error {
	log.Println("Database migration started.....")

	err := d.Client.AutoMigrate(&User{})
	if err != nil {
		return err
	}
	return nil
}
