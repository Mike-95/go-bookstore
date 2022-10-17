package models

import (
	"github.com/Mike-95/go-bookstore/pkg/config"
	"github.com/jinzhu/gorm"
)

// Creating models
var db *gorm.DB

type Book struct {
	gorm.Model
	Name string `gorm: ""json:"name"`
	Author string `json: "author"`
	Publication string `json: "publication"`
}

// Creating init function to connect with db
func init(){
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}
