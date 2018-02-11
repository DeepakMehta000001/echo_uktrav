package main

import (
  "uktrav_echo/app" 
  "uktrav_echo/app/controllers"  
)


func main() {
    //init server
    app.Init()
    
    //init controllers
    controllers.Init()
    
    //registering static renderer
    app.Server.Renderer = controllers.IndexPageRenderer
    
    // run server
    app.Server.Logger.Fatal(app.Server.Start(":1323"))
}
