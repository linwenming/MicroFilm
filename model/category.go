package model

import (
	"github.com/gocraft/dbr"
	"MicroFilm/util"
)

type Category struct {
	Id     int64   `json:"id"`
	Name   string   `db:"name" json:"name"`
	Parent int64    `json:"parent"`
	Level  int64    `json:"level"`
}

func NextBillNumber() {

}


func (m *Category) Save(tx *dbr.Tx) error {

	_, err := tx.InsertInto("mv_category").
		Columns(util.BuildColumnName(m)...).
		Record(m).
		Exec()

	return err
}

func (m *Category) Delete(tx *dbr.Tx, id int64) error {

	_, err := tx.DeleteFrom("mv_category").
		Where("id = ?", id).
		Exec()

	return err
}

func (m *Category) Update(tx *dbr.Tx) error {

	_, err := tx.Update("mv_category").
		SetMap(util.StructMap(m)).
		Where("id = ?", m.Id).
		Exec()

	return err
}

func (m *Category) Load(tx *dbr.Tx, id int64) error {

	return tx.Select("*").
		From("mv_category").
		Where("id = ?", id).
		LoadStruct(m)
}

type CategoryList []Category

func (m *CategoryList) Load(tx *dbr.Tx) error {

	return tx.Select("*").
		From("mv_category").
		Where("parent = 0").
		LoadStruct(m)
}

