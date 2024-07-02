package types


type Book struct {
	ID        int            `json:"id" db:"id"`
	Name      string         `json:"name" db:"name"`
	Author    string         `json:"author" db:"author"`
	OwnedFrom string         `json:"owned_from" db:"owned_from"`
	IsAvail   bool           `json:"is_avail" db:"is_avail"`
	Price     float32        `json:"price" db:"price"`
}

type Books struct {
    Books []Book `json:"books"`
}

type User struct{

	ID         int
	Name       string
	IsAdmin    bool
	Password   string

}