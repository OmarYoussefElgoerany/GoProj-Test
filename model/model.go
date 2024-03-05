package model

import "gorm.io/gorm"

//Models
type Order struct {
	ID         uint `gorm:"primary key;autoIncrement" json:"id"`
	Title     string `json:"title"`
}


func MigrateDB(db *gorm.DB) error {
	err := db.AutoMigrate(&Order{})
	return err
}
