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
      //Routes
      
      //GET all records
      app.Server.GET("/users",get_users(DB))
      
      //POST
      
      
}


func get_users(db *gorm.DB) echo.HandlerFunc {

    return func(c echo.Context) error {
            users :=  []models.User{}
            db.Find(&users)
	        return c.JSON(http.StatusOK,users)
      }

}




