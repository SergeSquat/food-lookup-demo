package api

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

const (
	q = `select ifnull(carbohydrate_g, 0), ifnull(protein_g, 0), ifnull(fa_sat_g, 0), ifnull(fa_mono_g, 0),
	         ifnull(fa_poly_g, 0), ifnull(kcal, 0), description
                 from entries where description like ? limit ?`
)

func GetProducts(mask string, limit int) ([]Product, error) {
	db, err := sql.Open("sqlite3", "./db/usda-nnd.sqlite3")

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	stmt, err := db.Prepare(q)

	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	mask = "%" + mask + "%"

	rows, err := stmt.Query(fmt.Sprintf("%%%s%%", mask), limit)

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	var ps []Product

	for rows.Next() {

		var carbohydrate_g, protein_g, fa_sat_g, fa_mono_g, fa_poly_g, kcal float64
		var description string

		err = rows.Scan(&carbohydrate_g, &protein_g, &fa_sat_g, &fa_mono_g, &fa_poly_g, &kcal,
			&description)

		if err != nil {
			log.Fatal(err)
		}

		p := Product{carbohydrate_g, protein_g, fa_sat_g, fa_mono_g,
			fa_poly_g, kcal, description}

		ps = append(ps, p)

	}

	return ps, err
}
