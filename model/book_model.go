package model

import "fmt"

type Book struct {
	Title     string `json:"title"`
	Author    string `json:"author"`
	Publisher string `json:"publisher"`
	Category  string `json:"category"`
	Quantity  int    `json:"quantity"`
	Isbn      int    `json:"isbn"`
}

func (b *Book) Validate() bool {
	fmt.Println(b)
	if b.Title == "" {
		return false
	}
	if b.Quantity < 1 {
		return false
	}
	if b.Isbn < 9780000000000 || b.Isbn > 9799999999999 {
		return false
	}

	return true
}
