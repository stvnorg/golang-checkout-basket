package main

import "testing"

func Test_DiscountPrice(t *testing.T) {
  if DiscountPrice([]string{"PEN", "TSHIRT", "MUG"}) != 32.50 {
    t.Fail()
  }

  if DiscountPrice([]string{"PEN", "TSHIRT", "PEN"}) != 25.00 {
    t.Fail()
  }

  if DiscountPrice([]string{"TSHIRT", "TSHIRT", "TSHIRT", "PEN", "TSHIRT"}) != 65.00 {
    t.Fail()
  }

  if DiscountPrice([]string{"PEN", "TSHIRT", "PEN", "PEN", "MUG", "TSHIRT", "TSHIRT"}) != 62.50 {
    t.Fail()
  }
}
