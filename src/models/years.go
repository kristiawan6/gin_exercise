package models

import (
	"gingonic/src/config"

	"github.com/jinzhu/gorm"
)

type Year struct {
	gorm.Model
	Name string
}

func SelectAllYear() []Year {
	items := []Year{}
	config.DB.Raw("SELECT * FROM years").Scan(&items)
	return items
}

func SelectYearById(ID int) Year {
	item := Year{}
	config.DB.Raw("SELECT * FROM years WHERE id = ?", ID).Scan(&item)
	return item
}

func CreateYear(Name string) []Year {
	items := []Year{}
	config.DB.Raw("INSERT INTO years (created_at, updated_at, deleted_at, name) Values(NULL, NULL, NULL, ?)", Name).Scan(&items)
	return items
}

func UpdateYear(ID int, Name string) []Year {
	items := []Year{}
	config.DB.Raw("UPDATE years SET name = ? WHERE id = ?", Name, ID).Scan(&items)
	return items
}

func DeleteYear(ID int) []Year {
	items := []Year{}
	config.DB.Raw("DELETE FROm years WHERE id  = ?", ID).Scan(&items)
	return items
}