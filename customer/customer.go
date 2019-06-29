package customer

import (
	"github.com/gin-gonic/gin"
	"github.com/nompungz/finalexam/database"
	"net/http"
	"strconv"
)

type Customer struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Status string `json:"status"`
}

func CreateCustomer(c *gin.Context) {
	customer := Customer{}
	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	rows, err := database.CreateCustomer(customer.Name, customer.Email, customer.Status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	response := Customer{}
	err = rows.Scan(&response.ID, &response.Name, &response.Email, &response.Status)
	if (err != nil) {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, response)
}

func GetCustomerById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	row, err := database.GetCustomerById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	results := Customer{}
	err = row.Scan(&results.ID, &results.Name, &results.Email, &results.Status)
	if (err != nil) {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, results)
}

func GetCustomers(c *gin.Context) {
	rows, err := database.GetCustomers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	results := make([]Customer, 0)
	for rows.Next() {
		cc := Customer{}
		err = rows.Scan(&cc.ID, &cc.Name, &cc.Email, &cc.Status)
		if (err != nil) {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		results = append(results, cc)
	}
	c.JSON(http.StatusOK, results)
}

func UpdateCustomer(c *gin.Context) {
	customer := Customer{}
	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	rows, err := database.UpdateCustomer(customer.ID, customer.Name, customer.Email, customer.Status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	results := Customer{}
	err = rows.Scan(&results.ID, &results.Name, &results.Email, &results.Status)
	if (err != nil) {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, results)
}
func DeleteCustomer(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = database.DeleteCustomer(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "customer deleted"})
}
