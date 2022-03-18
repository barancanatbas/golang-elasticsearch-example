package requests

type Create struct {
	Name  string `json:"name"`
	Color string `json:"color"`
	Price int    `json:"price"`
	Stock int    `json:"stock"`
}

type Update struct {
	Name  string `json:"name"`
	Color string `json:"color"`
	Price int    `json:"price"`
	Stock int    `json:"stock"`
	ID    string `json:"id"`
}
