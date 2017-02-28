package model

import (
	"github.com/gocraft/dbr"
	"MicroFilm/util"
)

type Product struct {
	Id          int64   `json:"id"`
	ProductName string   `db:"product_name" json:"productName"`
	ProductType int64   `db:"product_type" json:"productType"`
	Channel     string    `json:"channel"`
	Price       int64    `json:"price"`
	PriceOrg    int64   `db:"price_org" json:"priceOrg"`
	PriceUnit   int64   `db:"price_unit" json:"priceUnit"`
}

func (m *Product) Save(tx *dbr.Tx) error {

	_, err1 := tx.InsertInto("bis_product").
		Columns(util.BuildColumnName(m)...).
		Record(m).
		Exec()

	return err1
}

func (m *Product) Load(tx *dbr.Tx, id int64) error {

	return tx.Select(util.BuildColumnName(m)...).
		From("bis_product").
		Where("id = ?", id).
		LoadStruct(m)
}
