package models

import (
       "time"
)

/*
type User struct {
    Id           int       `gorm:"AUTO_INCREMENT;primary_key;column:id"` `json:"id"`
    Fname        string    `gorm:"size:30;column:fname"`                 `json:"fname"`
    Lname        string    `gorm:"size:30;column:lname"`                 `json:"lname"`
    Email        string    `gorm:"size:30;column:email"`                 `json:"email"`
    Password     string    `gorm:"size:30;column:password_salt"`         `json:"password"`
    Authcode     string    `gorm:"size:30;column:authcode"`              `json:"authcode"`
    Phone        string    `gorm:"size:15;column:phone"`                 `json:"phone"`
    CreatedDt    time.Time `gorm:"column:created_dt"`                    `json:"created_dt"`
    Status       int       `gorm:"column:status"`                        `json:"status"`
}
*/


type User struct {
    Id           int       `json:"id"`
    Fname        string    `json:"fname"`
    Lname        string    `json:"lname"`
    Email        string    `json:"email"`
    PasswordSalt string    `json:"password"`
    Authcode     string    `json:"authcode"`
    Phone        string    `json:"phone"`
    CreatedDt    time.Time `json:"created_dt"`
    Status       int       `json:"status"`
}


