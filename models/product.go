package models

type Product struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Image       string `json:"image"`
	Price       string `json:"price"`
	Category    string `json:"category"`
}
