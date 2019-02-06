package model

import "github.com/jinzhu/gorm"

//Success response
// swagger:response UserWallet
type UserWallet struct {
	Fid int `gorm:"not null"`
	Wallet string `gorm:"not null;unique"`
	Coin string `gorm:"not null"`
}


type OAuthUser struct {
	gorm.Model
	Username string `gorm:"not null"`
	Email string `gorm:"not null;unique"`
	Password string `gorm:"not null"`
	Token string
	Active bool `gorm:"not null"`
}