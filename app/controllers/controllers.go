package controllers

import (
    "fmt"
    //"time"
    "net/http"
   _"github.com/jinzhu/gorm/dialects/mysql"
    "github.com/labstack/echo"
   	"uktrav_echo/app"
   	"github.com/jinzhu/gorm"
	//cgorm "uktrav_echo/db/gorm"
    "uktrav_echo/app/models"
    
)

var DB *gorm.DB

func initDB(){
	DB, err := gorm.Open("mysql", "smartworks:smartworks@/uktrav?charset=utf8&parseTime=True&loc=127.0.0.1:3306")
	if err!=nil {
                fmt.Println("DB Connected")       
     }
    defer DB.Close()
    
}

/*
type User struct {
    gorm.Model
    id          int       `gorm:"AUTO_INCREMENT;primary_key;column:id"`
    fname       string    `gorm:"size:30;column:fname"`
    lname       string    `gorm:"size:30;column:lname"`
    email       string    `gorm:"size:30;column:email"`
    password    string    `gorm:"size:30;column:password_salt"`
    authcode    string    `gorm:"size:30;column:authcode"`
    phone       string    `gorm:"size:15;column:phone"`
    created_dt  time.Time `gorm:"column:created_dt"`
    status      int       `gorm:"column:status"`
}*/


func Init() {
        initDB()
        app.Server.GET("/users", func(c echo.Context) error {
             var user models.User	
             //obj = DB.Find(&user).GetErrors()
             if err := DB.Find(&user).GetErrors(); err != nil {
             } else {
                fmt.Println(err)
                fmt.Println(user)
             }  
	         return c.String(http.StatusOK, "OK")
        })
        
     
}
