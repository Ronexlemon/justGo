package models


type Store struct {
	StoreId       int       `json:id"`
	StoreName     string    `json:"name"`
	StoreAddress  string    `json:"address"`
	StoreProducts *Product `json:"products"`
}

type Product struct {
	ProductId   int    `json:"id"`
	ProductName string `json:"name"`

	ProductPrice int `json:"price"`
}

var Stores []Store

func (s *Store) IsEmpty() bool {
	return s.StoreName == " "

}
