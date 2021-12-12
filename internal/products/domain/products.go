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
		day = "Monday"
	case 1:
		day = "Tuesday"
	case 2:
		day = "Wednesday"
	case 3:
		day = "Thursday"
	case 4:
		day = "Friday"
	case 5:
		day = "Saturday"
	case 6:
		day = "Sunday"
	}

	return day
}

func (p *Product) DiscountInPercent() int {
	if time.Now().Day() == p.happyDay {
		return 25
	}
	return 0
}
