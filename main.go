package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"message": "Page not found"})
	})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/basket/total", func(c *gin.Context) {
		if IsBasketExist() {
			mybasket := GetBasketTotalAmount()
			c.JSON(200, gin.H{
				"products": mybasket.Products,
				"total_amount": mybasket.Total,
			})
		}
	})

	r.Run()
}

