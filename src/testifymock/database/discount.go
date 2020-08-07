package database

import "database/sql"

//DiscountRepository is
type DiscountRepository struct {
	db *sql.DB
}

//Discount is
type Discount interface {
	FindCurrentDiscount() int
}

//NewDiscountRepository is
func NewDiscountRepository(db *sql.DB) *DiscountRepository {
	return &DiscountRepository{
		db: db,
	}
}

//FindCurrentDiscount is
func (dc *DiscountRepository) FindCurrentDiscount() (discount int) {
	sql := "select value from discount limit 1"
	row := dc.db.QueryRow(sql)
	row.Scan(&discount)
	return discount
}
