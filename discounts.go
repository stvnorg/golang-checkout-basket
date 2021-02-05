package main

func DiscountPrice(products []string) float64 {
  ProductsAmount := make(map[string]int)

  for _, product := range products {
    ProductsAmount[product]++
  }

  total := 0.0

  for key, val := range ProductsAmount {
    qty := float64(val)

    switch key {
    case "MUG":
      total += qty * 7.5
    case "PEN":
      qty = float64((val/2) + (val%2))
      total += qty * 5.0
    case "TSHIRT":
      if qty >= 3.0 {
        total += (qty * 20.0) - (qty * 20.0 * 0.25)
      } else {
        total += qty * 20.0
      }
    }
  }

  return total
}
