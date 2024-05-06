package Application

type Product struct {
	ID          int    `json:"ID" db:"id"`
	Name        string `json:"Name" db:"name"`
	Description string `json:"Description" db:"description"`
	Price       int    `json:"Price" db:"price"`
	Count       int    `json:"Count" db:"count"`
}
