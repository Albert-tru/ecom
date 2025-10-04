package product

import (
"database/sql"

"github.com/Albert-tru/ecom/types"
)

type Store struct {
db *sql.DB
}

func NewStore(db *sql.DB) *Store {
return &Store{db: db}
}

func (s *Store) GetProducts() ([]types.Product, error) {
rows, err := s.db.Query("SELECT id, name, description, image, price, quantity, createdat FROM products")
if err != nil {
return nil, err
}
defer rows.Close()

products := []types.Product{}
for rows.Next() {
var p types.Product
// 注意：Scan 的顺序必须和 SELECT 的顺序一致
// SELECT: id, name, description, image, price, quantity, createdat
if err := rows.Scan(&p.ID, &p.Name, &p.Description,
&p.ImageURL, &p.Price, &p.Quantity, &p.CreatedAt); err != nil {
return nil, err
}
products = append(products, p)
}

return products, nil

}

func scanRowIntoProduct(row *sql.Row) (*types.Product, error) {
product := new(types.Product)

err := row.Scan(
&product.ID,
&product.Name,
&product.Description,
&product.ImageURL,
&product.Price,
&product.Quantity,
&product.CreatedAt,
)
if err != nil {
return nil, err
}
return product, err
}
