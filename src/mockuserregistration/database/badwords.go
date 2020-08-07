package database

import (
	"database/sql"
)

//BadWords is
type BadWords interface {
	FindAll() ([]string, error)
}

//BadWordsRepository is
type BadWordsRepository struct {
	db *sql.DB
}

//NewBadWordsRepository is
func NewBadWordsRepository(db *sql.DB) *BadWordsRepository {
	return &BadWordsRepository{
		db: db,
	}
}

//FindAll is
func (dc *BadWordsRepository) FindAll() (badWordList []string, err error) {
	sql := "select name from bad-word"
	rows, err := dc.db.Query(sql)
	if err != nil {
		return badWordList, err
	}

	var badWord string
	for rows.Next() {
		err := rows.Scan(&badWord)
		if err != nil {
			return badWordList, err
		}
		badWordList = append(badWordList, badWord)
	}
	return badWordList, nil
}
