package main

import (
    "fmt"
    "github.com/gin-gonic/gin"
    "strings"
)

var message string

func main() {
  r := gin.Default()
  r.NoRoute(func(c *gin.Context) {
    c.JSON(404, gin.H{"message": "Page not found"})
  })

  r.POST("/basket/create", func(c *gin.Context) {
    CreateBasket()

    if IsBasketExist() {
      message = "Basket has been created"
    } else {
      message = "Error creating basket"
    }

    c.JSON(201, gin.H{
      "message": message,
    })
  })


  r.POST("/basket/add_product", func(c *gin.Context) {
    code := strings.ToUpper(c.Query("code"))

    if code == "MUG" || code == "PEN" || code == "TSHIRT" {
      AddProductToBasket(code)
      message = "Product successfully added to the basket"
    } else {
      message = "Product code doesn't exist"
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
        "total_amount": fmt.Sprintf("%.2f%s", mybasket.Total, "€"),
      })
    } else {
      c.JSON(200, gin.H{
        "message": "Basket doesn't exist",
      })
    }
  })

  r.DELETE("/basket/delete", func(c *gin.Context) {
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
