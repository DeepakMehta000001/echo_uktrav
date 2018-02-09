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

/*
var DB *gorm.DB

func initDB(){
	DB, err := gorm.Open("mysql", "smartworks:smartworks@/uktrav?charset=utf8")
	if err!=nil {
            fmt.Println("DB not Connected")       
     }
    defer DB.Close()
    
}



type User struct {
    Id           int       `gorm:"AUTO_INCREMENT;primary_key;column:id"`
    Fname        string    `gorm:"size:30;column:fname"`
    Lname        string    `gorm:"size:30;column:lname"`
    Email        string    `gorm:"size:30;column:email"`
    PasswordSalt string    `gorm:"size:30;column:password_salt"`
    Authcode     string    `gorm:"size:30;column:authcode"`
    Phone        string    `gorm:"size:15;column:phone"`
    CreatedDt    time.Time `gorm:"column:created_dt"`
    Status       int       `gorm:"column:status"`
}

*/


type H map[string]interface{}

func Init() {
      var DB *gorm.DB
      DB, err := gorm.Open("mysql", "smartworks:smartworks@/uktrav?charset=utf8&parseTime=True")
	  if err!=nil {
            fmt.Println("DB not Connected")       
          }
      //defer DB.Close()
      
      app.Server.GET("/users",get_last_user(DB))
}


func get_last_user(db *gorm.DB) echo.HandlerFunc {

    return func(c echo.Context) error {
            var user models.User
            //obj := DB.First(&user,1)
            //fmt.Println(DB.HasTable(&User{}))
            db.Last(&user)
            fmt.Println(user.Fname)
            //defer DB.Close()
	        return c.JSON(http.StatusCreated, H{"id" : user.Id})
      }

}




