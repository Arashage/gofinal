package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Arashage/gofinal/customer"
	"github.com/Arashage/gofinal/database"
	"github.com/gin-gonic/gin"
)

func postCustomerHandler(c *gin.Context) {
	var rq customer.Customer
	err := c.ShouldBindJSON(&rq)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error,
		})
		return
	}

	err = database.InsertCustomer(&rq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error,
		})
		return
	}

	c.JSON(http.StatusCreated, rq)
}

func getCustomerByIDHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error,
		})
		return
	}

	rs, err2 := database.GetCustomerByID(id)
	if err2 != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err2.Error,
		})
		return
	}

	c.JSON(http.StatusOK, rs)
}

func getAllCustomerHandler(c *gin.Context) {
	rs, err := database.GetAllCustomer()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error,
		})
		return
	}

	c.JSON(http.StatusOK, rs)
}

func putCustomerHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error,
		})
		return
	}

	var rq customer.Customer
	err = c.ShouldBindJSON(&rq)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error,
		})
		return
	}

	rq.ID = id

	err = database.UpdateCustomer(&rq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error,
		})
		return
	}

	c.JSON(http.StatusOK, rq)
}

func deleteCustomerHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error,
		})
		return
	}

	err = database.DeleteCustomer(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "customer deleted",
	})
}

func main() {
	fmt.Println("customer service")
	//run port ":2009"
	r := gin.Default()

	database.Connect()
	defer database.DB.Close()

	database.CreateCustomer()

	r.POST("/customers", postCustomerHandler)
	r.GET("/customers/:id", getCustomerByIDHandler)
	r.GET("/customers", getAllCustomerHandler)
	r.PUT("/customers/:id", putCustomerHandler)
	r.DELETE("/customers/:id", deleteCustomerHandler)
	r.Run(":2009")
}
