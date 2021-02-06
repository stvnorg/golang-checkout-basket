package main

import "testing"

import (
	"github.com/nanobox-io/golang-scribble"
)

var test_db, _ = scribble.New("./" + GetEnv("DB_NAME", "test_db"), nil)

func Test_CreateBasket(t *testing.T) {
    CreateBasket()
}

func Test_BasketExist(t *testing.T) {
  if !IsBasketExist() {
    t.Fail()
  }
}

func Test_AddProductToBasket(t *testing.T) {
	AddProductToBasket("PEN")
}

func Test_GetBasketTotalAmount(t *testing.T) {
	test_basket := GetBasketTotalAmount()
	if test_basket.Total != 5.0 {
		t.Fail()
	}
}

func Test_DeleteBasket(t *testing.T) {
	DeleteBasket()
}
