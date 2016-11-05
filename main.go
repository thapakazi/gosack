package main

import (
	"fmt"
	"time"
)

type Khaja struct {
	Date  time.Time
	Lunch []string
	Snack []string
}

func main() {
	khaja := Khaja{Date: time.Now(), Lunch: []string{"Rice", "Daal", "Mix Veg", "Salad", "Tomato Chutney"}, Snack: []string{"Noodles Soup"}}
	// fmt.Println(khaja)
	fmt.Printf("So our lunch for %s is %s", khaja.Date.Weekday().String(), khaja.Lunch)

}
