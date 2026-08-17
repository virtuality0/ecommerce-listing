package product

const (
	queryCreateProduct = `INSERT INTO product(id, name, description, price, stock)
  VALUES (
    $1, $2, $3, $4, $5
  )
  `
	queryGetProductById = `SELECT id, name, description, price, stock, created_at, modified_at
  FROM product 
  WHERE id = $1
  `

	queryGetProductList = `SELECT id, name, description, price, stock, created_at, modified_at 
  FROM product 
  ORDER BY created_at DESC
  OFFSET $1 
  LIMIT $2
  `

	queryUpdateProduct = `UPDATE product 
  SET name = COALESCE($1, name), description = COALESCE($2, description), price = COALESCE($3, price), stock = COALESCE($4, stock), modified_at = NOW() 
  WHERE id = $5
  RETURNING id, name, description, price, stock, created_at, modified_at
  `

	queryDeleteProduct = `DELETE FROM product 
  WHERE id = $1`

	queryGetProductCount = `SELECT Count(id) FROM product 
  OFFSET $1 
  LIMIT $2`
)
