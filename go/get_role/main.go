package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	router.GET("/role", GET_ROLE)

	router.GET("/role/:id", GET_ROLE_ID)

	router.POST("/role/", POST_ROLE)

	router.PUT("/role/:id", PUT_ROLE)

	router.DELETE("/role/:id", DELETE_ROLE)

	router.Run(":8084")
}

func DELETE_ROLE(c *gin.Context) {
	id := c.Param("id")
	idValue, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusNotFound, "Error")
	}
	Data = append(Data[:int(idValue)-1], Data[int(idValue):]...)
	fmt.Println(Data)
}

func PUT_ROLE(c *gin.Context) {
	var Update_Role Role
	id := c.Param("id")
	idValue, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusNotFound, "Error")
	}
	if c.ShouldBind(&Update_Role) == nil {
		Data[int(idValue)-1] = Update_Role
	}

}

func POST_ROLE(c *gin.Context) {

	var new_role Role
	if c.ShouldBind(&new_role) == nil {
		log.Println(new_role.ID)
		log.Println(new_role.Name)
		log.Println(new_role.Summary)
		Data = append(Data, new_role)

	} else {
		c.JSON(http.StatusBadRequest, "ERROR")

	}

}

func GET_ROLE(c *gin.Context) {
	c.JSON(http.StatusOK, Data)
}

func GET_ROLE_ID(c *gin.Context) {
	id := c.Param("id")
	idValue, err := strconv.Atoi(id)
	fmt.Println(idValue)
	if err != nil {
		c.JSON(http.StatusNotFound, "Error")
	}
	for _, Item := range Data {
		fmt.Println("Item", Item.ID, "ID_Value", idValue)
		if Item.ID == uint(idValue) {
			c.JSON(http.StatusOK, Item)
			break
		} else {
			fmt.Println("ERROR")
		}
	}

}
