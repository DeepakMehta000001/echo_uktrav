package models

import (
       "time"
)

type User struct {
    Id           int       `gorm:"AUTO_INCREMENT;primary_key;column:id"`
    Fname        string    `gorm:"size:30;column:fname"`
    Lname        string    `gorm:"size:30;column:lname"`
    Email        string    `gorm:"size:30;column:email"`
    Password     string    `gorm:"size:30;column:password_salt"`
    Authcode     string    `gorm:"size:30;column:authcode"`
    Phone        string    `gorm:"size:15;column:phone"`
    CreatedDt    time.Time `gorm:"column:created_dt"`
    Status       int       `gorm:"column:status"`
}


