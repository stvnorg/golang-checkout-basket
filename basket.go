package main

import (
	"github.com/nanobox-io/golang-scribble"
	"os"
)

type Basket struct {
	Name string
	Products []string
	Total float64
}

func getEnv(key, fallback string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		value = fallback
	}
	return value
}

var db_name = getEnv("DB_NAME", "basket_db")

var db, _ = scribble.New("./" + db_name, nil)

var mybasket = Basket{}

// Check if basket exist before adding products
func IsBasketExist() bool {
	db.Read(db_name, "basket", &mybasket)
	return mybasket.Name != ""
}

// Create the basket
func CreateBasket() {
	mybasket.Name = "mybasket"
	db.Write(db_name, "basket", mybasket)
}

// Get the basket
func GetBasketTotalAmount() Basket {
	db.Read(db_name, "basket", &mybasket)
	return mybasket
}

// Add product to basket
func AddProductToBasket(product_name string) {
	db.Read(db_name, "basket", &mybasket)
	mybasket.Products = append(mybasket.Products, product_name)
	mybasket.Total = DiscountPrice(mybasket.Products)
	db.Write(db_name, "basket", mybasket)
}

// Delete the basket
func DeleteBasket() bool {
	db.Delete(db_name, "")
	mybasket.Name = ""
	return !IsBasketExist()
}
