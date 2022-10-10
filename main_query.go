package main

import (
	"fmt"
	"net/http"

	// "net/http"
	"database/sql"
	"log"

	// "context"
	// "github.com/denisenkom/go-mssqldb"
	// "github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type property struct {
	ID        string `json:"id"`
	Address   string `json:"address"`
	OwnerName string `json:"ownername"`
	Price     int    `json:"price"`
}

// var server = "localhost"
// var port = 3306
// var user = "abhishek"
// var password = "1234"
// var database = "//properties"

var db *sql.DB

// var propertys = []property{
// 	{ID: "Prop1", Address: "Guwahati", OwnerName: "Abhishek", Price: 10000000},
// 	{ID: "Prop2", Address: "Delhi", OwnerName: "Arjun", Price: 9000000},
// 	{ID: "Prop3", Address: "Chandigarh", OwnerName: "Karn", Price: 15000000},
// }

func getProps(c *gin.Context) {

	result, err := db.Query("SELECT * FROM propertydata")
	if err != nil {
		panic(err.Error())
	}
	// result.Close()
	var prop []property
	for result.Next() {
		var p property
		// func (rs *Rows) Scan(dest ...interface{}) error
		err = result.Scan(&p.ID, &p.Address, &p.OwnerName, &p.Price)
		if err != nil {
			panic(err)
		}
		// func append(slice []Type, elems ...Type) []Type
		prop = append(prop, p)
	}
	result.Close()
	c.IndentedJSON(http.StatusOK, prop)
}

func main() {
	var err error
	// connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
	// 	server, user, password, port, database)
	fmt.Println("!... Hello World ...!")
	db, err = sql.Open("mysql", "root:1234@tcp(localhost:3306)/properties")
	if err != nil {
		log.Fatal("Error creating connection pool: " + err.Error())
	}
	log.Printf("Connected!\n")

	router := gin.Default()
	router.GET("/props", getProps)
	// router.POST("/addprop", addprop)
	// router.PUT("/updateprops/:id", UpdateProp)
	// router.DELETE("/deleteprops/:id", deleteprop)
	router.Run("localhost:8080")

}
