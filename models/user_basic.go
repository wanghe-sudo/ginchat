package models

import "gorm.io/gorm"

type UserBasic struct {
	gorm.Model
	Name         string
	Password     string
	Phone        string
	Email        string
	Identify     string
	ClientIp     string
	ClientPort   string
	LoginTime    uint64
	HeartBeat    uint64
	LoginOutTime uint64
	IsLogout     bool
	DeviceInfo   string
}

func (table *UserBasic) TableName() string {
	return "user_basic"
}
