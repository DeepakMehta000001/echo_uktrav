package controllers

import (
    "fmt"
    "io"
    "net/http"
    "strconv"
   _"github.com/jinzhu/gorm/dialects/mysql"
    "github.com/labstack/echo"
   	"uktrav_echo/app"
   	"github.com/jinzhu/gorm"
    "uktrav_echo/app/models"
    //"reflect"
    "html/template"
)


type Template struct {
	templates *template.Template
}


func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

var IndexPageRenderer *Template

type H map[string]interface{}

func Init() {
      var DB *gorm.DB
      DB, err := gorm.Open("mysql", "smartworks:smartworks@/uktrav?charset=utf8&parseTime=True")
	  if err!=nil {
            fmt.Println("DB not Connected")       
          }
      
      //static files renderer
      IndexPageRenderer = &Template{
		templates: template.Must(template.ParseGlob("templates/*.html")),
      }
   
       
      //app.Server.File("/img/woods.jpg","static/img/woods.jpg")
   
      //Routes
      app.Server.GET("/", get_index_page(DB))
      //GET all records
      app.Server.GET("/users",get_users(DB))
      
      //GET all Bloggers
      app.Server.GET("/bloggers",get_bloggers(DB))
      
      //GET record by id param
      app.Server.GET("/users/:id",get_user(DB))
      
      //GET blogger by id param
      app.Server.GET("/bloggers/:id",get_blogger(DB))
      
      //POST Create a new user
      app.Server.POST("/users",create_user(DB))
      
      //PUT Update a user
      app.Server.PUT("/users/:id",update_user(DB))
      
      //PUT Update a blogger 
      app.Server.PUT("/bloggers/:id",update_blogger(DB))
      
      //Delete delete a user
      app.Server.DELETE("/users/:id",delete_user(DB))
      
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
func get_bloggers(db *gorm.DB) echo.HandlerFunc {

    return func(c echo.Context) error {
            users :=  []models.Blogger{}
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
	        if user.Email==""{
	            c.Logger().Print("Email does not exist")
	        }
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
	            user_temp := H{}
	            //user := models.User{}
	            if err := c.Bind(&user_temp); err != nil {
		            return err
	            }
	            id, _ := strconv.Atoi(c.Param("id"))
	            
	            var user models.User
	            db.First(&user,id)
	            
	            if Lname, ok := user_temp["lname"].(string); ok{
	                user.Lname = Lname
	            }
	            Status := int(user_temp["status"].(float64))
	            user.Status = Status
	            db.Save(&user)
	        return c.JSON(http.StatusOK,user)
    }
}



func delete_user(db *gorm.DB) echo.HandlerFunc {
    return func(c echo.Context) error {
	            id, _ := strconv.Atoi(c.Param("id"))
	            var user models.User
	            db.First(&user,id)
	            db.Delete(&user)
	        return c.JSON(http.StatusOK,"Record Deleted")
    }
}

//Custom Response

func get_blogger(db *gorm.DB) echo.HandlerFunc {

    return func(c echo.Context) error {
            id, _ := strconv.Atoi(c.Param("id"))
            var blogger models.Blogger
            var user    models.User
            if db.First(&blogger,id).RecordNotFound(){
                return c.JSON(http.StatusOK,"No Record Found")    
            }
            db.Model(&blogger).Related(&user)
	        blogger_details := H{ "user_id": blogger.UserId,
	                              "blogger_name": user.Fname +" "+ user.Lname, 
	                              "short_bio": blogger.ShortBio,
	                              "status" : blogger.Status,
	                              "no_of_posts": blogger.Posts, 
	                             }
	        data := H{ "blogger_details" : blogger_details}
	        return c.JSON(http.StatusOK,data)
      }

}


func update_blogger(db *gorm.DB) echo.HandlerFunc {
    return func(c echo.Context) error {
                id, _ := strconv.Atoi(c.Param("id"))
	            var blogger models.Blogger
	            if db.First(&blogger,id).RecordNotFound(){
                    return c.JSON(http.StatusOK,"No Record Found")    
                }
	            form, err := c.MultipartForm()
                if err != nil {
                    c.Logger().Print(err.Error())
                    return err
                }
                //fmt.Println(reflect.TypeOf(form.Value["no_of_posts"][0]))
	            blogger.Posts,_  = strconv.Atoi(form.Value["no_of_posts"][0])
	            if db.Save(&blogger).Error == nil {
	                fmt.Println("update successful")
	            }
	            
	            //reading a file in form
	            fmt.Println(form.File["img"])
	            for _,file := range(form.File["img"]){
	                fmt.Println(file.Filename)
	            }
	        return c.JSON(http.StatusOK,blogger)
    }
}


func get_index_page(db *gorm.DB) echo.HandlerFunc {
    return func(c echo.Context) error {
            var user models.User
	        db.First(&user,1)
	        user_details := H{"name": user.Fname +" "+ user.Lname,
	                        }
	      
	        return c.Render(http.StatusOK, "index.html",user_details)
    }
}








