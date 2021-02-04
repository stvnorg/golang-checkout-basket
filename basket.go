package main

import (
	"github.com/nanobox-io/golang-scribble"
)

type Basket struct {
	Name string
	Products []string
	Total float64
}

var db, _ = scribble.New("./basket_db", nil)

var mybasket = Basket{}

// Check if basket exist before we add products
func IsBasketExist() bool {
	db.Read("basket_db", "basket", &mybasket)
	return mybasket.Name != ""
}

// Create the basket
func CreateBasket() {
	mybasket.Name = "mybasket"
	db.Write("basket_db", "basket", mybasket)
}

// Get the basket
func GetBasketTotalAmount() Basket {
	db.Read("basket_db", "basket", &mybasket)
	return mybasket
}

// Add product to basket
func AddProductToBasket(product_name string) {
	db.Read("basket_db", "basket", &mybasket)
	mybasket.Products = append(mybasket.Products, product_name)
	mybasket.Total = DiscountPrice(mybasket.Products)
	db.Write("basket_db", "basket", mybasket)
}

// Delete the basket
func DeleteBasket() bool {
	db.Delete("basket_db", "")
	mybasket.Name = ""
	return !IsBasketExist()
}
