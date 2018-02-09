package controllers

import (
    "fmt"
    "net/http"
   _"github.com/jinzhu/gorm/dialects/mysql"
    "github.com/labstack/echo"
   	"uktrav_echo/app"
   	"github.com/jinzhu/gorm"
    "uktrav_echo/app/models"
    
)

type H map[string]interface{}

func Init() {
      var DB *gorm.DB
      DB, err := gorm.Open("mysql", "smartworks:smartworks@/uktrav?charset=utf8&parseTime=True")
	  if err!=nil {
            fmt.Println("DB not Connected")       
          }   
      app.Server.GET("/users",get_last_user(DB))
}


func get_last_user(db *gorm.DB) echo.HandlerFunc {

    return func(c echo.Context) error {
            var user models.User
            db.Last(&user)
            fmt.Println(user.Fname)
	        return c.JSON(http.StatusCreated, 
	        H{"id" : user.Id,
	        })
      }

}




