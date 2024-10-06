package main

import (
	"context"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type PlaceInfo struct {
	gorm.Model
	PlaceName string `gorm:"uniqueIndex"`
	Passwords []PasswordModel
}

type PasswordModel struct {
	gorm.Model
	PlaceInfoID uint
	Name        string
	Password    string
}

// App struct
type App struct {
	ctx context.Context
	db  *gorm.DB
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	a.db = db
	a.db.AutoMigrate(&PlaceInfo{}, &PasswordModel{})
}

func (a *App) DBAccess() []PlaceInfo {

	p1 := PasswordModel{Name: "取引", Password: "XXXXX"}
	p2 := PasswordModel{Name: "取引2", Password: "YYYYY"}
	a.db.Create(&PlaceInfo{PlaceName: "テスト5", Passwords: []PasswordModel{p1, p2}})

	var placeInfos []PlaceInfo
	err := a.db.Model(&PlaceInfo{}).Preload("Passwords").Find(&placeInfos).Error
	if err != nil {
		log.Println(err.Error())
		return nil
	}

	return placeInfos
}

func (a *App) AddPassword(placeID uint, name string, password string) string {
	result := a.db.Create(&PasswordModel{PlaceInfoID: placeID, Name: name, Password: password})
	if result.Error != nil {
		log.Println(result.Error.Error())
		return result.Error.Error()
	}
	return "success"
}
