package main

import (
    "fmt"
    "github.com/gin-gonic/gin"
    "github.com/banzaicloud/go-gin-prometheus"
    "strings"
)

var message string

// Route for custom http error 404
func GET_CustomHTTP404(c *gin.Context) {
  c.JSON(404, gin.H{"message": "Page not found"})
}

// Route for creating a basket
func POST_CreateBasket(c *gin.Context) {
  CreateBasket()

  if IsBasketExist() {
    message = "Basket has been created"
  } else {
    message = "Error creating basket"
  }

  c.JSON(201, gin.H{
    "message": message,
  })
}

// Route for adding a product to basket
func POST_AddProductToBasket(c *gin.Context) {
  if IsBasketExist() {
    code := strings.ToUpper(c.Param("code"))
    if code == "MUG" || code == "PEN" || code == "TSHIRT" {
      AddProductToBasket(code)
      message = "Product successfully added to the basket"
    } else {
      message = "Product code doesn't exist"
    }
  } else {
    message = "Basket doesn't exist"
  }

  c.JSON(200, gin.H{
    "message": message,
  })
}

// Route for getting the total amount in basket
func GET_TotalAmountInBasket(c *gin.Context) {
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
}

// Route to remove basket
func DELETE_DeleteBasket(c *gin.Context) {
  if DeleteBasket() {
    message = "Basket has been deleted"
  } else {
    message = "Error deleting basket"
  }

  c.JSON(200, gin.H{
    "message": message,
  })
}

func main() {
  r := gin.Default()

  p := ginprometheus.NewPrometheus("gin", []string{})
  p.SetListenAddress(":9900")
  p.Use(r, "/metrics")

  r.NoRoute(GET_CustomHTTP404)
  r.POST("/basket/create", POST_CreateBasket)
  r.POST("/basket/add_product/:code", POST_AddProductToBasket)
  r.GET("/basket/total", GET_TotalAmountInBasket)
  r.DELETE("/basket/delete", DELETE_DeleteBasket)

  r.Run()
}
