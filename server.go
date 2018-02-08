package main

import (
  "uktrav_echo/app" 
  "uktrav_echo/app/controllers"  
  "uktrav_echo/db/gorm"
)


func main() {
    //init server
    app.Init()
    
    //init db
    gorm.Init()
    
    //init controllers
    controllers.Init()
    // run server
    app.Server.Logger.Fatal(app.Server.Start(":1323"))
}
