package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type property struct {
	ID        string `json:"id"`
	Address   string `json:"address"`
	OwnerName string `json:"ownername"`
	Price     int    `json:"price"`
}

var propertys = []property{
	{ID: "Prop1", Address: "Guwahati", OwnerName: "Abhishek", Price: 10000000},
	{ID: "Prop2", Address: "Delhi", OwnerName: "Arjun", Price: 9000000},
	{ID: "Prop3", Address: "Chandigarh", OwnerName: "Karn", Price: 15000000},
}

func getProps(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, propertys)
}

func addprop(c *gin.Context) {
	var newprop property

	if err := c.BindJSON(&newprop); err != nil {
		return
	}

	propertys = append(propertys, newprop)
	c.IndentedJSON(http.StatusCreated, propertys)

}

func UpdateProp(c *gin.Context) {
	id := c.Param("id")
	var propdata property

	if err := c.BindJSON(&propdata); err != nil {
		return
	}

	for i, p := range propertys {
		if p.ID == id {
			propertys[i] = propdata
		}
	}

	c.IndentedJSON(http.StatusOK, propertys)
}

func deleteprop(c *gin.Context) {
	id := c.Param("id")

	for i, p := range propertys {
		if p.ID == id {
			propertys = append(propertys[:i], propertys[i+1:]...)
		}
	}

	c.IndentedJSON(http.StatusOK, propertys)
}

func main() {
	fmt.Println("Started")
	router := gin.Default()
	router.GET("/props", getProps)
	router.POST("/addprop", addprop)
	router.PUT("/updateprops/:id", UpdateProp)
	router.DELETE("/deleteprops/:id", deleteprop)
	router.Run("localhost:8080")
}
