package controllers

import (
    "fmt"
    "net/http"
    "strconv"
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
      
      //GET record by id param
      app.Server.GET("/users/:id",get_user(DB))
      
      //POST Create a new user
      app.Server.POST("/users",create_user(DB))
      
      //PUT Update a user
      app.Server.PUT("/users/:id",update_user(DB))
      
}


func get_users(db *gorm.DB) echo.HandlerFunc {

    return func(c echo.Context) error {
            users :=  []models.User{}
            if db.Find(&users).RecordNotFound(){
                return c.JSON(http.StatusOK,"No Record Found") 
            }
	        return c.JSON(http.StatusOK,users)
      }

}

func get_user(db *gorm.DB) echo.HandlerFunc {

    return func(c echo.Context) error {
            id, _ := strconv.Atoi(c.Param("id"))
            var user models.User
            if db.First(&user,id).RecordNotFound(){
                return c.JSON(http.StatusOK,"No Record Found")    
            }
	        fmt.Println(user.Authcode)
	        return c.JSON(http.StatusOK,user)
      }

}


func create_user(db *gorm.DB) echo.HandlerFunc {

    return func(c echo.Context) error {
                user := models.User{}
	            if err := c.Bind(&user); err != nil {
		            return err
	            }
	            db.Create(&user)
	            if !db.NewRecord(user){
	                fmt.Println("new record created")
	            }
	       return c.JSON(http.StatusCreated, user)
      }

}


func update_user(db *gorm.DB) echo.HandlerFunc {
    return func(c echo.Context) error {
	            user_temp := echo.Map{}
	            //user := models.User{}
	            if err := c.Bind(&user_temp); err != nil {
		            return err
	            }
	            id, _ := strconv.Atoi(c.Param("id"))
	            var user models.User
	            db.First(&user,id)
	            user.Lname = user_temp["lname"].(string)
	            db.Save(&user)
	        return c.JSON(http.StatusOK, user)
    }
}
