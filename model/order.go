package model

import (
	"github.com/gocraft/dbr"
	"MicroFilm/util"
)

type OrderDetail struct {
	Id     int64   `json:"id"`
	Name   string   `db:"name" json:"name"`
	Parent int64    `json:"parent"`
	Level  int64    `json:"level"`
}

type OrderLog struct {
	Id     int64   `json:"id"`
	Name   string   `db:"name" json:"name"`
	Parent int64    `json:"parent"`
	Level  int64    `json:"level"`
}

func (m *OrderDetail) Save(tx *dbr.Tx) error {

	_, err := tx.InsertInto("bis_order").
		Columns(util.BuildColumnName(m)...).
		Record(m).
		Exec()

	return err
}

