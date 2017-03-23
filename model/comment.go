package model

import (
	"github.com/gocraft/dbr"
	"MicroFilm/util"
)

type Comment struct {
	Id         int64  `json:"id"`
	Uid        int64  `json:"uid"`
	Mid        int64  `db:"mid" json:"mid"`
	Content    string  `db:"content" json:"content"`
	ZanCount   int     `db:"zan_count" json:"zanCount"`
	PhoneType  int   `db:"phone_type" json:"phoneType"`
	AnonEnable int   `db:"anon_enable" json:"anonEnable"`
	Status     int   `db:"status" json:"status"`
	CreateTime int64   `db:"create_time" json:"CreateTime"`
}

func (m *Comment) Save(tx *dbr.Tx) error {

	_, err := tx.InsertInto("mv_comment").
		Columns(util.BuildColumnName(m)...).
		Record(m).
		Exec()

	return err
}

func (m *Comment) Delete(tx *dbr.Tx, id int64) error {

	_, err := tx.DeleteFrom("mv_comment").
		Where("id = ?", id).
		Exec()

	return err
}

func (m *Comment) Update(tx *dbr.Tx) error {

	_, err := tx.Update("mv_comment").
		SetMap(util.StructMap(m)).
		Where("id = ?", m.Id).
		Exec()

	return err
}

type Comments []Comment

func (m *Comments) Load(tx *dbr.Tx, uid int64, mid int64) error {

	return tx.Select(util.BuildColumnName(&Comment{})...).
		From("mv_comment").
		Where("uid = ? and mid = ?", uid, mid).
		LoadStruct(m)
}