package main

import (
  "database/sql"
  "fmt"
  "net/http"

  _ "github.com/go-sql-driver/mysql"
  "github.com/labstack/echo"
  "github.com/labstack/echo/middleware"
  "github.com/JonathanMH/goClacks/echo"
)

type (
  Excuse struct {
    Error string `json:"error"`
    Id string `json:"id"`
    Quote string `json:"quote"`
  }
)

func main() {
  // Echo instance
  e := echo.New()
  e.Use(goClacks.Terrify)

  // Middleware
  e.Use(middleware.Logger())
  e.Use(middleware.Recover())

  e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
    AllowOrigins: []string{"*"},
    AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
  }))

  // Route => handler
  e.GET("/", func(c echo.Context) error {
    db, err := sql.Open("mysql", "smartworks:smartworks@tcp(127.0.0.1:3306)/uktrav")

    if err != nil {
      fmt.Println(err.Error())
      response := Excuse{Id: "", Error: "true", Quote: ""}
      return c.JSON(http.StatusInternalServerError, response)
    }
    defer db.Close()

    var quote string;
    var id string;
    err = db.QueryRow("SELECT id,fname FROM user LIMIT 1").Scan(&id, &quote)

    if err != nil {
      fmt.Println(err)
    }

    fmt.Println(quote);
    response := Excuse{Id: id, Error: "false", Quote: quote}
    return c.JSON(http.StatusOK, response)
  })
/*
  e.GET("/id/:id", func(c echo.Context) error {
    requested_id := c.Param("id")
    fmt.Println(requested_id);
      db, err := sql.Open("mysql", "db_user:db_password@tcp(SERVER_IP:PORT)/database_name")

    if err != nil {
      fmt.Println(err.Error())
      response := Excuse{Id: "", Error: "true", Quote: ""}
      return c.JSON(http.StatusInternalServerError, response)
    }
    defer db.Close()

    var quote string;
    var id string;
    err = db.QueryRow("SELECT id, quote FROM excuses WHERE id = ?", requested_id).Scan(&id, &quote)

    if err != nil {
      fmt.Println(err)
    }

    response := Excuse{Id: id, Error: "false", Quote: quote}
    return c.JSON(http.StatusOK, response)
  })
 */
  e.Logger.Fatal(e.Start(":1323"))
}
