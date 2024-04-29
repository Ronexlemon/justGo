package models


type Store struct {
	StoreId       string      `json:id"`
	StoreName     string    `json:"name"`
	StoreAddress  string    `json:"address"`
	StoreProducts *Product `json:"products"`
}

type Product struct {
	ProductId   string   `json:"id"`
	ProductName string `json:"name"`

	ProductPrice int `json:"price"`
}

var Stores []Store

func (s *Store) IsEmpty() bool {
	return s.StoreName == " "

}
