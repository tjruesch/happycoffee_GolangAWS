package domain

import (
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

func (p *Product) Price() float32 {
	return p.price
}

func (p *Product) HappyDay() int {
	return p.happyDay
}

func (p *Product) DiscountInPercent() int {
	if int(time.Now().Weekday()) == p.happyDay {
		return 25
	}
	return 0
}
