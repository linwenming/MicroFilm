package model

import (
	"time"
	"github.com/gocraft/dbr"
	"MicroFilm/util"
)

type Movie struct {
	Id         int64   `json:"id"`
	Name       string  `json:"name"`
	Intro      string  `json:"intro"`
	CategoryId int64   `db:"category_id" json:"categoryId"`
	PreviewImg string  `db:"preview_img" json:"previewImg"`
	PlayUrl    string  `db:"play_url" json:"playUrl"`
	PlayLength int64  `db:"play_length" json:"playLength"`
	FileSize   int64  `db:"file_size" json:"fileSize"`
	Tags       string  `json:"tags"`
	Status     int64     `json:"status"`
	Uploader   int64   `json:"uploader"`
	CreateTime int64   `db:"create_time" json:"createTime"`
	OnlineTime int64   `db:"online_time" json:"onlineTime"`
	Director   string  `json:"director"`
	Actor      string  `json:"actor"`
	Highlight  string  `json:"highlight"`
	Score      int64     `json:"score"`
	PlayCount  int64   `db:"play_count" json:"playCount"`
	ReplyCount int64   `db:"reply_count" json:"replyCount"`
	ZanCount   int64   `db:"zan_count" json:"zanCount"`
}

type MovieForm struct {
	Id         int64   `json:"id"`
	Name       string  `json:"name"`
	Intro      string  `json:"intro"`
	CategoryId int64   `db:"category_id" json:"categoryId"`
	PreviewImg string  `db:"preview_img" json:"previewImg"`
	PlayUrl    string  `db:"play_url" json:"playUrl"`
	PlayLength int64  `db:"play_length" json:"playLength"`
	FileSize   int64  `db:"file_size" json:"fileSize"`
	Tags       string  `json:"tags"`
	Status     int     `json:"status"`
	Uploader   int64   `json:"uploader"`
	CreateTime int64   `db:"create_time" json:"createTime"`
	OnlineTime int64   `db:"online_time" json:"onlineTime"`
}

func NewMovieForm() *MovieForm {
	return &MovieForm{
		Status: 0,
		CreateTime: time.Now().Unix(),
		OnlineTime: time.Now().Unix(),
	}
}

func NewMovie() *Movie {
	return &Movie{
		Status: 0,
		CreateTime: time.Now().Unix(),
		OnlineTime: time.Now().Unix(),
	}
}

func (m *MovieForm) Save(tx *dbr.Tx) error {

	_, err := tx.InsertInto("mv_film").
		Columns(util.BuildColumnName(m)...).
		Record(m).
		Exec()

	return err
}

func (m *MovieForm) Delete(tx *dbr.Tx, id int64) error {

	_, err := tx.DeleteFrom("mv_film").
		Where("id = ?", id).
		Exec()

	return err
}

func (m *MovieForm) Update(tx *dbr.Tx) error {

	_, err := tx.Update("mv_film").
		SetMap(util.StructMap(m)).
		Where("id = ?", m.Id).
		Exec()

	return err
}

func (m *Movie) UpdateBy(tx *dbr.Tx, value map[string]interface{}) error {

	_, err := tx.Update("mv_film").
		SetMap(value).
		Where("id = ?", m.Id).
		Exec()

	return err
}

func (m *MovieForm) Load(tx *dbr.Tx, id int64) error {

	return tx.Select(util.BuildColumnName(m)...).
		From("mv_film").
		Where("id = ?", id).
		LoadStruct(m)
}

func (m *Movie) Load(tx *dbr.Tx, id int64) error {

	return tx.Select(util.BuildColumnName(m)...).
		From("mv_film").
		Where("id = ?", id).
		LoadStruct(m)
}

type Movies []Movie

func (m *Movies) LoadByCate(tx *dbr.Tx, cateId int64) error {

	return tx.Select(util.BuildColumnName(&Movie{})...).
		From("mv_film").
		Where("category_id = ?", cateId).
		LoadStruct(m)
}

func (m *Movies) LoadByFuzzy(tx *dbr.Tx, keyword string) error {

	return tx.Select(util.BuildColumnName(&Movie{})...).
		From("mv_film").
		Where("CONCAT(`name`,`tags`) LIKE ?", "%" + keyword + "%").
		LoadStruct(m)
}