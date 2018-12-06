package entities

// Article - base type of object that everything evolves on.
type Article struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}
