package main

import (
	"github.com/gin-gonic/gin"
	"github.com/nompungz/finalexam/customer"
	"github.com/nompungz/finalexam/middleware"
)

func main() {
	r := setupRouter()
	r.Run(":2019")
	//fmt.Print("555")
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.Authorization)
	r.POST("/customers", customer.CreateCustomer)
	r.GET("/customers/:id", customer.GetCustomerById)
	r.GET("/customers", customer.GetCustomers)
	r.PUT("/customers/:id", customer.UpdateCustomer)
	r.DELETE("/customers/:id", customer.DeleteCustomer)
	return r
}
