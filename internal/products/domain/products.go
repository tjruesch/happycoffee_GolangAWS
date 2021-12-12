package domain

import (
	"fmt"
	"time"
)

type Product struct {
	name     string
	price    float32
	happyDay int
}

func NewProduct(name string, price float32, happyDay int) *Product {
	return &Product{
		name:     name,
		price:    price,
		happyDay: happyDay,
	}
}

func (p *Product) Name() string {
	return p.name
}

func (p *Product) Price() string {
	return fmt.Sprintf("%.2f", p.price) + " EUR"
}

func (p *Product) HappyDay() string {
	var day string
	switch p.happyDay {
	case 0:
		day = "Sunday"
	case 1:
		day = "Monday"
	case 2:
		day = "Tuesday"
	case 3:
		day = "Wednesday"
	case 4:
		day = "Thursday"
	case 5:
		day = "Friday"
	case 6:
		day = "Saturday"
	}

	return day
}

func (p *Product) DiscountInPercent() int {
	if int(time.Now().Weekday()) == p.happyDay {
		return 25
	}
	return 0
}
