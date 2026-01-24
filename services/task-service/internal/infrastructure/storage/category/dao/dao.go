package dao

type Category struct {
	ID    string  `json:"id"`
	Title string  `json:"title"`
	Desc  string  `json:"desc"`
	Price float64 `json:"price"`
}
