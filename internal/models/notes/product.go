package notes

import "github.com/lilpipidron/sugar-backend/internal/models/products"

type NoteProduct struct {
	Product products.Product
	Amount  int
}
