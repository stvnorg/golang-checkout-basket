package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"message": "Page not found"})
	})

	r.POST("/basket/create", func(c *gin.Context) {
		message := ""
		
		CreateBasket()
		
		if IsBasketExist() {
			message = "Basket has been created"
		} else {
			message = "Failed to create a basket"
		}

		c.JSON(200, gin.H{
			"message": message,
		})
	})

	r.GET("/basket/total", func(c *gin.Context) {
		if IsBasketExist() {
			mybasket := GetBasketTotalAmount()
			c.JSON(200, gin.H{
				"products": mybasket.Products,
				"total_amount": mybasket.Total,
			})
		} else {
			c.JSON(200, gin.H{
				"message": "Basket doesn't exist",
			})
		}
	})

	r.DELETE("/basket/delete", func(c *gin.Context) {
		message := "" 
		
		if DeleteBasket() {
			message = "Basket has been deleted"
		} else {
			message = "Error deleting basket"
		}

		c.JSON(200, gin.H{
            "message": message,
        })
	})

	r.Run()
}

