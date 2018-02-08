package models

import (
       "github.com/jinzhu/gorm"
       "time"
)

/*
    desc user;
+---------------+-------------+------+-----+---------+----------------+
| Field         | Type        | Null | Key | Default | Extra          |
+---------------+-------------+------+-----+---------+----------------+
| id            | int(11)     | NO   | PRI | NULL    | auto_increment |
| fname         | varchar(30) | NO   |     | NULL    |                |
| lname         | varchar(30) | YES  |     | NULL    |                |
| email         | varchar(30) | YES  |     | NULL    |                |
| password_salt | varchar(30) | NO   |     | NULL    |                |
| authcode      | varchar(30) | YES  |     | NULL    |                |
| phone         | varchar(15) | NO   |     | NULL    |                |
| created_dt    | datetime    | NO   |     | NULL    |                |
| status        | tinyint(1)  | YES  |     | 0       |                |
+---------------+-------------+------+-----+---------+----------------+
*/


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
}
 
/*
get_users := func(c echo.Context) error {
             var user []User
             res := getGormDB().Find(&user)
             fmt.Println(res)
	         return c.String(http.StatusOK, "OK")
}


    Birthday     time.Time
    Age          int
    Name         string  `gorm:"size:255"` // Default size for string is 255, reset it with this tag
    

    CreditCard        CreditCard      // One-To-One relationship (has one - use CreditCard's UserID as foreign key)
    Emails            []Email         // One-To-Many relationship (has many - use Email's UserID as foreign key)

    BillingAddress    Address         // One-To-One relationship (belongs to - use BillingAddressID as foreign key)
    BillingAddressID  sql.NullInt64

    ShippingAddress   Address         // One-To-One relationship (belongs to - use ShippingAddressID as foreign key)
    ShippingAddressID int

    IgnoreMe          int `gorm:"-"`   // Ignore this field
    Languages         []Language `gorm:"many2many:user_languages;"` // Many-To-Many relationship, 'user_languages' is join table
    
 */

