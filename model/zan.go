package model

import (
	"github.com/gocraft/dbr"
	"MicroFilm/util"
)

type Zan struct {
	Id         string  `json:"id"`
	Uid        int64   `json:"uid"`
	Mid        int64   `json:"mid"`
	Cid        int64   `json:"cid"`
	CreateTime int64   `db:"create_time" json:"createTime"`
}

func (m *Zan) Save(tx *dbr.Tx) error {

	_, err := tx.InsertInto("mv_zan").
		Columns(util.BuildColumnName(m)...).
		Record(m).
		Exec()
	return err
}

func (m *Zan) Delete(tx *dbr.Tx, id int64) error {

	_, err := tx.DeleteFrom("mv_zan").
		Where("id = ?", id).
		Exec()

	return err
}

func (m *Zan) Update(tx *dbr.Tx) error {

	_, err := tx.Update("mv_zan").
		SetMap(util.StructMap(m)).
		Where("id = ?", m.Id).
		Exec()

	return err
}

func (m *Zan) UpdateBy(tx *dbr.Tx, value map[string]interface{}) error {

	_, err := tx.Update("mv_zan").
		SetMap(value).
		Where("id = ?", m.Id).
		Exec()

	return err
}

func (m *Zan) LoadBy(tx *dbr.Tx, uid int64,mid int64) error {

	return tx.Select(util.BuildColumnName(m)...).
		From("mv_zan").
		Where("uid = ? and mid = ?", uid,mid).
		LoadStruct(m)
}