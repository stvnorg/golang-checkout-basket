package main

import (
	"fmt"
	"github.com/nanobox-io/golang-scribble"
)

type Basket struct {
	Name string
	Products []string
	Total float64
}

var db, _ = scribble.New("./basket_db", nil)
var mybasket = Basket{}

// Make sure the basket exist before we add products
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
func GetBasketTotalAmount() float64 {
	db.Read("basket_db", "basket", &mybasket)
	fmt.Printf("%+v\n", mybasket)
	return mybasket.Total
}

// Add product to basket
func AddProductToBasket(product_name string) {
	db.Read("basket_db", "basket", &mybasket)
	mybasket.Products = append(mybasket.Products, product_name)
	db.Write("basket_db", "basket", mybasket)
}

func main() {
	fmt.Println(IsBasketExist())
	CreateBasket()
	fmt.Println(GetBasketTotalAmount())
	//AddProductToBasket("PEN")
	//fmt.Println(GetBasketTotalAmount())

	//AddProductToBasket("TSHIRT")
    //fmt.Println(GetBasketTotalAmount())
}
